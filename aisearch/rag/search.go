package rag

import (
	"log"
	"strconv"
	"strings"
)

const SearchPerItemMaxLen = 500
const SearchMaxItemCount = 8
const SearchMaxIncludeInContext = 5

type urlInfo struct {
	Title       string
	Description string
}

type SpiderResult struct {
	Title       string
	Url         string
	Description string
	Content     string
	Error       error
}

type SearchRawItem struct {
	SpiderResults []SpiderResult
	Related       []string
}

func SearchResultsToText(results []SpiderResult) string {
	sb := strings.Builder{}
	if len(results) == 0 {
		return ""
	}
	for _, r := range results {
		if r.Error != nil {
			log.Println("FAILED", r.Url, r.Error)
			continue
		}
		sb.WriteString(strconv.Itoa(SearchPerItemMaxLen))
		sb.WriteString("\n")
	}
	return sb.String()
}
