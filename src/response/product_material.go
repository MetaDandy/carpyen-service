package response

import (
	"time"

	"github.com/MetaDandy/carpyen-service/src/model"
)

type ProductMaterial struct {
	ID        string `json:"id"`
	Quantity  string `json:"quantity"`
	UnitPrice string `json:"unit_price"`
	TotalCost string `json:"total_cost"`

	Material Material `json:"material"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ProductMaterialToDto(m *model.ProductMaterial) ProductMaterial {
	return ProductMaterial{
		ID:        m.ID.String(),
		Quantity:  m.Quantity.String(),
		UnitPrice: m.UnitPrice.String(),
		TotalCost: m.TotalCost.String(),

		Material: MaterialToDto(&m.Material),

		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
	}
}

func ProductMaterialToListDto(m []model.ProductMaterial) []ProductMaterial {
	var dtos []ProductMaterial
	for _, pm := range m {
		dtos = append(dtos, ProductMaterialToDto(&pm))
	}
	return dtos
}
