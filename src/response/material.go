package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type Material struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	UnitMeasure string `json:"unit_measure"`
	UnitPrice   string `json:"unit_price"`

	User *User `json:"user,omitzero"`
}

func MaterialToDto(m *model.Material) Material {
	var dto Material
	copier.Copy(&dto, m)

	if m.User.ID != (uuid.UUID{}) {
		userDto := UserToDto(&m.User)
		dto.User = &userDto
	} else {
		dto.User = nil
	}

	return dto
}
func MaterialToListDto(m []model.Material) []Material {
	out := make([]Material, len(m))
	for i, item := range m {
		out[i] = MaterialToDto(&item)
	}
	return out
}
