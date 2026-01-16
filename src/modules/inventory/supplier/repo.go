package supplier

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	Create(m model.Supplier) error
	FindByID(id string) (model.Supplier, error)
	FindAll(opts *helper.FindAllOptions) ([]model.Supplier, int64, error)
	Update(m model.Supplier) error
	SoftDelete(id string) error

	ValidateChiefInstaller(id string, iduser string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) Create(m model.Supplier) error {
	return r.db.Create(&m).Error
}

func (r *repo) FindByID(id string) (model.Supplier, error) {
	var supplier model.Supplier
	err := r.db.Preload("User").First(&supplier, "id = ?", id).Error
	return supplier, err
}

func (r *repo) FindAll(opts *helper.FindAllOptions) ([]model.Supplier, int64, error) {
	var finded []model.Supplier
	query := r.db.Model(model.Supplier{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *repo) Update(m model.Supplier) error {
	return r.db.Save(&m).Error
}

func (r *repo) SoftDelete(id string) error {
	return r.db.Delete(&model.Supplier{}, "id = ?", id).Error
}

func (r *repo) ValidateChiefInstaller(id string, iduser string) error {
	var supplier model.Supplier
	return r.db.
		Where("id = ? AND user_id = ?", id, iduser).
		First(&supplier).
		Error
}
