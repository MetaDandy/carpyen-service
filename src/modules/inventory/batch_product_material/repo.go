package batchproductmaterial

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	create(m model.BatchProductMaterial) error
	FindByID(id string) (model.BatchProductMaterial, error)
	findAll(opts *helper.FindAllOptions) ([]model.BatchProductMaterial, int64, error)
	update(m model.BatchProductMaterial) error
	softDelete(id string) error

	validateInstaller(id string, iduser string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) create(m model.BatchProductMaterial) error {
	return r.db.Create(&m).Error
}

func (r *repo) FindByID(id string) (model.BatchProductMaterial, error) {
	var batchProductMaterial model.BatchProductMaterial
	err := r.db.
		Preload("Product").
		Preload("User").
		Preload("ProductMaterials").
		First(&batchProductMaterial, "id = ?", id).Error
	return batchProductMaterial, err
}

func (r *repo) findAll(opts *helper.FindAllOptions) ([]model.BatchProductMaterial, int64, error) {
	var finded []model.BatchProductMaterial
	query := r.db.Model(model.BatchProductMaterial{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.
		Preload("Product").
		Find(&finded).Error
	return finded, total, err
}

func (r *repo) update(m model.BatchProductMaterial) error {
	return r.db.Save(&m).Error
}

func (r *repo) softDelete(id string) error {
	return r.db.Delete(&model.BatchProductMaterial{}, "id = ?", id).Error
}

func (r *repo) validateInstaller(id string, iduser string) error {
	var material model.Material
	return r.db.
		Where("id = ? AND user_id = ?", id, iduser).
		First(&material).
		Error
}
