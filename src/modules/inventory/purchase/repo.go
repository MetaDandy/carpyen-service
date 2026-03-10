package purchase

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	create(m model.Purchase) error
	findByID(id string) (model.Purchase, error)
	findAll(opts *helper.FindAllOptions) ([]model.Purchase, int64, error)
	update(m model.Purchase) error
	softDelete(id string) error

	validateInstaller(id string, iduser string) error
}
type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) create(m model.Purchase) error {
	return r.db.Create(&m).Error
}
func (r *repo) findByID(id string) (model.Purchase, error) {
	var purchase model.Purchase
	err := r.db.Preload("User").Preload("Material").Preload("Warehouse").First(&purchase, "id = ?", id).Error
	return purchase, err
}

func (r *repo) findAll(opts *helper.FindAllOptions) ([]model.Purchase, int64, error) {
	var finded []model.Purchase

	query := r.db.Model(model.Purchase{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Preload("Warehouse").Preload("Material").Find(&finded).Error
	return finded, total, err
}

func (r *repo) update(m model.Purchase) error {
	return r.db.Save(&m).Error
}

func (r *repo) softDelete(id string) error {
	return r.db.Delete(&model.Material{}, "id = ?", id).Error
}

func (r *repo) validateInstaller(id string, iduser string) error {
	var purchase model.Purchase
	return r.db.
		Where("id = ? AND user_id = ?", id, iduser).
		First(&purchase).
		Error
}
