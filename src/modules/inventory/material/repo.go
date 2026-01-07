package material

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	create(m model.Material) error
	findByID(id string) (model.Material, error)
	findAll(opts *helper.FindAllOptions) ([]model.Material, int64, error)
	update(m model.Material) error
	softDelete(id string) error

	validateInstaller(id string, iduser string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) create(m model.Material) error {
	return r.db.Create(&m).Error
}

func (r *repo) findByID(id string) (model.Material, error) {
	var material model.Material
	err := r.db.Preload("User").First(&material, "id = ?", id).Error
	return material, err
}

func (r *repo) findAll(opts *helper.FindAllOptions) ([]model.Material, int64, error) {
	var finded []model.Material
	query := r.db.Model(model.Material{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *repo) update(m model.Material) error {
	return r.db.Save(&m).Error
}

func (r *repo) softDelete(id string) error {
	return r.db.Delete(&model.Material{}, "id = ?", id).Error
}

func (r *repo) validateInstaller(id string, iduser string) error {
	err := r.db.Model(&model.Material{}).
		Where("id = ? AND user_id = ?", id, iduser).Error
	if err != nil {
		return err
	}

	return nil
}
