package datastore

// modify the db so that we store the FavColor as a hex code

import (
	"context"

	"github.com/google/uuid"
)

// Store users in this map to mimic a db
var userStore = make(map[string]User)

type User struct {
	Id       string
	Name     string
	Email    string
	FavColor string
}

func CreateUser(ctx context.Context, name string, email string, favColor string) (string, error) {
	id := uuid.New().String()
	// save the user
	userStore[id] = User{
		Id:       id,
		Name:     name,
		Email:    email,
		FavColor: favColor,
	}

	return id, nil
}
