package user

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	Create(m model.User) error
	FindByID(id string) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindAll(opts *helper.FindAllOptions) ([]model.User, int64, error)
	Update(m model.User) error
	SoftDelete(id string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) Create(m model.User) error {
	return r.db.Create(&m).Error
}

func (r *repo) FindByID(id string) (model.User, error) {
	var user model.User
	err := r.db.First(&user, "id = ?", id).Error
	return user, err
}

func (r *repo) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}

func (r *repo) FindAll(opts *helper.FindAllOptions) ([]model.User, int64, error) {
	var finded []model.User
	query := r.db.Model(model.User{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *repo) Update(m model.User) error {
	return r.db.Save(&m).Error
}

func (r *repo) SoftDelete(id string) error {
	return r.db.Delete(&model.User{}, "id = ?", id).Error
}
