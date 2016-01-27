package db

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
