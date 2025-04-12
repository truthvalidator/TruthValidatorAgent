package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"ssaisearch/api/model"
	"ssaisearch/api/reqmodel"
	"ssaisearch/api/shared"
	"ssaisearch/llm"
	"ssaisearch/rag"
	"ssaisearch/rag/searxng"
	"ssaisearch/utils"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

// VerifyProposal handles the core truth verification workflow:
// 1. Parses and validates the incoming query
// 2. Searches for evidence using SearxNG metasearch 
// 3. Analyzes results with LLM to generate judgment
// 4. Stores verification evidence on Web3 storage
// 5. Streams response back to client
func VerifyProposal(c *gin.Context) {
	var query reqmodel.Query // Incoming verification request
	if err := c.BindJSON(&query); err != nil {
		utils.GinErrorMsg(c, err)
		return
	}
	query = query.Regularize()
	if query.Question == "" {
		c.String(200, "")
		return
	}

	ans := model.Answer{
		Question: query.Question,
	}

	searchResultsText := "" // Aggregated text from decentralized search results

	// Configure decentralized search via SearxNG metasearch engine
	config := searxng.SearxNGConfig{
		Endpoint: "http://127.0.0.1:18080/search", // SearxNG instance endpoint
	}

	queryStr := query.Question

	// Execute decentralized search and process results
	searchRaw, err := searxng.SearchRawFromSearxNG(config, query.Country(), query.Locale(), queryStr)
	if err != nil {
		log.Println("SearchRawFromSearxNG error:", err)
	} else {
		// Convert and aggregate search results for LLM analysis
		serpResults := ConvertSearxNGResultsToSERPResults(searchRaw.SpiderResults)
		searchResultsText = rag.SearchResultsToText(serpResults)
	}

	ragRaw := ConvertSearxNGRawItemToRagRawItem(searchRaw)

	meta := reqmodel.NewMeta(ragRaw)
	meta.ID = ans.ID
	c.String(200, "%s\r\n", meta.Json())
	ans.Evidence = utils.Json(meta.Evidences)

	if query.ProgLang == "" {
		query.ProgLang = "golang"
	}
	guideStr := query.ProgLang

	if str, err := GenerateProgramGuideStr(query.ProgLang); err == nil {
		guideStr = str
	}

	// Prepare LLM prompt with search results and execute AI analysis
	content := llm.Promptc(query.Language, query.Field, query.Question, guideStr, searchResultsText)

	// Configure LLM request with system prompts and user query
	req := utils.ModelGPT35Request([]openai.ChatCompletionMessage{
		utils.ChatMsgFromSystem(content),
		utils.ChatMsgFromSystem(`{
        "output_format": "json",
        "output_schema": {
            "judge": "boolean",
            "reason": "string"
        }
    }`),
		utils.ChatMsgFromUser(query.Question),
	})

	ans.FirstAnswer = chatStreamToGin(c, req)
	ans.IsOk = true

	// Store verification evidence on decentralized Web3 storage
	// Prepare JSON data containing question, AI analysis and search evidence
	data := map[string]interface{}{
		"Question":    query.Question,
		"FirstAnswer": ans.FirstAnswer,
		"RagRaw":      ragRaw,
		"IsOk":        ans.IsOk,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Failed to marshal JSON data:", err)
		return
	}

	// save data to json
	filePath := fmt.Sprintf("/tmp/search_results_%s.json", ans.ID)
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		log.Println("Failed to write JSON data to file:", err)
		return
	}

	// us w3 up to upload file
	cmd := exec.Command("w3", "up", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Failed to upload file using w3 up:", err)
		log.Println("Output:", string(output))
		return
	}

	//
	uploadURL := extractUploadURL(string(output))
	if uploadURL == "" {
		log.Println("Failed to extract upload URL from output:", string(output))
		return
	}
	log.Println("File uploaded successfully. URL:", uploadURL)

	// ans.UploadURL = uploadURL
	webStorageInfo := fmt.Sprintf(`Web3 Storage URL:%s
`, uploadURL)

	ans.FirstAnswer += webStorageInfo

	c.String(200, webStorageInfo)

}

// extractUploadURL extracts the Filecoin/IPFS upload URL from w3 CLI output
// Args:
//   output: Raw output string from w3 up command
// Returns:
//   The extracted upload URL or empty string if not found
func extractUploadURL(output string) string {
	re := regexp.MustCompile(`https://w3s\.link/ipfs/[a-zA-Z0-9]+`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

// chatStreamToGin streams LLM responses back to client via Gin context
// Args:
//   c: Gin context for streaming response
//   req: OpenAI chat completion request
// Returns:
//   completeAns: Full aggregated response from LLM
func chatStreamToGin(c *gin.Context, req openai.ChatCompletionRequest) (completeAns string) {
	var resp *openai.ChatCompletionStream

	resp, err := shared.Cli.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		utils.GinErrorMsgTxt(c, "LLM backend broken")
		log.Println(err)
		return
	}
	defer resp.Close()

	sb := strings.Builder{}
	buf := strings.Builder{}

	c.Stream(func(w io.Writer) bool {
		msg, err := resp.Recv()
		if errors.Is(err, io.EOF) || err != nil {
			bufS := buf.String()
			if bufS != "" {
				w.Write([]byte(bufS))
				sb.WriteString(bufS)
			}
			return false
		}
		delta := msg.Choices[0].Delta.Content
		delta = utils.CleanStr(delta)
		sb.WriteString(delta)

		if utils.WriteSplitByRune(w, &buf, delta, '\n', '.', ';', '。', '？', '?') {
			return true
		}
		return true
	})
	completeAns = sb.String()
	return
}
