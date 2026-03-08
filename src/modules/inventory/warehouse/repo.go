package warehouse

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	Create(m model.Warehouse) error
	FindByID(id string) (model.Warehouse, error)
	FindAll(opts *helper.FindAllOptions) ([]model.Warehouse, int64, error)
	Update(m model.Warehouse) error
	SoftDelete(id string) error

	ValidateChiefInstaller(id string, iduser string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) Create(m model.Warehouse) error {
	return r.db.Create(&m).Error
}

func (r *repo) FindByID(id string) (model.Warehouse, error) {
	var warehouse model.Warehouse
	err := r.db.Preload("User").First(&warehouse, "id = ?", id).Error
	return warehouse, err
}

func (r *repo) FindAll(opts *helper.FindAllOptions) ([]model.Warehouse, int64, error) {
	var finded []model.Warehouse
	query := r.db.Model(model.Warehouse{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *repo) Update(m model.Warehouse) error {
	return r.db.Save(&m).Error
}

func (r *repo) SoftDelete(id string) error {
	return r.db.Delete(&model.Warehouse{}, "id = ?", id).Error
}

func (r *repo) ValidateChiefInstaller(id string, iduser string) error {
	var warehouse model.Warehouse
	return r.db.
		Where("id = ? AND user_id = ?", id, iduser).
		First(&warehouse).
		Error
}
