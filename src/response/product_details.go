package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
)

type ProductDetails struct {
	ID string `json:"id"`

	BatchProduct *BatchProduct `json:"batch_product"`
	User         *User         `json:"user,omitzero"`
}

func ProductDetailsToDto(m *model.ProductDetails) ProductDetails {
	dto := ProductDetails{
		ID: m.ID.String(),
	}

	if m.BatchProduct.ID != (uuid.UUID{}) {
		batch := BatchProductToDto(&m.BatchProduct)
		dto.BatchProduct = &batch
	}

	if m.User.ID != (uuid.UUID{}) {
		usr := UserToDto(&m.User)
		dto.User = &usr
	}

	return dto
}

func ProductDetailsToListDto(m []model.ProductDetails) []ProductDetails {
	out := make([]ProductDetails, len(m))
	for i, item := range m {
		out[i] = ProductDetailsToDto(&item)
	}
	return out
}
