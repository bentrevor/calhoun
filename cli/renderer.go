package cli

import (
	"fmt"
	"io"
	"log"

	. "github.com/bentrevor/calhoun/app"
)

type ConsoleRenderer struct{}

func (ConsoleRenderer) Render(action CalhounAction, w io.Writer, args ...RenderArgs) {
	switch action {
	case UploadPhoto:
		fmt.Fprintf(w, "you just uploaded a photo")
	default:
		log.Fatal(action, " not implemented for cli")
	}
}
