package models

import (
	"indoquran-golang/helpers"

	"gopkg.in/mgo.v2/bson"
)

// Signup handles registering a user
func (u *UserModel) Signup(data SignupUserCommand) error {
	collection := DBConnect.MGOUse(DatabaseName, CollUser)
	user := &UserModel{
		Name:     data.Name,
		Email:    data.Email,
		Password: helpers.GeneratePasswordHash([]byte(data.Password)),
	}

	return collection.Insert(user)
}

// GetUserByEmail : search user by email
func (u *UserModel) GetUserByEmail(email string) (user User, err error) {
	collection := DBConnect.MGOUse(DatabaseName, CollUser)

	err = collection.Find(bson.M{"email": email}).One(&user)

	return user, err
}
