package models

type Paragraph struct {

	FnRefs []string `json:"fnRefs,omitempty"`

	Ordinal int32 `json:"ordinal"`

	Text string `json:"text"`

	SummaryRef string `json:"summaryRef,omitempty"`
}
