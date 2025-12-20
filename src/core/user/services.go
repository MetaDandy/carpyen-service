package user

import (
	"fmt"

	"github.com/MetaDandy/go-fiber-skeleton/helper"
	"github.com/MetaDandy/go-fiber-skeleton/src/enum"
	"github.com/MetaDandy/go-fiber-skeleton/src/model"
	"github.com/MetaDandy/go-fiber-skeleton/src/response"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserService interface {
	Login(input LoginDTO) (string, error)
	Create(input CreateUserDTO) error
	Update(id string, input UpdateUserDTO) error
	UpdateProfile(id string, input UpdateUserProfileDTO) error
	GetByID(id string) (*response.User, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.User], error)
	SoftDelete(id string) error
}

type Service struct {
	repo UserRepo
}

func NewService(repo UserRepo) UserService {
	return &Service{repo: repo}
}

func (s *Service) Login(input LoginDTO) (string, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	if !helper.CheckPasswordHash(input.Password, user.Password) {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := helper.GenerateJwt(user.ID.String(), user.Email, user.Role.String())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Create(input CreateUserDTO) error {
	if enum.IsValidRole(input.Role) == false {
		return fmt.Errorf("invalid role in create user: %s", input.Role)
	}

	if input.Password != input.ConfirmPassword {
		return fmt.Errorf("password and confirm password do not match")
	}

	hash, err := helper.HashPassword(input.Password)
	if err != nil {
		return err
	}

	user := model.User{}
	copier.Copy(&user, &input)
	user.ID = uuid.New()
	user.Password = hash
	user.Role = enum.Role(input.Role)

	if err := s.repo.Create(user); err != nil {
		return err
	}

	return nil
}

func (s *Service) Update(id string, input UpdateUserDTO) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	copier.CopyWithOption(&user, &input, copier.Option{IgnoreEmpty: true})

	if input.Role != nil {
		if enum.IsValidRole(*input.Role) == false {
			return fmt.Errorf("invalid role in update user: %s", *input.Role)
		}
		user.Role = enum.Role(*input.Role)
	}

	if input.Password != nil {
		if input.ConfirmPassword == nil || *input.Password != *input.ConfirmPassword {
			return fmt.Errorf("password and confirm password do not match")
		}

		hash, err := helper.HashPassword(*input.Password)
		if err != nil {
			return err
		}
		user.Password = hash
	}

	if err := s.repo.Update(user); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateProfile(id string, input UpdateUserProfileDTO) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	copier.CopyWithOption(&user, &input, copier.Option{IgnoreEmpty: true})

	if err := s.repo.Update(user); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetByID(id string) (*response.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	dto := response.UserToDto(&user)

	return &dto, nil
}

func (s *Service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.User], error) {
	finded, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}
	dtos := response.UserToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &response.Paginated[response.User]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) SoftDelete(id string) error {
	return s.repo.SoftDelete(id)
}
