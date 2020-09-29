package repository

import (
	"golang-gin-cassandra/src/domain/users/model"
	"golang-gin-cassandra/src/utils/errors"
)

type UserRepository interface {
	GetByID(userID string) (*model.User, *errors.RestErr)
	Create(user model.User) (*model.User, *errors.RestErr)
}
