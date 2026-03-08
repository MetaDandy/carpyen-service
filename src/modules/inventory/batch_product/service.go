package batchproduct

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
	FindByID(id string) (*response.BatchProduct, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchProduct], error)
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

type WarehouseRepo interface {
	FindByID(id string) (model.Warehouse, error)
}

type service struct {
	repo          Repo
	userRepo      UserRepo
	productRepo   ProductRepo
	warehouseRepo WarehouseRepo
}

func NewService(repo Repo, userRepo UserRepo, productRepo ProductRepo, warehouseRepo WarehouseRepo) Service {
	return &service{repo: repo, userRepo: userRepo, productRepo: productRepo, warehouseRepo: warehouseRepo}
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

	warehouse, err := s.warehouseRepo.FindByID(input.WarehouseID)
	if err != nil {
		return err
	}

	batchproduct := model.BatchProduct{}
	batchproduct.ID = uuid.New()
	batchproduct.UserID = user.ID
	batchproduct.ProductID = product.ID
	batchproduct.WarehouseID = warehouse.ID

	batchproduct.UnitPrice, err = decimal.NewFromString(input.UnitPrice)
	if err != nil {
		return errors.New("invalid unit price")
	}

	batchproduct.Quantity, err = decimal.NewFromString(input.Quantity)
	if err != nil {
		return errors.New("invalid quantity")
	}

	batchproduct.Stock = batchproduct.Quantity

	batchproduct.TotalPrice = batchproduct.Quantity.Mul(batchproduct.UnitPrice)

	return s.repo.create(batchproduct)
}

func (s *service) FindByID(id string) (*response.BatchProduct, error) {
	batchproduct, err := s.repo.findByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.BatchProductToDto(&batchproduct)
	return &dto, nil
}

func (s *service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchProduct], error) {
	finded, total, err := s.repo.findAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := response.BatchProductToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.BatchProduct]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}
func (s *service) Update(id string, input Update) error {
	batchproduct, err := s.repo.findByID(id)
	if err != nil {
		return err
	}

	if input.UnitPrice != nil {
		batchproduct.UnitPrice, err = decimal.NewFromString(*input.UnitPrice)
		if err != nil {
			return errors.New("invalid unit price")
		}
		batchproduct.TotalPrice = batchproduct.Quantity.Mul(batchproduct.UnitPrice)
	}

	if input.Quantity != nil && batchproduct.Stock.Equal(batchproduct.Quantity) {
		batchproduct.Quantity, err = decimal.NewFromString(*input.Quantity)
		if err != nil {
			return errors.New("invalid quantity")
		}
		batchproduct.Stock = batchproduct.Quantity
		batchproduct.TotalPrice = batchproduct.Quantity.Mul(batchproduct.UnitPrice)
	}

	if input.ProductID != nil && batchproduct.Stock.Equal(batchproduct.Quantity) {
		product, err := s.productRepo.FindByID(*input.ProductID)
		if err != nil {
			return err
		}
		batchproduct.ProductID = product.ID
		batchproduct.Product = product
	}

	if input.WarehouseID != nil && batchproduct.Stock.Equal(batchproduct.Quantity) {
		warehouse, err := s.warehouseRepo.FindByID(*input.WarehouseID)
		if err != nil {
			return err
		}
		batchproduct.WarehouseID = warehouse.ID
		batchproduct.Warehouse = warehouse
	}
	return s.repo.update(batchproduct)
}

func (s *service) SoftDelete(id string) error {
	return s.repo.softDelete(id)
}

func (s *service) ValidateInstaller(id string, iduser string) error {
	return s.repo.validateInstaller(id, iduser)
}
