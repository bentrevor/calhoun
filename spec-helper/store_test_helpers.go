package spec_helper

import (
	"fmt"

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
	id := len(mk.Photos) + 1
	photo := opts.Photo
	photo.Id = id
	mk.Photos[opts.User.Name] = append(mk.Photos[opts.User.Name], photo)

	return id
}

func (mk *MemoryDB) Select(opts QueryOpts) []Photo {
	return mk.Photos[opts.User.Name]
}

type MemoryFS struct {
	SrvPath string
	Photos  map[string]Photo
}

func NewMemoryFS(srvPath string) *MemoryFS {
	memoryFS := MemoryFS{Photos: make(map[string]Photo), SrvPath: srvPath}
	return &memoryFS
}

func (fs *MemoryFS) PhotoSrc(id int) string {
	return fmt.Sprintf("%s/%d", fs.SrvPath, id)
}

func (fs *MemoryFS) WritePhoto(photo Photo) {
	fs.Photos[string(photo.Id)] = photo
}

func (fs *MemoryFS) CountPhotos() int {
	return len(fs.Photos)
}
