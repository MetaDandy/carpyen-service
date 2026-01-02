package client

import (
	"fmt"

	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/MetaDandy/carpyen-service/src/response"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type Service interface {
	Login(input Login) (string, error)
	Create(input Create, userID string) error
	Update(id string, input Update) error
	UpdateProfile(id string, input UpdateProfile) error
	GetByID(id string) (*response.Client, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.Client], error)
	SoftDelete(id string) error

	ValidateSeller(id string, idclient string) error
}

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{repo: repo}
}

func (s *service) Login(input Login) (string, error) {
	client, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	if !helper.CheckPasswordHash(input.Password, client.Password) {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := helper.GenerateJwt(client.ID.String(), client.Email, enum.RoleClient.String())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) Create(input Create, userID string) error {

	if input.Password != input.ConfirmPassword {
		return fmt.Errorf("password and confirm password do not match")
	}

	hash, err := helper.HashPassword(input.Password)
	if err != nil {
		return err
	}

	client := model.Client{}
	copier.Copy(&client, &input)
	client.ID = uuid.New()
	client.Password = hash
	client.UserID = uuid.MustParse(userID)

	if err := s.repo.Create(client); err != nil {
		return err
	}

	return nil
}

func (s *service) Update(id string, input Update) error {
	client, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	copier.CopyWithOption(&client, &input, copier.Option{IgnoreEmpty: true})

	if input.Password != nil {
		if input.ConfirmPassword == nil || *input.Password != *input.ConfirmPassword {
			return fmt.Errorf("password and confirm password do not match")
		}

		hash, err := helper.HashPassword(*input.Password)
		if err != nil {
			return err
		}
		client.Password = hash
	}

	if err := s.repo.Update(client); err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateProfile(id string, input UpdateProfile) error {
	client, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	copier.CopyWithOption(&client, &input, copier.Option{IgnoreEmpty: true})

	if err := s.repo.Update(client); err != nil {
		return err
	}

	return nil
}

func (s *service) GetByID(id string) (*response.Client, error) {
	client, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	dto := response.ClientToDto(&client)

	return &dto, nil
}

func (s *service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.Client], error) {
	finded, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}
	dtos := response.ClientToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &response.Paginated[response.Client]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *service) SoftDelete(id string) error {
	return s.repo.SoftDelete(id)
}

func (s *service) ValidateSeller(id string, idclient string) error {
	return s.repo.ValidateSeller(id, idclient)
}
