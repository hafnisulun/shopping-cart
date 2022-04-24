package models

type Cart struct {
	Base
	Items []CartItem `json:"items"`
	Total float64    `json:"total" gorm:"-"`
}

type CartResponse struct {
	Data Cart `json:"data"`
}
