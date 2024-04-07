// models/response.go
package models

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Errors []string    `json:"errors,omitempty"`
}
