package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
)

type MaterialDetails struct {
	ID            string `json:"id"`

	BatchMaterial *BatchMaterial `json:"batch_material"`
	User            *User            `json:"user,omitzero"`
}

func MaterialDetailsToDto(m *model.MaterialDetails) MaterialDetails {
	dto := MaterialDetails{
		ID:            m.ID.String(),
	}

	if m.BatchMaterial.ID != (uuid.UUID{}) {
		batch := BatchMaterialToDto(&m.BatchMaterial)
		dto.BatchMaterial = &batch
	}

	if m.User.ID != (uuid.UUID{}) {
		usr := UserToDto(&m.User)
		dto.User = &usr
	}

	return dto
}

func MaterialDetailsToListDto(m []model.MaterialDetails) []MaterialDetails {
	out := make([]MaterialDetails, len(m))
	for i, item := range m {
		out[i] = MaterialDetailsToDto(&item)
	}
	return out
}
