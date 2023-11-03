package models

type Work struct {

	Id int32 `json:"id"`

	Title string `json:"title"`

	Abbreviation string `json:"abbreviation,omitempty"`

	Ordinal int32 `json:"ordinal"`

	Year string `json:"year,omitempty"`

	VolumeId int32 `json:"volumeId"`
}
