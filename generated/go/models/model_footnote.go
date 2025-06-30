package models

type Footnote struct {

	Ordinal int32 `json:"ordinal"`

	Ref string `json:"ref"`

	Text string `json:"text"`
}
