package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/jinzhu/copier"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Role    string `json:"role"`
}

func UserToDto(m *model.User) User {
	var dto User
	copier.Copy(&dto, m)

	dto.Role = m.Role.String()

	return dto
}

func UserToListDto(m []model.User) []User {
	out := make([]User, len(m))
	for i := range m {
		out[i] = UserToDto(&m[i])
	}
	return out
}
