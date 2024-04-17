package model

type Currency struct {
	BaseEntity
	Name   string `json:"name"`
	Symbol string `json:"symbol" validate:"required"`
}
