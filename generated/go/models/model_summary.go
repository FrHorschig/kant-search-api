package models

type Summary struct {

	FnRefs []string `json:"fnRefs,omitempty"`

	Ordinal int32 `json:"ordinal"`

	Ref string `json:"ref"`

	Text string `json:"text"`
}
