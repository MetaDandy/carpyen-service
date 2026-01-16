package material

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
	FindByID(id string) (*response.Material, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.Material], error)
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
		return errors.New("invalid material type")
	}
	if !input.UnitMeasure.IsValid() {
		return errors.New("invalid unit measure")
	}

	material := model.Material{}
	copier.Copy(&material, &input)
	material.ID = uuid.New()
	material.UserID = user.ID

	material.UnitPrice, err = decimal.NewFromString(input.UnitPrice)
	if err != nil {
		return errors.New("invalid unit price")
	}

	return s.repo.Create(material)
}

func (s *service) FindByID(id string) (*response.Material, error) {
	material, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.MaterialToDto(&material)
	return &dto, nil
}

func (s *service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.Material], error) {
	finded, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := response.MaterialToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.Material]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}
func (s *service) Update(id string, input Update) error {
	material, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if input.Type != nil && !input.Type.IsValid() {
		return errors.New("invalid material type")
	}
	if input.UnitMeasure != nil && !input.UnitMeasure.IsValid() {
		return errors.New("invalid unit measure")
	}

	copier.CopyWithOption(&material, &input, copier.Option{IgnoreEmpty: true})

	if input.UnitPrice != nil {
		material.UnitPrice, err = decimal.NewFromString(*input.UnitPrice)
		if err != nil {
			return errors.New("invalid unit price")
		}
	}

	return s.repo.Update(material)
}

func (s *service) SoftDelete(id string) error {
	return s.repo.SoftDelete(id)
}

func (s *service) ValidateInstaller(id string, iduser string) error {
	return s.repo.ValidateInstaller(id, iduser)
}
