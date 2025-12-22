package user

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(m model.User) error
	FindByID(id string) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindAll(opts *helper.FindAllOptions) ([]model.User, int64, error)
	Update(m model.User) error
	SoftDelete(id string) error
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) UserRepo {
	return &Repo{db: db}
}

func (r *Repo) Create(m model.User) error {
	return r.db.Create(&m).Error
}

func (r *Repo) FindByID(id string) (model.User, error) {
	var user model.User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *Repo) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}

func (r *Repo) FindAll(opts *helper.FindAllOptions) ([]model.User, int64, error) {
	var finded []model.User
	query := r.db.Model(model.User{})
	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *Repo) Update(m model.User) error {
	return r.db.Save(&m).Error
}

func (r *Repo) SoftDelete(id string) error {
	return r.db.Delete(&model.User{}, "id = ?", id).Error
}
