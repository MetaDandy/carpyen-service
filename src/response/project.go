package response

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	State       string `json:"state"`

	User   *User   `json:"user,omitzero"`
	Client *Client `json:"client"`
}

func ProjectToDto(m *model.Project) Project {
	var dto Project
	copier.Copy(&dto, m)

	if m.User.ID != (uuid.UUID{}) {
		userDto := UserToDto(&m.User)
		dto.User = &userDto
	} else {
		dto.User = nil
	}
	if m.Client.ID != (uuid.UUID{}) {
		clientDto := ClientToDto(&m.Client)
		dto.Client = &clientDto
	} else {
		dto.Client = nil
	}

	return dto
}
func ProjectToListDto(m []model.Project) []Project {
	out := make([]Project, len(m))
	for i, item := range m {
		out[i] = ProjectToDto(&item)
	}
	return out
}
