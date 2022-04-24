package models

type MinQtyPromo struct {
	Base
	PromoID   uint    `json:"promo_id" gorm:"not null"`
	Promo     Promo   `json:"promo" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Product   Product `json:"product" gorm:"not null"`
	Qty       int     `json:"qty" gorm:"not null"`
	Discount  int     `json:"discount" gorm:"not null"`
}
