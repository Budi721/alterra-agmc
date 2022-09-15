package models

// Response struct model in order to consistent structure when response api
type Response struct {
	Status string `json:"status,omitempty"`
	Code   uint   `json:"code,omitempty"`
	Data   any    `json:"data,omitempty"`
}
