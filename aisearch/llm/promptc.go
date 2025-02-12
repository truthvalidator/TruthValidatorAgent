package llm

import (
	"fmt"
	"log"
	"strings"

	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/promptc/promptc-go/prompt"
)

func systemPrompt(ptsName Field, varMap map[string]string) string {
	compiled := _pts[ptsName].CompileWithOption(varMap, false)
	return strings.TrimSpace(compiled.Prompts[0].Prompt)
}

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

type Field string

const (
	SearchField Field = "search"
)

var _pts map[Field]*prompt.PromptC

func init() {
	_pts = map[Field]*prompt.PromptC{
		SearchField: loadPromptc("search.promptc"),
	}
}

func loadPromptc(path string) *prompt.PromptC {
	pt, err := iox.ReadAllText(path)
	if err != nil {
		log.Println("Failed to load promptc", err)
		panic(err)
	}
	return prompt.ParsePromptC(pt)
}
