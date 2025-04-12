package handler

import (
	"ssaisearch/rag"
	"ssaisearch/rag/searxng"
)

// ConvertSearxNGResultsToSERPResults converts search results from SearxNG format to RAG format
// Args:
//   searxngResults: Raw search results from SearxNG metasearch engine
// Returns:
//   []rag.SpiderResult: Converted results in RAG-compatible format
func ConvertSearxNGResultsToSERPResults(searxngResults []searxng.SpiderResult) []rag.SpiderResult {
	serpResults := make([]rag.SpiderResult, len(searxngResults))
	for i, result := range searxngResults {
		serpResults[i] = rag.SpiderResult{
			Title:       result.Title,
			Url:         result.URL,
			Description: result.Description,
			Content:     "",  // Assuming Content is not available in searxng.SpiderResult
			Error:       nil, // Assuming Error is not available in searxng.SpiderResult
		}
	}
	return serpResults
}

// ConvertSearxNGRawItemToRagRawItem converts complete search response from SearxNG to RAG format
// Args:
//   searxngRaw: Complete raw search response from SearxNG
// Returns:
//   rag.SearchRawItem: Converted results including both search items and related queries
func ConvertSearxNGRawItemToRagRawItem(searxngRaw searxng.SearchRawItem) rag.SearchRawItem {
	return rag.SearchRawItem{
		SpiderResults: ConvertSearxNGResultsToSERPResults(searxngRaw.SpiderResults),
		Related:       searxngRaw.Related,
	}
}

type ProgramGuide struct {
	ProgramLanguage string `json:"program_language"`
	Guide           string `json:"guide"`
}

// GenerateProgramGuideStr creates language-specific guidance for LLM responses
// Args:
//   programmingLanguage: Target programming language (e.g. "golang")
// Returns:
//   string: Formatted guidance string for LLM prompt
//   error: Always nil in current implementation
func GenerateProgramGuideStr(programmingLanguage string) (string, error) {
	if programmingLanguage != "" {
		return programmingLanguage + ",Ensure to append the cited links at the end of your response.", nil
	}
	return "Ensure to append the cited links at the end of your response.", nil
}
