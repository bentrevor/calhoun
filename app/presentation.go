package app

import "io"

type CalhounServer interface {
	RegisterRoutes()
	Start()
}

type CalhounRenderer interface {
	UploadPhotoForm(io.Writer)
	UploadPhoto(io.Writer)
	ViewPhotos(io.Writer, []Photo)
}
