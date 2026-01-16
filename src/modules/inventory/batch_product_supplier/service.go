package batchproductsupplier

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
	FindByID(id string) (*response.BatchProductSupplier, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchProductSupplier], error)
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

type SupplierRepo interface {
	FindByID(id string) (model.Supplier, error)
}

type service struct {
	repo         Repo
	userRepo     UserRepo
	productRepo  ProductRepo
	supplierRepo SupplierRepo
}

func NewService(repo Repo, userRepo UserRepo, productRepo ProductRepo, supplierRepo SupplierRepo) Service {
	return &service{repo: repo, userRepo: userRepo, productRepo: productRepo, supplierRepo: supplierRepo}
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

	supplier, err := s.supplierRepo.FindByID(input.SupplierID)
	if err != nil {
		return err
	}

	batchproductsupplier := model.BatchProductSupplier{}
	batchproductsupplier.ID = uuid.New()
	batchproductsupplier.UserID = user.ID
	batchproductsupplier.ProductID = product.ID
	batchproductsupplier.SupplierID = supplier.ID

	batchproductsupplier.UnitPrice, err = decimal.NewFromString(input.UnitPrice)
	if err != nil {
		return errors.New("invalid unit price")
	}

	batchproductsupplier.Quantity, err = decimal.NewFromString(input.Quantity)
	if err != nil {
		return errors.New("invalid quantity")
	}

	batchproductsupplier.Stock = batchproductsupplier.Quantity

	batchproductsupplier.TotalPrice = batchproductsupplier.Quantity.Mul(batchproductsupplier.UnitPrice)

	return s.repo.create(batchproductsupplier)
}

func (s *service) FindByID(id string) (*response.BatchProductSupplier, error) {
	batchproductsupplier, err := s.repo.findByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.BatchProductSupplierToDto(&batchproductsupplier)
	return &dto, nil
}

func (s *service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchProductSupplier], error) {
	finded, total, err := s.repo.findAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := response.BatchProductSupplierToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.BatchProductSupplier]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}
func (s *service) Update(id string, input Update) error {
	batchproductsupplier, err := s.repo.findByID(id)
	if err != nil {
		return err
	}

	if input.UnitPrice != nil {
		batchproductsupplier.UnitPrice, err = decimal.NewFromString(*input.UnitPrice)
		if err != nil {
			return errors.New("invalid unit price")
		}
		batchproductsupplier.TotalPrice = batchproductsupplier.Quantity.Mul(batchproductsupplier.UnitPrice)
	}

	if input.Quantity != nil && batchproductsupplier.Stock.Equal(batchproductsupplier.Quantity) {
		batchproductsupplier.Quantity, err = decimal.NewFromString(*input.Quantity)
		if err != nil {
			return errors.New("invalid quantity")
		}
		batchproductsupplier.Stock = batchproductsupplier.Quantity
		batchproductsupplier.TotalPrice = batchproductsupplier.Quantity.Mul(batchproductsupplier.UnitPrice)
	}

	if input.ProductID != nil && batchproductsupplier.Stock.Equal(batchproductsupplier.Quantity) {
		product, err := s.productRepo.FindByID(*input.ProductID)
		if err != nil {
			return err
		}
		batchproductsupplier.ProductID = product.ID
		batchproductsupplier.Product = product
	}

	if input.SupplierID != nil && batchproductsupplier.Stock.Equal(batchproductsupplier.Quantity) {
		supplier, err := s.supplierRepo.FindByID(*input.SupplierID)
		if err != nil {
			return err
		}
		batchproductsupplier.SupplierID = supplier.ID
		batchproductsupplier.Supplier = supplier
	}
	return s.repo.update(batchproductsupplier)
}

func (s *service) SoftDelete(id string) error {
	return s.repo.softDelete(id)
}

func (s *service) ValidateInstaller(id string, iduser string) error {
	return s.repo.validateInstaller(id, iduser)
}
