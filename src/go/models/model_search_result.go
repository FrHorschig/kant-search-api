package models

type SearchResult struct {

	WorkId int32 `json:"workId"`

	Matches []Match `json:"matches"`
}
