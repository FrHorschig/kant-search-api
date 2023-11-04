package models

type SearchCriteria struct {

	WorkIds []int32 `json:"workIds"`

	SearchString string `json:"searchString"`

	Options SearchOptions `json:"options"`
}
