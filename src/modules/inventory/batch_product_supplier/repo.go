package batchproductsupplier

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	create(m model.BatchProductSupplier) error
	findByID(id string) (model.BatchProductSupplier, error)
	findAll(opts *helper.FindAllOptions) ([]model.BatchProductSupplier, int64, error)
	update(m model.BatchProductSupplier) error
	softDelete(id string) error

	validateInstaller(id string, iduser string) error
}
type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) create(m model.BatchProductSupplier) error {
	return r.db.Create(&m).Error
}
func (r *repo) findByID(id string) (model.BatchProductSupplier, error) {
	var batchProductSupplier model.BatchProductSupplier
	err := r.db.Preload("User").Preload("Product").Preload("Supplier").First(&batchProductSupplier, "id = ?", id).Error
	return batchProductSupplier, err
}

func (r *repo) findAll(opts *helper.FindAllOptions) ([]model.BatchProductSupplier, int64, error) {
	var finded []model.BatchProductSupplier

	query := r.db.Model(model.BatchProductSupplier{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Preload("Supplier").Preload("Product").Find(&finded).Error
	return finded, total, err
}

func (r *repo) update(m model.BatchProductSupplier) error {
	return r.db.Save(&m).Error
}

func (r *repo) softDelete(id string) error {
	return r.db.Delete(&model.Product{}, "id = ?", id).Error
}

func (r *repo) validateInstaller(id string, iduser string) error {
	var batchProductSupplier model.BatchProductSupplier
	return r.db.
		Where("id = ? AND user_id = ?", id, iduser).
		First(&batchProductSupplier).
		Error
}
