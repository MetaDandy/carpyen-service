package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type Supplier struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
	User    User   `json:"user"`
}

func SupplierToDto(m *model.Supplier) Supplier {
	var dto Supplier

	copier.Copy(&dto, m)

	if m.User.ID != (uuid.UUID{}) {
		dto.User = UserToDto(&m.User)
	}

	return dto
}

func SupplierToListDto(m []model.Supplier) []Supplier {
	out := make([]Supplier, len(m))
	for i := range m {
		out[i] = SupplierToDto(&m[i])
	}
	return out
}
