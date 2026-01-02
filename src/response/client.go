package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type Client struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
}

func ClientToDto(m *model.Client) Client {
	var dto Client
	copier.Copy(&dto, m)

	return dto
}
func ClientToListDto(m []model.Client) []Client {
	out := make([]Client, len(m))
	for i := range m {
		out[i] = ClientToDto(&m[i])
	}
	return out
}
