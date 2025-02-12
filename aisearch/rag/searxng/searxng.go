package searxng

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type SearxNGConfig struct {
	Endpoint string
}

type SearxNGResponse struct {
	Results []SearxNGResult `json:"results"`
}

type SearxNGResult struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"content"`
}

func SearchSearxNG(config SearxNGConfig, query, language string) (SearxNGResponse, error) {
	params := url.Values{}
	params.Add("format", "json")
	params.Add("categories", "general")
	params.Add("safesearch", "1")
	params.Add("language", language)
	params.Add("q", query)

	url := config.Endpoint + "?" + params.Encode()

	resp, err := http.Get(url)
	if err != nil {
		return SearxNGResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearxNGResponse{}, err
	}

	var searxNGResponse SearxNGResponse
	err = json.Unmarshal(body, &searxNGResponse)
	if err != nil {
		return SearxNGResponse{}, err
	}

	return searxNGResponse, nil
}

type SearchRawItem struct {
	Related       []string       `json:"related"`
	SpiderResults []SpiderResult `json:"spider_results"`
}

type SpiderResult struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

func SearchRawFromSearxNG(config SearxNGConfig, country, locale, query string) (SearchRawItem, error) {
	searxNGResponse, err := SearchSearxNG(config, query, locale)
	if err != nil {
		return SearchRawItem{}, err
	}

	var searchRaw SearchRawItem
	for i, result := range searxNGResponse.Results {
		if i >= 13 {
			break
		}
		searchRaw.SpiderResults = append(searchRaw.SpiderResults, SpiderResult{
			Title:       result.Title,
			URL:         result.URL,
			Description: result.Description,
		})
	}

	return searchRaw, nil
}
