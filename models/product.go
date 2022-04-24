package models

type Product struct {
	Base
	SKU   string  `json:"sku" gorm:"not null"`
	Name  string  `json:"name" gorm:"not null"`
	Price float64 `json:"price" gorm:"not null"`
	Qty   int     `json:"qty" gorm:"not null"`
}
