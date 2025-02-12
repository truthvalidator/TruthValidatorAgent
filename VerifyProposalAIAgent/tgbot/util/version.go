package util

import (
	"time"
)

var (
	Version = "0.0.1"
	Hash    = "hash"
	BuildAt = "2006-01-02 00:00:00"
)

type Release struct {
	TagName     string     `json:"tag_name"`
	Name        string     `json:"name"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	Body        string     `json:"body"`
}
