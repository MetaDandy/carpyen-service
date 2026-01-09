package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type BatchMaterialSupplier struct {
	ID        uuid.UUID `json:"id"`
	Quantity  string    `json:"quantity"`
	UnitPrice string    `json:"unit_price"`
	TotalCost string    `json:"total_cost"`
	Stock     string    `json:"stock"`

	Material *Material `json:"material"`
	User     *User     `json:"user,omitzero"`
}

func BatchMaterialSupplierToDto(m *model.BatchMaterialSupplier) BatchMaterialSupplier {
	var dto BatchMaterialSupplier
	copier.Copy(&dto, m)

	dto.UnitPrice = m.UnitPrice.String()
	dto.TotalCost = m.TotalCost.String()
	dto.Stock = m.Stock.String()
	if m.User.ID != (uuid.UUID{}) {
		userDto := UserToDto(&m.User)
		dto.User = &userDto
	} else {
		dto.User = nil
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
