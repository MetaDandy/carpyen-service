package batchmaterial

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
	FindByID(id string) (*response.BatchMaterial, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchMaterial], error)
	Update(id string, input Update) error
	SoftDelete(id string) error

	ValidateInstaller(id string, iduser string) error
}
type UserRepo interface {
	FindByID(id string) (model.User, error)
}

type MaterialRepo interface {
	FindByID(id string) (model.Material, error)
}

type WarehouseRepo interface {
	FindByID(id string) (model.Warehouse, error)
}

type service struct {
	repo          Repo
	userRepo      UserRepo
	materialRepo  MaterialRepo
	warehouseRepo WarehouseRepo
}

func NewService(repo Repo, userRepo UserRepo, materialRepo MaterialRepo, warehouseRepo WarehouseRepo) Service {
	return &service{repo: repo, userRepo: userRepo, materialRepo: materialRepo, warehouseRepo: warehouseRepo}
}

func (s *service) Create(input Create, userID string) error {

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	material, err := s.materialRepo.FindByID(input.MaterialID)
	if err != nil {
		return err
	}

	warehouse, err := s.warehouseRepo.FindByID(input.WarehouseID)
	if err != nil {
		return err
	}

	batchmaterial := model.BatchMaterial{}
	batchmaterial.ID = uuid.New()
	batchmaterial.UserID = user.ID
	batchmaterial.MaterialID = material.ID
	batchmaterial.WarehouseID = warehouse.ID

	batchmaterial.UnitPrice, err = decimal.NewFromString(input.UnitPrice)
	if err != nil {
		return errors.New("invalid unit price")
	}

	batchmaterial.Quantity, err = decimal.NewFromString(input.Quantity)
	if err != nil {
		return errors.New("invalid quantity")
	}

	batchmaterial.Stock = batchmaterial.Quantity

	batchmaterial.TotalCost = batchmaterial.Quantity.Mul(batchmaterial.UnitPrice)

	return s.repo.create(batchmaterial)
}

func (s *service) FindByID(id string) (*response.BatchMaterial, error) {
	batchmaterial, err := s.repo.findByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.BatchMaterialToDto(&batchmaterial)
	return &dto, nil
}

func (s *service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchMaterial], error) {
	finded, total, err := s.repo.findAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := response.BatchMaterialToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.BatchMaterial]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}
func (s *service) Update(id string, input Update) error {
	batchmaterial, err := s.repo.findByID(id)
	if err != nil {
		return err
	}

	if input.UnitPrice != nil {
		batchmaterial.UnitPrice, err = decimal.NewFromString(*input.UnitPrice)
		if err != nil {
			return errors.New("invalid unit price")
		}
		batchmaterial.TotalCost = batchmaterial.Quantity.Mul(batchmaterial.UnitPrice)
	}

	if input.Quantity != nil && batchmaterial.Stock.Equal(batchmaterial.Quantity) {
		batchmaterial.Quantity, err = decimal.NewFromString(*input.Quantity)
		if err != nil {
			return errors.New("invalid quantity")
		}
		batchmaterial.Stock = batchmaterial.Quantity
		batchmaterial.TotalCost = batchmaterial.Quantity.Mul(batchmaterial.UnitPrice)
	}

	if input.MaterialID != nil && batchmaterial.Stock.Equal(batchmaterial.Quantity) {
		material, err := s.materialRepo.FindByID(*input.MaterialID)
		if err != nil {
			return err
		}
		batchmaterial.MaterialID = material.ID
		batchmaterial.Material = material
	}

	if input.WarehouseID != nil && batchmaterial.Stock.Equal(batchmaterial.Quantity) {
		warehouse, err := s.warehouseRepo.FindByID(*input.WarehouseID)
		if err != nil {
			return err
		}
		batchmaterial.WarehouseID = warehouse.ID
		batchmaterial.Warehouse = warehouse
	}
	return s.repo.update(batchmaterial)
}

func (s *service) SoftDelete(id string) error {
	return s.repo.softDelete(id)
}

func (s *service) ValidateInstaller(id string, iduser string) error {
	return s.repo.validateInstaller(id, iduser)
}
