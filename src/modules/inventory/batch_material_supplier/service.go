package batchmaterialsupplier

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
	FindByID(id string) (*response.BatchMaterialSupplier, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.BatchMaterialSupplier], error)
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

type SupplierRepo interface {
	FindByID(id string) (model.Supplier, error)
}

type service struct {
	repo         Repo
	userRepo     UserRepo
	materialRepo MaterialRepo
	supplierRepo SupplierRepo
}

func NewService(repo Repo, userRepo UserRepo, materialRepo MaterialRepo, supplierRepo SupplierRepo) Service {
	return &service{repo: repo, userRepo: userRepo, materialRepo: materialRepo, supplierRepo: supplierRepo}
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

	supplier, err := s.supplierRepo.FindByID(input.SupplierID)
	if err != nil {
		return err
	}

	batchmaterialsupplier := model.BatchMaterialSupplier{}
	batchmaterialsupplier.ID = uuid.New()
	batchmaterialsupplier.UserID = user.ID
	batchmaterialsupplier.MaterialID = material.ID
	batchmaterialsupplier.SupplierID = supplier.ID

	batchmaterialsupplier.UnitPrice, err = decimal.NewFromString(input.UnitPrice)
	if err != nil {
		return errors.New("invalid unit price")
	}

	batchmaterialsupplier.Quantity, err = decimal.NewFromString(input.Quantity)
	if err != nil {
		return errors.New("invalid quantity")
	}

	batchmaterialsupplier.Stock = batchmaterialsupplier.Quantity

	batchmaterialsupplier.TotalCost = batchmaterialsupplier.Quantity.Mul(batchmaterialsupplier.UnitPrice)

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

	if input.UnitPrice != nil {
		batchmaterialsupplier.UnitPrice, err = decimal.NewFromString(*input.UnitPrice)
		if err != nil {
			return errors.New("invalid unit price")
		}
		batchmaterialsupplier.TotalCost = batchmaterialsupplier.Quantity.Mul(batchmaterialsupplier.UnitPrice)
	}

	if input.Quantity != nil && batchmaterialsupplier.Stock.Equal(batchmaterialsupplier.Quantity) {
		batchmaterialsupplier.Quantity, err = decimal.NewFromString(*input.Quantity)
		if err != nil {
			return errors.New("invalid quantity")
		}
		batchmaterialsupplier.Stock = batchmaterialsupplier.Quantity
		batchmaterialsupplier.TotalCost = batchmaterialsupplier.Quantity.Mul(batchmaterialsupplier.UnitPrice)
	}

	if input.MaterialID != nil && batchmaterialsupplier.Stock.Equal(batchmaterialsupplier.Quantity) {
		material, err := s.materialRepo.FindByID(*input.MaterialID)
		if err != nil {
			return err
		}
		batchmaterialsupplier.MaterialID = material.ID
		batchmaterialsupplier.Material = material
	}

	if input.SupplierID != nil && batchmaterialsupplier.Stock.Equal(batchmaterialsupplier.Quantity) {
		supplier, err := s.supplierRepo.FindByID(*input.SupplierID)
		if err != nil {
			return err
		}
		batchmaterialsupplier.SupplierID = supplier.ID
		batchmaterialsupplier.Supplier = supplier
	}
	return s.repo.update(batchmaterialsupplier)
}

func (s *service) SoftDelete(id string) error {
	return s.repo.softDelete(id)
}

func (s *service) ValidateInstaller(id string, iduser string) error {
	return s.repo.validateInstaller(id, iduser)
}
