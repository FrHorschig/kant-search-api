package models

type Work struct {

	Siglum string `json:"siglum,omitempty"`

	Code string `json:"code"`

	Paragraphs []int32 `json:"paragraphs,omitempty"`

	Sections []Section `json:"sections"`

	Ordinal int32 `json:"ordinal"`

	Title string `json:"title"`

	Year string `json:"year"`
}
