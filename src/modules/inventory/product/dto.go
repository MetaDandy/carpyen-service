package product

import "github.com/MetaDandy/carpyen-service/src/enum"

type Create struct {
	Name      string       `json:"name" validate:"required"`
	Type      enum.Product `json:"type" validate:"required"`
	UnitPrice float64      `json:"unit_price" validate:"required,gt=0"`
}

type Update struct {
	Name      *string       `json:"name"`
	Type      *enum.Product `json:"type"`
	UnitPrice *float64      `json:"unit_price"`
}
