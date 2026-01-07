package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	UnitPrice float64 `json:"unit_price"`

	User *User `json:"user,omitzero"`
}

func ProductToDto(m *model.Product) Product {
	var dto Product
	copier.Copy(&dto, m)

	if m.User.ID != (uuid.UUID{}) {
		userDto := UserToDto(&m.User)
		dto.User = &userDto
	} else {
		dto.User = nil
	}

	return dto
}
func ProductToListDto(m []model.Product) []Product {
	out := make([]Product, len(m))
	for i, item := range m {
		out[i] = ProductToDto(&item)
	}
	return out
}
