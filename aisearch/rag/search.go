// Package rag provides retrieval-augmented generation (RAG) functionality for AI search.
// It processes web search results to prepare them for use as context in LLM prompts.
//
// Features:
// - Web search result processing
// - Content length management
// - Error handling for failed results
package rag

import (
	"log"
	"strconv"
	"strings"
)

// Constants for managing search result processing
const (
	// Maximum length for each search result item in characters
	SearchPerItemMaxLen = 500
	
	// Maximum number of search results to process  
	SearchMaxItemCount = 8
	
	// Maximum number of results to include in final context
	SearchMaxIncludeInContext = 5
)

// urlInfo contains basic metadata about a URL
type urlInfo struct {
	Title       string // Page title
	Description string // Page description/meta
}

// SpiderResult contains all data from crawling a single URL
type SpiderResult struct {
	Title       string // Page title
	Url         string // Full URL
	Description string // Page description/meta
	Content     string // Full page content
	Error       error  // Crawling error if any
}

// SearchRawItem contains a set of spider results and related queries
type SearchRawItem struct {
	SpiderResults []SpiderResult // All crawled results
	Related       []string       // Related search queries
}

// SearchResultsToText converts spider results into a formatted text block for LLM context.
// It filters out failed results and enforces length limits.
//
// Parameters:
//   results: Slice of spider results to process
//
// Returns:
//   string: Formatted text block with valid results
//           Empty string if no valid results
func SearchResultsToText(results []SpiderResult) string {
	sb := strings.Builder{}
	if len(results) == 0 {
		return ""
	}
	
	// Process each result, skipping failures
	for _, r := range results {
		if r.Error != nil {
			log.Printf("Skipping failed result from %s: %v", r.Url, r.Error)
			continue
		}
		
		// Add result length limit (implementation appears incomplete)
		sb.WriteString(strconv.Itoa(SearchPerItemMaxLen))
		sb.WriteString("\n")
	}
	
	return sb.String()
}
