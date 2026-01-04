package material

import "github.com/MetaDandy/carpyen-service/src/enum"

type Create struct {
	Name        string        `json:"name" validate:"required"`
	Type        enum.Material `json:"type" validate:"required"`
	UnitMeasure enum.Measure  `json:"unit_measure" validate:"required"`
	UnitPrice   float64       `json:"unit_price" validate:"required,gt=0"`
}

type Update struct {
	Name        *string        `json:"name"`
	Type        *enum.Material `json:"type"`
	UnitMeasure *enum.Measure  `json:"unit_measure"`
	UnitPrice   *float64       `json:"unit_price"`
}
