package user

import (
	"context"
	"log"
	"pockethealth/internchallenge/pkg/datastore"
)

// ApiServices are an instance of an ApiServicer that implement the api actions for an ApiServicer
type UserApiService struct{}

// NewUserApiService creates a new Api Service
func NewUserApiService() UserApiService {
	return UserApiService{}
}

// PostRegister - Register a User
func (s UserApiService) PostRegister(ctx context.Context, name string, email string, favColor string) (string, error) {
	// save user to datastore
	userId, err := datastore.CreateUser(ctx, name, email, favColor)
	if err != nil {
		log.Printf("error creating user: %q", err.Error())
		return "", err
	}
	log.Printf("created user with id: %s\n", userId)

	// return the user id
	// task 3: here is the bug, the user id is not being returned
	return userId, err
}
