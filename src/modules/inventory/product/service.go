package product

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
	FindByID(id string) (*response.Product, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.Product], error)
	Update(id string, input Update) error
	SoftDelete(id string) error

	ValidateInstaller(id string, iduser string) error
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

func (s *service) Create(input Create, userID string) error {

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	if !input.Type.IsValid() {
		return errors.New("invalid product type")
	}

	product := model.Product{}
	copier.Copy(&product, &input)
	product.UnitPrice, err = decimal.NewFromString(input.UnitPrice)
	if err != nil {
		return errors.New("invalid unit price")
	}

	product.ID = uuid.New()
	product.UserID = user.ID

	return s.repo.create(product)
}

func (s *service) FindByID(id string) (*response.Product, error) {
	product, err := s.repo.findByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.ProductToDto(&product)
	return &dto, nil
}

func (s *service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.Product], error) {
	finded, total, err := s.repo.findAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := response.ProductToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.Product]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}
func (s *service) Update(id string, input Update) error {
	product, err := s.repo.findByID(id)
	if err != nil {
		return err
	}

	if input.Type != nil && !input.Type.IsValid() {
		return errors.New("invalid product type")
	}

	copier.CopyWithOption(&product, &input, copier.Option{IgnoreEmpty: true})

	if input.UnitPrice != nil {
		product.UnitPrice, err = decimal.NewFromString(*input.UnitPrice)
		if err != nil {
			return errors.New("invalid unit price")
		}
	}

	return s.repo.update(product)
}

func (s *service) SoftDelete(id string) error {
	return s.repo.softDelete(id)
}

func (s *service) ValidateInstaller(id string, iduser string) error {
	return s.repo.validateInstaller(id, iduser)
}
