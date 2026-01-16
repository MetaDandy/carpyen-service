package product

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	Create(m model.Product) error
	FindByID(id string) (model.Product, error)
	FindAll(opts *helper.FindAllOptions) ([]model.Product, int64, error)
	Update(m model.Product) error
	SoftDelete(id string) error
	ValidateInstaller(id string, iduser string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) Create(m model.Product) error {
	return r.db.Create(&m).Error
}

func (r *repo) FindByID(id string) (model.Product, error) {
	var product model.Product
	err := r.db.Preload("User").First(&product, "id = ?", id).Error
	return product, err
}

func (r *repo) FindAll(opts *helper.FindAllOptions) ([]model.Product, int64, error) {
	var finded []model.Product
	query := r.db.Model(model.Product{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *repo) Update(m model.Product) error {
	return r.db.Save(&m).Error
}

func (r *repo) SoftDelete(id string) error {
	return r.db.Delete(&model.Product{}, "id = ?", id).Error
}

func (r *repo) ValidateInstaller(id string, iduser string) error {
	var product model.Product
	return r.db.
		Where("id = ? AND user_id = ?", id, iduser).
		First(&product).
		Error
}
