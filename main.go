package main

import (
	"fmt"
	"log"

	"github.com/bentrevor/calhoun/app"
	"github.com/bentrevor/calhoun/db"
	"github.com/bentrevor/calhoun/web"

	"github.com/namsral/flag"
)

func main() {
	var (
		rootDir   string
		assetPath string
		srvPath   string
		ui        string
	)

	flag.StringVar(&rootDir, "root-dir", "/home/vagrant/go/src/github.com/bentrevor/calhoun", "project root")
	flag.StringVar(&assetPath, "asset-path", "web/assets", "asset path")
	flag.StringVar(&srvPath, "srv-path", fmt.Sprintf("%s/images/srv", assetPath), "path to save uploaded files")

	// for now, just http/json server, but eventually cli inputs, mobile apps, etc.
	flag.StringVar(&ui, "ui", "web", "")

	flag.Parse()

	fullAssetPath := fmt.Sprintf("%s/%s", rootDir, assetPath)

	switch ui {
	case "web":
		postgresDB := db.NewPostgresDB("dev")
		realFS := db.RealFS{RootDir: srvPath}
		store := app.CalhounStore{DB: postgresDB, FS: realFS, SrvPath: srvPath}

		renderer := web.BrowserRenderer{
			ViewsPath:  fmt.Sprintf("%s/web/views", rootDir),
			PhotosPath: srvPath,
		}

		calhoun := app.Calhoun{
			Store:    store,
			Renderer: renderer,
		}

		server := web.WebServer{
			App:           calhoun,
			AssetPath:     assetPath,
			FullAssetPath: fullAssetPath,
		}

		app.Run("dev", server)
	case "cli":
		log.Fatal("cli not implemented yet")
	default:
		log.Fatal("can only use web ui for now: `", ui, "` not supported")
	}
}
