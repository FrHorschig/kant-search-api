package models

type SearchCriteria struct {

	Options SearchOptions `json:"options"`

	SearchTerms string `json:"searchTerms"`
}
