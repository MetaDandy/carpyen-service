package supplier

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	create(m model.Supplier) error
	findByID(id string) (model.Supplier, error)
	findAll(opts *helper.FindAllOptions) ([]model.Supplier, int64, error)
	update(m model.Supplier) error
	softDelete(id string) error

	validateChiefInstaller(id string, iduser string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) create(m model.Supplier) error {
	return r.db.Create(&m).Error
}

func (r *repo) findByID(id string) (model.Supplier, error) {
	var supplier model.Supplier
	err := r.db.Preload("User").First(&supplier, "id = ?", id).Error
	return supplier, err
}

func (r *repo) findAll(opts *helper.FindAllOptions) ([]model.Supplier, int64, error) {
	var finded []model.Supplier
	query := r.db.Model(model.Supplier{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *repo) update(m model.Supplier) error {
	return r.db.Save(&m).Error
}

func (r *repo) softDelete(id string) error {
	return r.db.Delete(&model.Supplier{}, "id = ?", id).Error
}

func (r *repo) validateChiefInstaller(id string, iduser string) error {
	err := r.db.Model(&model.Supplier{}).
		Where("id = ? AND user_id = ?", id, iduser).Error
	if err != nil {
		return err
	}

	return nil
}
