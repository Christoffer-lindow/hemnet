package models

import "time"

type Property struct {
	Id               int        `json:"id"`
	Url              string     `json:"url"`
	ImgUrl           string     `json:"img_url"`
	Street           string     `json:"street"`
	Address          string     `json:"address"`
	Price            float32    `json:"price"`
	Type             string     `json:"type"`
	ContractType     string     `json:"contract_type"`
	Rooms            *int       `json:"rooms"`
	Area             *float32   `json:"area"`
	Balcony          *bool      `json:"balcony"`
	YearBuilt        *time.Time `json:"year_built"`
	Rent             *float32   `json:"rent"`
	ConsumptionPrice *float32   `json:"consumption_price"`
	PriceSqm         *float32   `json:"price_sqm"`
}
