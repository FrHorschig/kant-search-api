package models

type Heading struct {

	FnRefs []string `json:"fnRefs,omitempty"`

	Ordinal int32 `json:"ordinal"`

	Pages []int32 `json:"pages"`

	Text string `json:"text"`

	TocText string `json:"tocText"`
}
