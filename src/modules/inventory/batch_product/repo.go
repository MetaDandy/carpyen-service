package batchproduct

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	create(m model.BatchProduct) error
	findByID(id string) (model.BatchProduct, error)
	findAll(opts *helper.FindAllOptions) ([]model.BatchProduct, int64, error)
	update(m model.BatchProduct) error
	softDelete(id string) error

	validateInstaller(id string, iduser string) error
}
type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) create(m model.BatchProduct) error {
	return r.db.Create(&m).Error
}
func (r *repo) findByID(id string) (model.BatchProduct, error) {
	var batchProduct model.BatchProduct
	err := r.db.Preload("User").Preload("Product").Preload("Supplier").First(&batchProduct, "id = ?", id).Error
	return batchProduct, err
}

func (r *repo) findAll(opts *helper.FindAllOptions) ([]model.BatchProduct, int64, error) {
	var finded []model.BatchProduct

	query := r.db.Model(model.BatchProduct{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Preload("Supplier").Preload("Product").Find(&finded).Error
	return finded, total, err
}

func (r *repo) update(m model.BatchProduct) error {
	return r.db.Save(&m).Error
}

func (r *repo) softDelete(id string) error {
	return r.db.Delete(&model.Product{}, "id = ?", id).Error
}

func (r *repo) validateInstaller(id string, iduser string) error {
	var batchProduct model.BatchProduct
	return r.db.
		Where("id = ? AND user_id = ?", id, iduser).
		First(&batchProduct).
		Error
}
