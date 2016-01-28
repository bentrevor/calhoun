package db

type MemoryFS struct {
	Photos map[string]Photo
	Asdf   []int
}

func NewMemoryFS() MemoryFS {
	memoryFS := MemoryFS{}
	memoryFS.Photos = make(map[string]Photo)

	return memoryFS
}

func (fs *MemoryFS) WritePhoto(photo Photo) {
	fs.Photos[string(photo.Id)] = photo
}

func (fs *MemoryFS) CountPhotos() int {
	return len(fs.Photos)
}
