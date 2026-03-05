package project

import (
	"errors"

	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/MetaDandy/carpyen-service/src/response"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type Service interface {
	Create(input Create, userID string) error
	FindByID(id string) (*response.Project, error)
	FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.Project], error)
	Update(id string, input Update) error
	SoftDelete(id string) error
	ValidateSeller(id string, iduser string) error
}
type UserRepo interface {
	FindByID(id string) (model.User, error)
}
type ClientRepo interface {
	FindByID(id string) (model.Client, error)
}

type service struct {
	repo       Repo
	userRepo   UserRepo
	clientRepo ClientRepo
}

func NewService(repo Repo, userRepo UserRepo, clientRepo ClientRepo) Service {
	return &service{repo: repo, userRepo: userRepo, clientRepo: clientRepo}
}

func (s *service) Create(input Create, userID string) error {

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	client, err := s.clientRepo.FindByID(input.ClientID)
	if err != nil {
		return err
	}

	project := model.Project{}
	copier.Copy(&project, &input)

	project.ID = uuid.New()
	project.UserID = user.ID
	project.ClientID = client.ID
	project.State = enum.StatusPending

	return s.repo.Create(project)
}

func (s *service) FindByID(id string) (*response.Project, error) {
	project, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	dto := response.ProjectToDto(&project)
	return &dto, nil
}

func (s *service) FindAll(opts *helper.FindAllOptions) (*response.Paginated[response.Project], error) {
	finded, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}

	dtos := response.ProjectToListDto(finded)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	paginated := &response.Paginated[response.Project]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}

	return paginated, nil
}
func (s *service) Update(id string, input Update) error {
	project, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if input.Name != nil {
		project.Name = *input.Name
	}

	if input.Description != nil {
		project.Description = *input.Description
	}

	if input.Location != nil {
		project.Location = *input.Location
	}

	if input.State != nil {
		if !input.State.IsValid() {
			return errors.New("invalid project state")
		}
		project.State = *input.State
	}

	if input.ClientID != nil {
		client, err := s.clientRepo.FindByID(*input.ClientID)
		if err != nil {
			return err
		}
		project.ClientID = client.ID
	}

	return s.repo.Update(project)
}

func (s *service) SoftDelete(id string) error {
	return s.repo.SoftDelete(id)
}

func (s *service) ValidateSeller(id string, iduser string) error {
	return s.repo.ValidateSeller(id, iduser)
}
