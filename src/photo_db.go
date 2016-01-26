package calhoun

type QueryOpts map[string]string

type PhotoDB interface {
	Insert(Photo, QueryOpts) (bool, error)
	Select(QueryOpts) []Photo
}

type MemoryDB struct {
	Photos map[string][]Photo
}

func NewMemoryDB() MemoryDB {
	memoryDB := MemoryDB{}
	memoryDB.Photos = make(map[string][]Photo)

	return memoryDB
}

func (mk *MemoryDB) Insert(photo Photo, opts QueryOpts) (bool, error) {
	user := opts["user"]
	mk.Photos[user] = append(mk.Photos[user], photo)

	return true, nil
}

func (mk *MemoryDB) Select(opts QueryOpts) []Photo {
	return mk.Photos[opts["user"]]
}

// func (pg *PostgresDB) Insert(photo Photo, opts QueryOpts) (bool, error) {
// 	pg.Query("INSERT INTO photos (%s) values (%s)", columns, values)

// 	return true, nil
// }

// func (pg *PostgresDB) Select(opts QueryOpts) []Photo {
// 	scope = opts["scope"] || "*"
// 	userId = opts["userId"]
// 	tags = opts["tags"]
// 	query = buildFrom(userId, tags, etc...)

// 	return pg.Query("SELECT %s FROM photos WHERE %s", scope, query)
// }
