package serviceuser

import (
	"github.com/arfan21/golang-mygram/exception"
	"github.com/arfan21/golang-mygram/repository/repositoryuser"
	"github.com/arfan21/golang-mygram/validation"
)

type ServiceUser interface {
	Create(data *model.UserRequest) (model.UserResponse, error)
}

type service struct {
	repo repositoryuser.RepositoryUser
}

func New(repo repositoryuser.RepositoryUser) ServiceUser {
	return &service{repo: repo}
}

func (s *service) Create(data *model.UserRequest) (model.UserResponse, error) {
	err := validation.ValidateUserCreate(data)
	if err != nil {
		return model.UserResponse{}, exception.ErrorValidation{Message: err}
	}

	createdData, err := s.repo.Create(entity.User{})
	if err != nil {
		return model.UserResponse{}, err
	}

	// createdData diubah ke response

	return model.UserResponse{}, nil
}
