package ui

import (
	"embed"
	"errors"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:generate yarn
//go:generate yarn run build
//go:embed all:build
var files embed.FS

func SvelteKitHandler(path string) http.Handler {
	fsys, fsysErr := fs.Sub(files, "build")

	if fsysErr != nil {
		panic(fsysErr)
	}
	filesystem := http.FS(fsys)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, path)

		_, err := filesystem.Open(path)
		if errors.Is(err, os.ErrNotExist) {
			path += ".html"
		}

		r.URL.Path = path

		http.FileServer(filesystem).ServeHTTP(w, r)
	})
}

func Start(router *gin.Engine) {
	router.Use(gin.WrapH(SvelteKitHandler("/")))
}

func Stop() {}
