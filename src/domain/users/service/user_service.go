package service

import (
	"errors"
	"fmt"
	"golang-gin-cassandra/src/domain/users/model"
	"golang-gin-cassandra/src/domain/users/repository"
	. "golang-gin-cassandra/src/utils/errors"
	"strings"
)

type UserService interface {
	GetByID(userID string) (*model.User, *RestErr)
	Create(user model.User) (*model.User, *RestErr)
}

type userService struct {
	repository repository.UserRepository
}

func (s *userService) Create(user model.User) (*model.User, *RestErr) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}
	return s.repository.Create(user)
}

func NewService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) GetByID(userID string) (*model.User, *RestErr) {
	userID = strings.TrimSpace(userID)
	if len(userID) == 0 {
		return nil, NewBadRequestError("invalid user id. UserId can't be empty")
	}
	user, err := s.repository.GetByID(userID)
	if err != nil {
		userNotFoundErr := fmt.Sprintf("user not found for user id %s", userID)
		return nil, NewInternalServerError(userNotFoundErr, errors.New("here"))
	}
	return user, nil
}


