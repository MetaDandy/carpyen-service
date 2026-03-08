package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type Warehouse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Phone       string `json:"phone"`

	User *User `json:"user,omitzero"`
}

func WarehouseToDto(w *model.Warehouse) Warehouse {
	var dto Warehouse
	copier.Copy(&dto, w)

	if w.User.ID != (uuid.UUID{}) {
		userDto := UserToDto(&w.User)
		dto.User = &userDto
	} else {
		dto.User = nil
	}

	return dto
}
func WarehouseToListDto(w []model.Warehouse) []Warehouse {
	out := make([]Warehouse, len(w))
	for i, item := range w {
		out[i] = WarehouseToDto(&item)
	}
	return out
}
