// Package llm provides prompt management and generation for large language models (LLMs).
// It uses promptc templates to dynamically generate prompts based on context.
//
// Features:
// - Dynamic prompt generation from templates
// - Support for multiple languages and contexts
// - Integration with promptc template system
//
// See https://github.com/promptc/promptc-go for template format details.
package llm

import (
	"fmt"
	"log"
	"strings"

	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/prompt"
)

// promptBaseDir defines where prompt template files are located
const promptBaseDir = "prompts/"

// systemPrompt compiles a prompt template with the given variables
func systemPrompt(ptsName Field, varMap map[string]string) string {
	if _pts[ptsName] == nil {
		log.Printf("Warning: No prompt template found for field %s", ptsName)
		return ""
	}
	compiled := _pts[ptsName].CompileWithOption(varMap, false)
	return strings.TrimSpace(compiled.Prompts[0].Prompt)
}

// Promptc generates a complete LLM prompt based on the given parameters.
//
// Parameters:
//   - lang: Language code (e.g. "en", "zh")
//   - field: Prompt field type (currently only supports "search")
//   - question: The user's query/question
//   - guide: Additional guidance for the response format
//   - context: Any relevant context for the query
//
// Example:
//   prompt := Promptc("en", "search", "What is AI?", "technical",
//       "User is a computer science student")
func Promptc(lang string, field string, question string, guide string, context any) string {
	ptsField := SearchField
	switch field {
	default:
		ptsField = SearchField
		if guide != "" {
			guide = fmt.Sprintf(" When providing answers, please consider the context and format the response appropriately for a %s search engine.", guide)
		} else {
			guide = " Please stand in the perspective of a search engine. When providing answers, please consider the context and format the response appropriately."
		}
	}
	
	varMap := map[string]string{
		"lang":     lang,
		"guide":    guide,
		"question": question,
		"context":  fmt.Sprint(context),
	}
	
	return systemPrompt(ptsField, varMap)
}

// Field represents different types of prompt templates available
type Field string

const (
	// SearchField generates prompts for search-related queries
	SearchField Field = "search"
)

// _pts holds loaded prompt templates
var _pts map[Field]*prompt.PromptC

func init() {
	_pts = make(map[Field]*prompt.PromptC)
	
	// Load all prompt templates during initialization
	_pts[SearchField] = loadPromptc(promptBaseDir + "search.promptc")
}

// loadPromptc loads and parses a promptc template file
//
// Parameters:
//   - path: Path to the .promptc template file
//
// Returns:
//   - *prompt.PromptC: Parsed template
//   - Panics if template cannot be loaded
func loadPromptc(path string) *prompt.PromptC {
	pt, err := iox.ReadAllText(path)
	if err != nil {
		log.Printf("Failed to load prompt template from %s: %v", path, err)
		panic(fmt.Sprintf("Critical error: Could not load required prompt template %s", path))
	}
	
	parsed := prompt.ParsePromptC(pt)
	if parsed == nil {
		log.Printf("Failed to parse prompt template %s", path)
		panic(fmt.Sprintf("Critical error: Could not parse prompt template %s", path))
	}
	
	return parsed
}
