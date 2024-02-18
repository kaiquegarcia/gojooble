package jooble

import (
	"encoding/json"
	"net/http"
)

type Jooble interface {
	Search(keywords string, location string) (*SearchResponse, Audit, error)
}

// New returns an implementation of Jooble interface.
func New(
	httpClient *http.Client,
	apiKey string,
) Jooble {
	return &jooble{
		httpClient: httpClient,
		baseURL:    "https://jooble.org/api",
		apiKey:     apiKey,
	}
}

type jooble struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

func (j *jooble) Search(
	keywords string,
	location string,
) (*SearchResponse, Audit, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"keywords": keywords,
		"location": location,
	})
	if err != nil {
		return nil, &audit{}, err
	}

	var searchResponse SearchResponse
	a, err := j.do(
		http.MethodPost,
		j.baseURL+"/"+j.apiKey,
		payload,
		&searchResponse,
	)
	if err != nil {
		return nil, a, err
	}

	return &searchResponse, a, err
}
