package productmaterial

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	Create(m model.ProductMaterial) error
	FindByID(id string) (model.ProductMaterial, error)
	FindAll(id string, opts *helper.FindAllOptions) ([]model.ProductMaterial, int64, error)
	Update(m model.ProductMaterial) error
	Delete(id string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) Create(m model.ProductMaterial) error {
	return r.db.Create(&m).Error
}

func (r *repo) FindByID(id string) (model.ProductMaterial, error) {
	var finded model.ProductMaterial
	err := r.db.First(&finded, "id = ?", id).Error
	return finded, err
}

func (r *repo) FindAll(id string, opts *helper.FindAllOptions) ([]model.ProductMaterial, int64, error) {
	var finded []model.ProductMaterial
	query := r.db.Model(model.ProductMaterial{}).Where("batch_product_material_id = ?", id)
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.
		Preload("Material").
		Find(&finded).Error
	return finded, total, err
}

func (r *repo) Update(m model.ProductMaterial) error {
	return r.db.Save(&m).Error
}

func (r *repo) Delete(id string) error {
	return r.db.Delete(&model.ProductMaterial{}, "id = ?", id).Unscoped().Error
}
