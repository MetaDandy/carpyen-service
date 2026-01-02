package client

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"gorm.io/gorm"
)

type Repo interface {
	Create(m model.Client) error
	FindByID(id string) (model.Client, error)
	FindByEmail(email string) (model.Client, error)
	FindAll(opts *helper.FindAllOptions) ([]model.Client, int64, error)
	Update(m model.Client) error
	SoftDelete(id string) error

	ValidateSeller(id string, idclient string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db: db}
}

func (r *repo) Create(m model.Client) error {
	return r.db.Create(&m).Error
}

func (r *repo) FindByID(id string) (model.Client, error) {
	var Client model.Client
	err := r.db.First(&Client, "id = ?", id).Error
	return Client, err
}

func (r *repo) FindByEmail(email string) (model.Client, error) {
	var Client model.Client
	err := r.db.First(&Client, "email = ?", email).Error
	return Client, err
}

func (r *repo) FindAll(opts *helper.FindAllOptions) ([]model.Client, int64, error) {
	var finded []model.Client
	query := r.db.Model(model.Client{})
	var total int64
	query, total = opts.ApplyFindAllOptions(query)

	err := query.Find(&finded).Error
	return finded, total, err
}

func (r *repo) Update(m model.Client) error {
	return r.db.Save(&m).Error
}

func (r *repo) SoftDelete(id string) error {
	return r.db.Delete(&model.Client{}, "id = ?", id).Error
}

func (r *repo) ValidateSeller(id string, idclient string) error {
	err := r.db.Model(&model.Client{}).Where("id = ? AND user_id = ?", idclient, id).Error
	if err != nil {
		return err
	}

	return nil
}
