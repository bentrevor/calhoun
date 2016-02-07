package web

import (
	"fmt"
	"html/template"
	"io"
	"log"

	. "github.com/bentrevor/calhoun/app"
)

type Page struct {
	PhotoTags []PhotoTag
}

type PhotoTag struct {
	Src       string
	ClassName string
}

type BrowserRenderer struct {
	ViewsPath  string
	PhotosPath string
}

func (br *BrowserRenderer) renderHtmlFile(filename string, writer io.Writer, page Page) {
	layoutPath := fmt.Sprintf("%s/layout.html", br.ViewsPath)
	filepath := fmt.Sprintf("%s/%s.html", br.ViewsPath, filename)

	tmpl, err := template.ParseFiles(layoutPath, filepath)

	if err != nil {
		log.Fatal("error on template.ParseFiles: ", err)
	}

	tmpl.Execute(writer, page)
}

func (br BrowserRenderer) Render(action CalhounAction, w io.Writer, args ...RenderArgs) {
	switch action {
	case UploadPhotoForm:
		br.renderHtmlFile("upload_photo_form", w, Page{})
	case UploadPhoto:
		br.renderHtmlFile("upload_success", w, Page{})
	case ViewPhotos:
		photoTags := []PhotoTag{}

		for _, photo := range args[0].Photos {
			photoTags = append(photoTags, PhotoTag{
				Src:       fmt.Sprintf("%s/%s", br.PhotosPath, photo.Src),
				ClassName: fmt.Sprintf("image_%d", photo.Id),
			})
		}

		br.renderHtmlFile("view_photos", w, Page{PhotoTags: photoTags})
	default:
		log.Fatal(action, " not implemented for web")
	}
}
