package models

type Section struct {

	Heading int32 `json:"heading"`

	Paragraphs []int32 `json:"paragraphs,omitempty"`

	Sections []Section `json:"sections,omitempty"`
}
