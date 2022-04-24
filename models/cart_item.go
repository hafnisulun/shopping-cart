package models

import "github.com/gofrs/uuid"

type CartItem struct {
	Base
	CartID    uint    `json:"-" gorm:"not null"`
	Cart      Cart    `json:"-"`
	ProductID uint    `json:"-" gorm:"not null"`
	Product   Product `json:"product" gorm:"not null"`
	Qty       int     `json:"qty" gorm:"not null"`
}

type CartItemInput struct {
	ProductUUID uuid.UUID `json:"product_uuid" binding:"required"`
}
