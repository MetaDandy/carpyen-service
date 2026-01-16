package productmaterial

import (
	"fmt"

	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/MetaDandy/carpyen-service/src/response"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Service interface {
	Create(input Create) error
	FindAll(id string, opts *helper.FindAllOptions) (*response.Paginated[response.ProductMaterial], error)
	Update(id string, input Update) error
	Delete(id string) error
}

type materialRepo interface {
	FindByID(id string) (model.Material, error)
}

type bpmRepo interface {
	FindByID(id string) (model.BatchProductMaterial, error)
}

type service struct {
	repo    Repo
	bpmRepo bpmRepo
	mRepo   materialRepo
}

func NewService(repo Repo, bpmRepo bpmRepo, mRepo materialRepo) Service {
	return &service{repo: repo, bpmRepo: bpmRepo, mRepo: mRepo}
}

func (s *service) Create(input Create) error {
	material, err := s.mRepo.FindByID(input.MaterialID)
	if err != nil {
		return err
	}

	bpm, err := s.bpmRepo.FindByID(input.BatchProductMaterialID)
	if err != nil {
		return err
	}

	var productMaterial model.ProductMaterial
	productMaterial.BatchProductMaterialID = bpm.ID
	productMaterial.MaterialID = material.ID
	productMaterial.ID = uuid.New()
	productMaterial.Quantity, err = decimal.NewFromString(input.Quantity)
	if err != nil {
		return fmt.Errorf("invalid quantity: %v", err)
	}

	productMaterial.UnitPrice = material.UnitPrice
	productMaterial.TotalCost = productMaterial.UnitPrice.Mul(productMaterial.Quantity)

	return s.repo.Create(productMaterial)
}

func (s *service) FindAll(id string, opts *helper.FindAllOptions) (*response.Paginated[response.ProductMaterial], error) {
	finded, total, err := s.repo.FindAll(id, opts)
	if err != nil {
		return nil, err
	}

	dtos := response.ProductMaterialToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.ProductMaterial]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}

func (s *service) Update(id string, input Update) error {
	productMaterial, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if input.Quantity != nil {
		productMaterial.Quantity, err = decimal.NewFromString(*input.Quantity)
		if err != nil {
			return fmt.Errorf("invalid quantity: %v", err)
		}
		productMaterial.TotalCost = productMaterial.UnitPrice.Mul(productMaterial.Quantity)
	}

	if input.UnitPrice != nil {
		productMaterial.UnitPrice, err = decimal.NewFromString(*input.UnitPrice)
		if err != nil {
			return fmt.Errorf("invalid unit price: %v", err)
		}
		productMaterial.TotalCost = productMaterial.UnitPrice.Mul(productMaterial.Quantity)
	}

	if input.MaterialID != nil {
		material, err := s.mRepo.FindByID(*input.MaterialID)
		if err != nil {
			return err
		}
		productMaterial.MaterialID = material.ID
		productMaterial.Material = material
	}

	return s.repo.Update(productMaterial)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
