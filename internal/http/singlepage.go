package http

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ayushg3112/dirlist/internal/templates"
	"github.com/ayushg3112/dirlist/sort"
	"github.com/ayushg3112/dirlist/walk"
)

type ServerOptions struct {
	Port           string
	RootDirAbsPath string
}

func StartSinglePageServer(walker walk.Walker, sorter sort.DirEntrySorter, options ServerOptions) error {
	log.Printf("starting the server at port %s", options.Port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "" || path == "/" {
			structure, err := walker.Walk(sorter)
			if err != nil {
				fmt.Fprintf(w, "error while traversing structure: %v", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			html, err := templates.GenerateSinglePageTemplateHTML(structure)

			if err != nil {
				fmt.Fprintf(w, "error while generating HTML: %v", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			fmt.Fprint(w, html)
			return
		}

		filePath := filepath.Join(options.RootDirAbsPath, path)

		http.ServeFile(w, r, filePath)
	})

	return http.ListenAndServe(":"+options.Port, nil)
}
