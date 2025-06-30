package models

type Volume struct {

	Title string `json:"title"`

	VolumeNumber int32 `json:"volumeNumber"`

	Works []Work `json:"works"`
}
