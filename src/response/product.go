package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/jinzhu/copier"
)

type Product struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	UnitPrice string `json:"unit_price"`
}

func ProductToDto(m *model.Product) Product {
	var dto Product
	copier.Copy(&dto, m)
	return dto
}
func ProductToListDto(m []model.Product) []Product {
	out := make([]Product, len(m))
	for i, item := range m {
		out[i] = ProductToDto(&item)
	}
	return out
}
