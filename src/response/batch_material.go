package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
)

type BatchMaterial struct {
	ID        string `json:"id"`
	Quantity  string `json:"quantity"`
	UnitPrice string `json:"unit_price"`
	TotalCost string `json:"total_cost"`
	Stock     string `json:"stock"`

	Material  *Material  `json:"material"`
	Warehouse *Warehouse `json:"warehouse"`
	User      *User      `json:"user,omitzero"`
}

func BatchMaterialToDto(m *model.BatchMaterial) BatchMaterial {
	dto := BatchMaterial{
		ID:        m.ID.String(),
		Quantity:  m.Quantity.String(),
		UnitPrice: m.UnitPrice.String(),
		TotalCost: m.TotalCost.String(),
		Stock:     m.Stock.String(),
	}

	if m.Material.ID != (uuid.UUID{}) {
		mat := MaterialToDto(&m.Material)
		dto.Material = &mat
	}

	if m.Warehouse.ID != (uuid.UUID{}) {
		wh := WarehouseToDto(&m.Warehouse)
		dto.Warehouse = &wh
	}

	if m.User.ID != (uuid.UUID{}) {
		usr := UserToDto(&m.User)
		dto.User = &usr
	}

	return dto
}

func BatchMaterialToListDto(m []model.BatchMaterial) []BatchMaterial {
	out := make([]BatchMaterial, len(m))
	for i, item := range m {
		out[i] = BatchMaterialToDto(&item)
	}
	return out
}
