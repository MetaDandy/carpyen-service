package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
)

type BatchProduct struct {
	ID         string `json:"id"`
	Quantity   string `json:"quantity"`
	UnitPrice  string `json:"unit_price"`
	TotalPrice string `json:"total_cost"`
	Stock      string `json:"stock"`

	Product   *Product   `json:"product"`
	Warehouse *Warehouse `json:"warehouse"`
	User      *User      `json:"user,omitzero"`
}

func BatchProductToDto(m *model.BatchProduct) BatchProduct {
	dto := BatchProduct{
		ID:         m.ID.String(),
		Quantity:   m.Quantity.String(),
		UnitPrice:  m.UnitPrice.String(),
		TotalPrice: m.TotalPrice.String(),
		Stock:      m.Stock.String(),
	}

	if m.Product.ID != (uuid.UUID{}) {
		mat := ProductToDto(&m.Product)
		dto.Product = &mat
	}

	if m.Warehouse.ID != (uuid.UUID{}) {
		warehouse := WarehouseToDto(&m.Warehouse)
		dto.Warehouse = &warehouse
	}

	if m.User.ID != (uuid.UUID{}) {
		usr := UserToDto(&m.User)
		dto.User = &usr
	}

	return dto
}

func BatchProductToListDto(m []model.BatchProduct) []BatchProduct {
	out := make([]BatchProduct, len(m))
	for i, item := range m {
		out[i] = BatchProductToDto(&item)
	}
	return out
}
