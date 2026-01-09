package batchmaterialsupplier

import (
	"errors"

	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/MetaDandy/carpyen-service/src/response"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/shopspring/decimal"
)

type Service interface {
	Create(input Create, userID string) error
	FindByID(id string) (*response.BatchMaterialSupplier, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchMaterialSupplier], error)
	Update(id string, input Update) error
	SoftDelete(id string) error

	ValidateInstaller(id string, iduser string) error
}
type UserRepo interface {
	FindByID(id string) (model.User, error)
}

type service struct {
	repo     Repo
	userRepo UserRepo
}

func NewService(repo Repo, userRepo UserRepo) Service {
	return &service{repo: repo, userRepo: userRepo}
}

func (s *service) Create(input Create, userID string) error {

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	batchmaterialsupplier := model.BatchMaterialSupplier{}
	copier.Copy(&batchmaterialsupplier, &input)
	batchmaterialsupplier.ID = uuid.New()
	batchmaterialsupplier.UserID = user.ID

	batchmaterialsupplier.UnitPrice, err = decimal.NewFromString(input.UnitPrice)
	if err != nil {
		return errors.New("invalid unit price")
	}

	return s.repo.create(batchmaterialsupplier)
}

func (s *service) FindByID(id string) (*response.BatchMaterialSupplier, error) {
	batchmaterialsupplier, err := s.repo.findByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.BatchMaterialSupplierToDto(&batchmaterialsupplier)
	return &dto, nil
}

func (s *service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchMaterialSupplier], error) {
	finded, total, err := s.repo.findAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := response.BatchMaterialSupplierToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.BatchMaterialSupplier]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}
func (s *service) Update(id string, input Update) error {
	batchmaterialsupplier, err := s.repo.findByID(id)
	if err != nil {
		return err
	}

	copier.CopyWithOption(&batchmaterialsupplier, &input, copier.Option{IgnoreEmpty: true})

	if input.UnitPrice != nil {
		batchmaterialsupplier.UnitPrice, err = decimal.NewFromString(*input.UnitPrice)
		if err != nil {
			return errors.New("invalid unit price")
		}
	}

	return s.repo.update(batchmaterialsupplier)
}

func (s *service) SoftDelete(id string) error {
	return s.repo.softDelete(id)
}

func (s *service) ValidateInstaller(id string, iduser string) error {
	return s.repo.validateInstaller(id, iduser)
}
