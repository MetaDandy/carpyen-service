package batchmaterialsupplier

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	create(m model.BatchMaterialSupplier) error
	findByID(id string) (model.BatchMaterialSupplier, error)
	findAll(opts *helper.FindAllOptions) ([]model.BatchMaterialSupplier, int64, error)
	update(m model.BatchMaterialSupplier) error
	softDelete(id string) error

	validateInstaller(id string, iduser string) error
}
type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) create(m model.BatchMaterialSupplier) error {
	return r.db.Create(&m).Error
}
func (r *repo) findByID(id string) (model.BatchMaterialSupplier, error) {
	var batchMaterialSupplier model.BatchMaterialSupplier
	err := r.db.Preload("User").Preload("Material").Preload("Supplier").First(&batchMaterialSupplier, "id = ?", id).Error
	return batchMaterialSupplier, err
}

func (r *repo) findAll(opts *helper.FindAllOptions) ([]model.BatchMaterialSupplier, int64, error) {
	var finded []model.BatchMaterialSupplier

	query := r.db.Model(model.BatchMaterialSupplier{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Preload("Supplier").Preload("Material").Find(&finded).Error
	return finded, total, err
}

func (r *repo) update(m model.BatchMaterialSupplier) error {
	return r.db.Save(&m).Error
}

func (r *repo) softDelete(id string) error {
	return r.db.Delete(&model.Material{}, "id = ?", id).Error
}

func (r *repo) validateInstaller(id string, iduser string) error {
	var batchMaterialSupplier model.BatchMaterialSupplier
	return r.db.
		Where("id = ? AND user_id = ?", id, iduser).
		First(&batchMaterialSupplier).
		Error
}
