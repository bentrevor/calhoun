package db

type MemoryDB struct {
	Photos map[string][]Photo
}

func NewMemoryDB() MemoryDB {
	memoryDB := MemoryDB{}
	memoryDB.Photos = make(map[string][]Photo)

	return memoryDB
}

func (mk *MemoryDB) Insert(opts QueryOpts) int {
	user := opts.User
	mk.Photos[user.Name] = append(mk.Photos[user.Name], opts.Photo)

	return len(mk.Photos)
}

func (mk *MemoryDB) Select(opts QueryOpts) []Photo {
	return mk.Photos[opts.User.Name]
}
