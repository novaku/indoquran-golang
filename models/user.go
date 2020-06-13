package models

import (
	"indoquran-golang/forms"
	"indoquran-golang/helpers"

	"gopkg.in/mgo.v2/bson"
)

// User defines user object structure
type User struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Email      string        `json:"email" bson:"email"`
	Password   string        `json:"password" bson:"password"`
	IsVerified bool          `json:"is_verified" bson:"is_verified"`
}

// UserModel defines the model structure
type UserModel struct{}

// Signup handles registering a user
func (u *UserModel) Signup(data forms.SignupUserCommand) error {
	// Connect to the user collection
	collection := dbConnect.MGOUse(databaseName, "user")
	// Assign result to error object while saving user
	err := collection.Insert(bson.M{
		"name":     data.Name,
		"email":    data.Email,
		"password": helpers.GeneratePasswordHash([]byte(data.Password)),
		// This will come later when adding verification
		"is_verified": false,
	})

	return err
}

// GetUserByEmail : search user by email
func (u *UserModel) GetUserByEmail(email string) (user User, err error) {
	collection := dbConnect.MGOUse(databaseName, "user")

	err = collection.Find(bson.M{"email": email}).One(&user)

	return user, err
}
