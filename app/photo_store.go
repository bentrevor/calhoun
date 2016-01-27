package app

import (
	. "github.com/bentrevor/calhoun/db"
)

type PhotoStore struct {
	DB PhotoDB
}

func (store PhotoStore) SavePhoto(user User, photo Photo) (bool, error) {
	options := QueryOpts{User: user, Photo: photo}

	store.DB.Insert(options)

	return true, nil
}

func (store PhotoStore) PhotosForUser(user User) []Photo {
	options := QueryOpts{User: user}

	return store.DB.Select(options)
}
