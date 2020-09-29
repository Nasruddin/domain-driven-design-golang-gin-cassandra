package model

import "golang-gin-cassandra/src/utils/errors"

/*
var (
	users = map[string]*User{
		"123": {
			ID:        "123",
			FirstName: "Nasir",
			LastName:  "Khan",
			FullName:  "Mohammed Nasir Khan",
			Age:       30,
			EmailId:   "nasir_khan@gmail.com",
		},
	}
)*/

type User struct {
	ID string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	FullName string `json:"full_name"`
	Age int `json:"age"`
	EmailId string `json:"email_id"`
}

func (user *User) ValidateUser() *errors.RestErr {
	if user.Age < 0 {
		return errors.NewInternalServerError("age can't be less than 0", nil)
	}

	if user.EmailId == "" {
		return errors.NewInternalServerError("email id can't be blank", nil)
	}
	// TODO:: Add more validation
	return nil
}

/*func (u *User) GetByID(userID string) (*User, *restError.RestErr) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, restError.NewInternalServerError(
		"Unable to get user. Database error", errors.New("user not found"),
		)
}*/