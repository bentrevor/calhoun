package calhoun

type PhotoStore struct {
	DB *PhotoDB
}

type User struct {
	Name string
}

type Photo struct {
	Filepath string
}

func (store PhotoStore) SavePhoto(user User, photo Photo) (bool, error) {
	options := map[string]string{"user": user.Name}

	(*store.DB).Insert(photo, options)

	return true, nil
}

func (store PhotoStore) PhotosForUser(user User) []Photo {
	options := map[string]string{"user": user.Name}

	return (*store.DB).Select(options)
}
