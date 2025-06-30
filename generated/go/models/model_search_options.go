package models

type SearchOptions struct {

	IncludeHeadings bool `json:"includeHeadings"`

	IncludeFootnotes bool `json:"includeFootnotes"`

	IncludeParagraphs bool `json:"includeParagraphs"`

	WithStemming bool `json:"withStemming"`

	WorkCodes []string `json:"workCodes"`
}
