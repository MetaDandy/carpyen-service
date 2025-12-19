package user

import (
	"github.com/MetaDandy/go-fiber-skeleton/helper"
	"github.com/MetaDandy/go-fiber-skeleton/src/model"
)

type UserRepo interface {
	Create(m model.User) error
	FindByID(id string) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindAll(opts *helper.FindAllOptions) ([]model.User, int64, error)
	Update(m model.User) error
	SoftDelete(id string) error
}
