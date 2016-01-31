package spec_helper

import (
	. "github.com/bentrevor/calhoun/app"
)

type MemoryDB struct {
	Photos map[string][]Photo
}

func NewMemoryDB() *MemoryDB {
	memoryDB := MemoryDB{Photos: make(map[string][]Photo)}
	return &memoryDB
}

func (mk *MemoryDB) Insert(opts QueryOpts) int {
	mk.Photos[opts.User.Name] = append(mk.Photos[opts.User.Name], opts.Photo)
	return len(mk.Photos)
}

func (mk *MemoryDB) Select(opts QueryOpts) []Photo {
	return mk.Photos[opts.User.Name]
}

type MemoryFS struct {
	Photos map[string]Photo
}

func NewMemoryFS() *MemoryFS {
	memoryFS := MemoryFS{Photos: make(map[string]Photo)}
	return &memoryFS
}

func (*MemoryFS) RootDir() string {
	return "/fake/srv/path"
}

func (fs *MemoryFS) WritePhoto(photo Photo) {
	fs.Photos[string(photo.Id)] = photo
}

func (fs *MemoryFS) CountPhotos() int {
	return len(fs.Photos)
}
