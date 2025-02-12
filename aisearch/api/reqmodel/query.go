package reqmodel

import (
	"strings"
)

type Query struct {
	Question string `json:"question"`
	ProgLang string `json:"prog_lang"`
	Language string `json:"language"`
	Field    string `json:"field"`
}

func (q Query) Regularize() Query {
	switch q.Language {
	default:
		q.Language = "English"
	}
	switch q.Field {
	case "code":
	default:
		q.Field = "code"
	}

	switch q.ProgLang {
	case "Go", "go":
		q.ProgLang = "golang"

	}
	q.Question = strings.TrimSpace(q.Question)
	return q
}

func (q Query) Locale() string {
	return "en"
}

func (q Query) Country() string {
	return "us"
}

type ContinuousAskQuery struct {
	Query
	ParentId string `json:"parent_id"`
}
