package purchase

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/MetaDandy/carpyen-service/src/response"
	"github.com/google/uuid"
)

type Service interface {
	Create(input Create, userID string) error
	FindByID(id string) (*response.Purchase, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.Purchase], error)
	Update(id string, input Update) error
	SoftDelete(id string) error

	ValidateInstaller(id string, iduser string) error
}
type UserRepo interface {
	FindByID(id string) (model.User, error)
}

type SupplierRepo interface {
	FindByID(id string) (model.Supplier, error)
}

type service struct {
	repo         Repo
	userRepo     UserRepo
	supplierRepo SupplierRepo
}

func NewService(repo Repo, userRepo UserRepo, supplierRepo SupplierRepo) Service {
	return &service{repo: repo, userRepo: userRepo, supplierRepo: supplierRepo}
}

func (s *service) Create(input Create, userID string) error {

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	supplier, err := s.supplierRepo.FindByID(input.SupplierID)
	if err != nil {
		return err
	}

	purchase := model.Purchase{}
	purchase.ID = uuid.New()
	purchase.UserID = user.ID
	purchase.SupplierID = supplier.ID

	return s.repo.create(purchase)
}

func (s *service) FindByID(id string) (*response.Purchase, error) {
	purchase, err := s.repo.findByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.PurchaseToDto(&purchase)
	return &dto, nil
}

func (s *service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.Purchase], error) {
	finded, total, err := s.repo.findAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := response.PurchaseToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.Purchase]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}
func (s *service) Update(id string, input Update) error {
	purchase, err := s.repo.findByID(id)
	if err != nil {
		return err
	}

	if input.SupplierID != nil {
		supplier, err := s.supplierRepo.FindByID(*input.SupplierID)
		if err != nil {
			return err
		}
		purchase.SupplierID = supplier.ID
		purchase.Supplier = supplier
	}
	return s.repo.update(purchase)
}

func (s *service) SoftDelete(id string) error {
	return s.repo.softDelete(id)
}

func (s *service) ValidateInstaller(id string, iduser string) error {
	return s.repo.validateInstaller(id, iduser)
}
