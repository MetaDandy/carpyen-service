package product

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	create(m model.Product) error
	findByID(id string) (model.Product, error)
	findAll(opts *helper.FindAllOptions) ([]model.Product, int64, error)
	update(m model.Product) error
	softDelete(id string) error

	validateChiefInstaller(installerID string, productID string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) create(m model.Product) error {
	return r.db.Create(&m).Error
}

func (r *repo) findByID(id string) (model.Product, error) {
	var product model.Product
	err := r.db.Preload("BatchMaterialSuppliers").First(&product, "id = ?", id).Error
	return product, err
}

func (r *repo) findAll(opts *helper.FindAllOptions) ([]model.Product, int64, error) {
	var finded []model.Product
	query := r.db.Model(model.Product{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *repo) update(m model.Product) error {
	return r.db.Save(&m).Error
}

func (r *repo) softDelete(id string) error {
	return r.db.Delete(&model.Product{}, "id = ?", id).Error
}

func (r *repo) validateChiefInstaller(id string, iduser string) error {
	err := r.db.Model(&model.Product{}).
		Where("id = ? AND user_id = ?", id, iduser).Error
	if err != nil {
		return err
	}

	return nil
}
