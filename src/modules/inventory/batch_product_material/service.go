package batchproductmaterial

import (
	"errors"

	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/MetaDandy/carpyen-service/src/response"
	"github.com/google/uuid"
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

	product, err := s.productRepo.FindByID(input.ProductID)
	if err != nil {
		return err
	}

	batchProductMaterial := model.BatchProductMaterial{}
	batchProductMaterial.ID = uuid.New()
	batchProductMaterial.UserID = user.ID
	batchProductMaterial.ProductID = product.ID

	batchProductMaterial.UnitPrice, err = decimal.NewFromString(input.UnitPrice)
	if err != nil {
		return errors.New("invalid unit price")
	}

	batchProductMaterial.Quantity, err = decimal.NewFromString(input.Quantity)
	if err != nil {
		return errors.New("invalid quantity")
	}

	batchProductMaterial.Stock = batchProductMaterial.Quantity

	batchProductMaterial.TotalCost = batchProductMaterial.Quantity.Mul(batchProductMaterial.UnitPrice)

	return s.repo.create(batchProductMaterial)
}

func (s *service) FindByID(id string) (*response.BatchProductMaterial, error) {
	batchProductMaterial, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.BatchProductMaterialToDto(&batchProductMaterial)
	return &dto, nil
}

func (s *service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchProductMaterial], error) {
	finded, total, err := s.repo.findAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := response.BatchProductMaterialToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.BatchProductMaterial]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}

func (s *service) Update(id string, input Update) error {
	batchProductMaterial, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if input.UnitPrice != nil {
		batchProductMaterial.UnitPrice, err = decimal.NewFromString(*input.UnitPrice)
		if err != nil {
			return errors.New("invalid unit price")
		}
		batchProductMaterial.TotalCost = batchProductMaterial.Quantity.Mul(batchProductMaterial.UnitPrice)
	}

	if input.Quantity != nil && batchProductMaterial.Stock.Equal(batchProductMaterial.Quantity) {
		batchProductMaterial.Quantity, err = decimal.NewFromString(*input.Quantity)
		if err != nil {
			return errors.New("invalid quantity")
		}
		batchProductMaterial.Stock = batchProductMaterial.Quantity
		batchProductMaterial.TotalCost = batchProductMaterial.Quantity.Mul(batchProductMaterial.UnitPrice)
	}

	if input.ProductID != nil && batchProductMaterial.Stock.Equal(batchProductMaterial.Quantity) {
		product, err := s.productRepo.FindByID(*input.ProductID)
		if err != nil {
			return err
		}
		batchProductMaterial.ProductID = product.ID
		batchProductMaterial.Product = product
	}

	return s.repo.update(batchProductMaterial)
}

func (s *service) SoftDelete(id string) error {
	return s.repo.softDelete(id)
}

func (s *service) ValidateInstaller(id string, iduser string) error {
	return s.repo.validateInstaller(id, iduser)
}
