package product

import (
	"github.com/MetaDandy/carpyen-service/src/enum"
)

type Create struct {
	Name      string       `json:"name" validate:"required"`
	Type      enum.Product `json:"type" validate:"required"`
	UnitPrice string       `json:"unit_price"`
}

type Update struct {
	Name      *string       `json:"name"`
	Type      *enum.Product `json:"type"`
	UnitPrice *string       `json:"unit_price"`
}
