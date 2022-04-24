package models

import "database/sql/driver"

type PromoType string

const (
	PromoTypeBuyAGetB PromoType = "buy_a_get_b"
	PromoTypeMinQty   PromoType = "min_qty"
)

func (g *PromoType) Scan(value interface{}) error {
	*g = PromoType(value.(string))
	return nil
}

func (g PromoType) Value() (driver.Value, error) {
	return string(g), nil
}

type Promo struct {
	Base
	Type PromoType `json:"type"  gorm:"type:promo_type not null"`
}
