package jooble

import "encoding/json"

type SearchResponse struct {
	TotalCount    json.Number   `json:"total_count" example:"1473"`
	Opportunities []Opportunity `json:"jobs"`
}
