package models

type HttpError struct {

	Code int32 `json:"code"`

	Message ErrorMessage `json:"message"`

	Params []string `json:"params,omitempty"`
}
