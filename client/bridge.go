package client

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed all:dist
var BuildFS embed.FS

func BuildHTTPFS() http.FileSystem {
	build, err := fs.Sub(BuildFS, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(build)
}
