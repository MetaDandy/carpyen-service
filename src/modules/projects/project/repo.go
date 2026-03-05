package project

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	Create(m model.Project) error
	FindByID(id string) (model.Project, error)
	FindAll(opts *helper.FindAllOptions) ([]model.Project, int64, error)
	Update(m model.Project) error
	SoftDelete(id string) error
	ValidateSeller(id string, iduser string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) Create(m model.Project) error {

	return r.db.Create(&m).Error
}

func (r *repo) FindByID(id string) (model.Project, error) {
	var project model.Project
	err := r.db.Preload("User").Preload("Client").First(&project, "id = ?", id).Error
	return project, err
}

func (r *repo) FindAll(opts *helper.FindAllOptions) ([]model.Project, int64, error) {
	var finded []model.Project
	query := r.db.Model(model.Project{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *repo) Update(m model.Project) error {
	return r.db.Save(&m).Error
}

func (r *repo) SoftDelete(id string) error {
	return r.db.Delete(&model.Project{}, "id = ?", id).Error
}

func (r *repo) ValidateSeller(id string, iduser string) error {
	var project model.Project
	return r.db.
		Where("id = ? AND user_id = ?", id, iduser).
		First(&project).
		Error
}
