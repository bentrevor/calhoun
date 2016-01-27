package db

type QueryOpts struct {
	User  User
	Photo Photo
}

type PhotoDB interface {
	Insert(QueryOpts) (bool, error)
	Select(QueryOpts) []Photo
}

type User struct {
	Id   int64
	Name string
}

type Photo struct {
	Filepath string
}
