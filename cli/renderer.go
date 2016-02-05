package cli

import (
	"fmt"
	"io"

	. "github.com/bentrevor/calhoun/app"
)

type ConsoleRenderer struct{}

func (ConsoleRenderer) UploadPhoto(w io.Writer) {
	fmt.Fprintf(w, "you just uploaded a photo")
}

// ISP!!
func (ConsoleRenderer) UploadPhotoForm(w io.Writer) {
	fmt.Fprintf(w, "noop (upload form)")
}

func (ConsoleRenderer) ViewPhotos(w io.Writer, _ []Photo) {
	fmt.Fprintf(w, "noop (view photos)")
}
