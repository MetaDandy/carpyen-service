package material

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	Create(m model.Material) error
	FindByID(id string) (model.Material, error)
	FindAll(opts *helper.FindAllOptions) ([]model.Material, int64, error)
	Update(m model.Material) error
	SoftDelete(id string) error
	ValidateInstaller(id string, iduser string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) Create(m model.Material) error {
	return r.db.Create(&m).Error
}

func (r *repo) FindByID(id string) (model.Material, error) {
	var material model.Material
	err := r.db.Preload("User").First(&material, "id = ?", id).Error
	return material, err
}

func (r *repo) FindAll(opts *helper.FindAllOptions) ([]model.Material, int64, error) {
	var finded []model.Material
	query := r.db.Model(model.Material{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *repo) Update(m model.Material) error {
	return r.db.Save(&m).Error
}

func (r *repo) SoftDelete(id string) error {
	return r.db.Delete(&model.Material{}, "id = ?", id).Error
}

func (r *repo) ValidateInstaller(id string, iduser string) error {
	var material model.Material
	return r.db.
		Where("id = ? AND user_id = ?", id, iduser).
		First(&material).
		Error
}
