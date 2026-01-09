package batchproductmaterial

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
	FindByID(id string) (*response.BatchProductMaterial, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchProductMaterial], error)
	Update(id string, input Update) error
	SoftDelete(id string) error

	ValidateInstaller(id string, iduser string) error
}

type UserRepo interface {
	FindByID(id string) (model.User, error)
}

type ProductRepo interface {
	FindByID(id string) (model.Product, error)
}

type service struct {
	repo        Repo
	userRepo    UserRepo
	productRepo ProductRepo
}

func NewService(repo Repo, userRepo UserRepo, productRepo ProductRepo) Service {
	return &service{repo: repo, userRepo: userRepo, productRepo: productRepo}
}
func (s *service) Create(input Create, userID string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	batchProductMaterial := model.BatchProductMaterial{}
	copier.Copy(&batchProductMaterial, &input)
	batchProductMaterial.ID = uuid.New()
	batchProductMaterial.UserID = user.ID

	batchProductMaterial.UnitPrice, err = decimal.NewFromString(input.UnitPrice)
	if err != nil {
		return errors.New("invalid unit price")
	}

	batchProductMaterial.Stock, err = decimal.NewFromString(input.Stock)
	if err != nil {
		return errors.New("invalid stock")
	}

	return s.repo.create(batchProductMaterial)
}

func (s *service) FindByID(id string) (*response.BatchProductMaterial, error) {
	batchProductMaterial, err := s.repo.findByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.BatchProductMaterialToDto(&batchProductMaterial)
	return &dto, nil
}

func (s *service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchProductMaterial], error) {
	batchProductMaterials, total, err := s.repo.findAll(opts)
	if err != nil {
		return nil, err
	}

	var dtos []response.BatchProductMaterial
	for _, batchProductMaterial := range batchProductMaterials {
		dto := response.BatchProductMaterialToDto(&batchProductMaterial)
		dtos = append(dtos, dto)
	}

	paginated := &response.Paginated[response.BatchProductMaterial]{
		Total: total,
		Data:  dtos,
	}
	return paginated, nil
}

// Actualizar el product id solo si no se toco el stock y esta igual al quantity
func (s *service) Update(id string, input Update) error {
	batchProductMaterial, err := s.repo.findByID(id)
	if err != nil {
		return err
	}
	if input.Quantity != nil {
		batchProductMaterial.Quantity = *input.Quantity
	}
	if input.UnitPrice != nil {
		batchProductMaterial.UnitPrice, err = decimal.NewFromString(*input.UnitPrice)
		if err != nil {
			return errors.New("invalid unit price")
		}
	}
	if input.Stock != nil {
		batchProductMaterial.Stock, err = decimal.NewFromString(*input.Stock)
		if err != nil {
			return errors.New("invalid stock")
		}
	}

	return s.repo.update(batchProductMaterial)
}

func (s *service) SoftDelete(id string) error {
	return s.repo.softDelete(id)
}

func (s *service) ValidateInstaller(id string, iduser string) error {
	return s.repo.validateInstaller(id, iduser)
}
