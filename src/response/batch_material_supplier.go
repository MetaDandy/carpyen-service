package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
)

type BatchMaterialSupplier struct {
	ID        string `json:"id"`
	Quantity  string `json:"quantity"`
	UnitPrice string `json:"unit_price"`
	TotalCost string `json:"total_cost"`
	Stock     string `json:"stock"`

	Material *Material `json:"material"`
	Supplier *Supplier `json:"supplier"`
	User     *User     `json:"user,omitzero"`
}

func BatchMaterialSupplierToDto(m *model.BatchMaterialSupplier) BatchMaterialSupplier {
	dto := BatchMaterialSupplier{
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

	if m.Supplier.ID != (uuid.UUID{}) {
		sup := SupplierToDto(&m.Supplier)
		dto.Supplier = &sup
	}

	if m.User.ID != (uuid.UUID{}) {
		usr := UserToDto(&m.User)
		dto.User = &usr
	}

	return dto
}

func BatchMaterialSupplierToListDto(m []model.BatchMaterialSupplier) []BatchMaterialSupplier {
	out := make([]BatchMaterialSupplier, len(m))
	for i, item := range m {
		out[i] = BatchMaterialSupplierToDto(&item)
	}
	return out
}
