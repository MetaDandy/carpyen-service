package supplier

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/MetaDandy/carpyen-service/src/response"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type Service interface {
	create(input Create, userID string) error
	findByID(id string) (*response.Supplier, error)
	findAll(opts *helper.FindAllOptions) (*response.Paginated[response.Supplier], error)
	update(id string, input Update) error
	softDelete(id string) error

	validateChiefInstaller(id string, iduser string) error
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

func (s *service) create(input Create, userID string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	supplier := model.Supplier{}
	copier.Copy(&supplier, &input)
	supplier.User = user
	supplier.ID = uuid.New()
	supplier.UserID = user.ID

	return s.repo.Create(supplier)
}

func (s *service) findByID(id string) (*response.Supplier, error) {
	supplier, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.SupplierToDto(&supplier)
	return &dto, nil
}

func (s *service) findAll(opts *helper.FindAllOptions) (*response.Paginated[response.Supplier], error) {
	finded, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := response.SupplierToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.Supplier]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}

func (s *service) update(id string, input Update) error {
	supplier, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	copier.CopyWithOption(&supplier, &input, copier.Option{IgnoreEmpty: true})

	return s.repo.Update(supplier)
}

func (s *service) softDelete(id string) error {
	return s.repo.SoftDelete(id)
}

func (s *service) validateChiefInstaller(id string, iduser string) error {
	return s.repo.ValidateChiefInstaller(id, iduser)
}
