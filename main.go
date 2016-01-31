package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/bentrevor/calhoun/app"
	"github.com/bentrevor/calhoun/db"
	"github.com/bentrevor/calhoun/web"

	"github.com/namsral/flag"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "some request info:\n%q\n%q\n%q\n%q",
		html.EscapeString(r.URL.Path),
		html.EscapeString(r.Host),
		html.EscapeString(r.RequestURI),
		html.EscapeString(r.RemoteAddr),
	)
}

func main() {
	var (
		rootDir   string
		assetPath string
		srvPath   string
		ui        string
	)

	flag.StringVar(&rootDir, "root-dir", "/home/vagrant/go/src/github.com/bentrevor/calhoun", "project root")
	flag.StringVar(&assetPath, "asset-path", fmt.Sprintf("web/assets", rootDir), "asset path")
	flag.StringVar(&srvPath, "srv-path", fmt.Sprintf("%s/images/srv", assetPath), "path to save uploaded files")

	// for now, just http/json server, but eventually cli inputs, mobile apps, etc.
	flag.StringVar(&ui, "ui", "web", "")

	flag.Parse()

	switch ui {
	case "web":
		postgresDB := db.NewPostgresDB("dev")
		realFS := db.NewRealFS(srvPath)
		store := app.CalhounStore{DB: postgresDB, FS: realFS, SrvPath: srvPath}
		renderer := web.BrowserRenderer{ViewsPath: fmt.Sprintf("%s/web/views", rootDir)}

		server := app.CalhounServer{
			Store:         store,
			Renderer:      renderer,
			AssetPath:     assetPath,
			FullAssetPath: fmt.Sprintf("%s/%s", rootDir, assetPath),
		}

		server.Run("dev")
	default:
		log.Fatal("can only use web ui for now: `", ui, "` not supported")
	}
}
