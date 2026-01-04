package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/jinzhu/copier"
)

type Material struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	UnitMeasure string `json:"unit_measure"`
	UnitPrice   string `json:"unit_price"`
}

func MaterialToDto(m *model.Material) Material {
	var dto Material
	copier.Copy(&dto, m)
	return dto
}
func MaterialToListDto(m []model.Material) []Material {
	out := make([]Material, len(m))
	for i, item := range m {
		out[i] = MaterialToDto(&item)
	}
	return out
}
