package models

type Match struct {

	Snippet string `json:"snippet"`

	Text string `json:"text"`

	Pages []int32 `json:"pages"`

	ParagraphId int32 `json:"paragraphId"`

	SentenceId int32 `json:"sentenceId,omitempty"`
}
