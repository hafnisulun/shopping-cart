package models

type BuyAGetBPromo struct {
	Base
	PromoID      uint    `json:"promo_id" gorm:"not null"`
	Promo        Promo   `json:"promo" gorm:"not null"`
	BuyProductID uint    `json:"buy_product_id" gorm:"not null"`
	BuyProduct   Product `json:"buy_product" gorm:"not null"`
	BuyQty       int     `json:"buy_qty" gorm:"not null"`
	GetProductID uint    `json:"get_product_id" gorm:"not null"`
	GetProduct   Product `json:"get_product" gorm:"not null"`
	GetQty       int     `json:"get_qty" gorm:"not null"`
}
