package models

type Paragraph struct {

	Id int32 `json:"id"`

	Text string `json:"text"`

	Pages []int32 `json:"pages"`

	WorkId int32 `json:"workId"`

	HeadingLevel int32 `json:"headingLevel,omitempty"`

	FootnoteName string `json:"footnoteName,omitempty"`
}
