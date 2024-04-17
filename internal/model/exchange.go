package model

type Exchange struct {
	BaseEntity
	Name             string `json:"name" validate:"required"`
	Description      string `json:"description"`
	Website          string `json:"website"`
	DocumentationUrl string `json:"documentation_url"`
	ApiUrl           string `json:"api_url"`
	Status           string `json:"status"`
}

type ExchangeStatus string

const (
	ENABLE  ExchangeStatus = "Enable"
	DISABLE ExchangeStatus = "Disable"
)
