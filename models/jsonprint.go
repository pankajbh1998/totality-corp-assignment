package models

type JsonPrint struct {
	Code    int         `json:"Code"`
	Message interface{} `json:"message"`
}
