package db

type QueryOpts map[string]string

type PhotoDB interface {
	Insert(Photo, QueryOpts) (bool, error)
	Select(QueryOpts) []Photo
}

type User struct {
	Name string
}

type Photo struct {
	Filepath string
}
