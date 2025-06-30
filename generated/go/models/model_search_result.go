package models

type SearchResult struct {

	Hits []Hit `json:"hits"`

	WorkCode string `json:"workCode"`
}
