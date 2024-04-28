package model

type Exchange struct {
	BaseEntity
	Name             string `json:"name" validate:"required"`
	Description      string `json:"description"`
	Website          string `json:"website"`
	DocumentationUrl string `json:"documentationUrl"`
	ApiUrl           string `json:"apiUrl"`
	Status           string `json:"status"`
}

type ExchangeStatus string

const (
	ENABLE  ExchangeStatus = "Enable"
	DISABLE ExchangeStatus = "Disable"
)
