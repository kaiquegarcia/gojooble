package jooble

import (
	"encoding/json"
)

type Opportunity struct {
	ID        json.Number `json:"id" example:"7506016653773926861"`
	Title     string      `json:"title" example:"Golang Developer"`
	Location  string      `json:"location" example:"Virginia"`
	Snippet   string      `json:"snippet"`
	Salary    string      `json:"salary"`
	Source    string      `json:"source" example:"app.linkedin.com"`
	Type      string      `json:"type"`
	Link      string      `json:"link" example:"https://jooble.org/desc/123?..."`
	Company   string      `json:"company"`
	UpdatedAt string   `json:"updated" example:"2024-02-16T00:00:00.0000000"`
}
