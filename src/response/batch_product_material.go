package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
)

type BatchProductSupplier struct {
	ID         string `json:"id"`
	Quantity   string `json:"quantity"`
	UnitPrice  string `json:"unit_price"`
	TotalPrice string `json:"total_cost"`
	Stock      string `json:"stock"`

	Product  *Product  `json:"product"`
	Supplier *Supplier `json:"supplier"`
	User     *User     `json:"user,omitzero"`
}

func BatchProductSupplierToDto(m *model.BatchProductSupplier) BatchProductSupplier {
	dto := BatchProductSupplier{
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

func BatchProductSupplierToListDto(m []model.BatchProductSupplier) []BatchProductSupplier {
	out := make([]BatchProductSupplier, len(m))
	for i, item := range m {
		out[i] = BatchProductSupplierToDto(&item)
	}
	return out
}
