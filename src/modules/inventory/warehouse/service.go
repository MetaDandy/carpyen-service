package warehouse

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/MetaDandy/carpyen-service/src/response"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type Service interface {
	create(input Create, userID string) error
	findByID(id string) (*response.Warehouse, error)
	findAll(opts *helper.FindAllOptions) (*response.Paginated[response.Warehouse], error)
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

	warehouse := model.Warehouse{}
	copier.Copy(&warehouse, &input)
	warehouse.User = user
	warehouse.ID = uuid.New()
	warehouse.UserID = user.ID

	return s.repo.Create(warehouse)
}

func (s *service) findByID(id string) (*response.Warehouse, error) {
	warehouse, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.WarehouseToDto(&warehouse)
	return &dto, nil
}

func (s *service) findAll(opts *helper.FindAllOptions) (*response.Paginated[response.Warehouse], error) {
	finded, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := response.WarehouseToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.Warehouse]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}

func (s *service) update(id string, input Update) error {
	warehouse, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	copier.CopyWithOption(&warehouse, &input, copier.Option{IgnoreEmpty: true})

	return s.repo.Update(warehouse)
}

func (s *service) softDelete(id string) error {
	return s.repo.SoftDelete(id)
}

func (s *service) validateChiefInstaller(id string, iduser string) error {
	return s.repo.ValidateChiefInstaller(id, iduser)
}
