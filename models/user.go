package models

import (
	"bitbucket.org/indoquran-api/helpers"
	"bitbucket.org/indoquran-api/models/modelstruct"

	"gopkg.in/mgo.v2/bson"
)

// Signup handles registering a user
func (u *UserModel) Signup(data modelstruct.SignupUserCommand) error {
	collection := DBConnect.MongoUse(DatabaseName, CollUser)
	user := &UserModel{
		Name:     data.Name,
		Email:    data.Email,
		Password: helpers.GeneratePasswordHash([]byte(data.Password)),
	}

	return collection.Insert(user)
}

// GetUserByEmail : search user by email
func (u *UserModel) GetUserByEmail(email string) (user modelstruct.User, err error) {
	collection := DBConnect.MongoUse(DatabaseName, CollUser)

	err = collection.Find(bson.M{"email": email}).One(&user)

	return user, err
}
