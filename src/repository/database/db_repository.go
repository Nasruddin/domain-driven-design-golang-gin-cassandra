package database

import (
	"errors"
	"fmt"
	"golang-gin-cassandra/src/clients/cassandra"
	"golang-gin-cassandra/src/domain/users/model"
	. "golang-gin-cassandra/src/utils/errors"
)

const (
	queryGetUserById = "SELECT user_id, first_name, last_name, full_name, age, email FROM users WHERE user_id =?"
	queryCreateUser  = "INSERT INTO users (user_id, first_name, last_name, full_name, age, email) VALUES (?, ?, ?, ?, ?, ?)"
)


func NewDbRepository() DbRepository {
	return &dbRepository{}
}


type DbRepository interface {
	GetByID(userID string) (*model.User, *RestErr)
	Create(user model.User) (*model.User, *RestErr)
}

type dbRepository struct {

}

func (repo *dbRepository) Create(user model.User) (*model.User, *RestErr) {

	if err := cassandra.GetSession().Query(queryCreateUser,
		user.ID, user.FirstName, user.LastName, user.FullName, user.Age, user.EmailId).Exec(); err != nil {
		return nil, NewInternalServerError("unable to insert user in db", errors.New(err.Error()))
	}
	return &user, nil
}

func (repo *dbRepository) GetByID(userID string) (*model.User, *RestErr) {

	var user model.User
	if err := cassandra.GetSession().Query(queryGetUserById, userID).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.FullName, &user.Age, &user.EmailId); err != nil {
		if err.Error() == "not found" {
			fmt.Println("here")
			return nil, NewInternalServerError("no user for given user id", errors.New(err.Error()))
		}
		return nil, NewInternalServerError("unable to find user in db", errors.New(err.Error()))
	}
	return &user, nil
}