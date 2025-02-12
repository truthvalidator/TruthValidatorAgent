package handler

import (
	"ssaisearch/rag"
	"ssaisearch/rag/searxng"
)

// ConvertSearxNGResultsToSERPResults converts []searxng.SpiderResult to []serp.SpiderResult
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

// ConvertSearxNGRawItemToRagRawItem converts searxng.SearchRawItem to rag.SearchRawItem
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

func GenerateProgramGuideStr(progLang string) (string, error) {

	if progLang != "" {
		return progLang + ",Ensure to append the cited links at the end of your response.", nil
	}

	return "Ensure to append the cited links at the end of your response.", nil
}
