package models

import (
	"gorm.io/gorm"
)

type Operation struct {
	gorm.Model
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"result"`
	Message   string  `json:"message"`
}

func NewOperation(operation string, a, b float64, msg string) *Operation {
	return &Operation{Operation: operation, A: a, B: b, Message: msg}
}
