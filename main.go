package main

import (
	"fmt"
	"log"

	"github.com/bentrevor/calhoun/app"
	"github.com/bentrevor/calhoun/db"

	// presentation layers
	"github.com/bentrevor/calhoun/cli"
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
	flag.StringVar(&ui, "ui", "web", "")

	flag.Parse()

	var server app.CalhounServer
	postgresDB := db.NewPostgresDB("dev")
	realFS := db.RealFS{RootDir: srvPath}
	store := app.CalhounStore{DB: postgresDB, FS: realFS, SrvPath: srvPath}
	calhoun := app.Calhoun{Store: store}

	switch ui {
	case "web":
		calhoun.Renderer = web.BrowserRenderer{
			ViewsPath:  fmt.Sprintf("%s/web/views", rootDir),
			PhotosPath: srvPath,
		}

		server = &web.WebServer{
			App:           calhoun,
			AssetPath:     assetPath,
			FullAssetPath: fmt.Sprintf("%s/%s", rootDir, assetPath),
		}
	case "cli":
		calhoun.Renderer = cli.ConsoleRenderer{}

		server = cli.ConsoleServer{
			App:  calhoun,
			Args: flag.Args(),
		}
	default:
		log.Fatal(ui, " not supported")
	}

	app.Run("dev", server)
}
