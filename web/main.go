package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

func main() {
	useOS := len(os.Args) > 1 && os.Args[1] == "live"
	http.Handle("/", http.FileServer(getFileSystem(useOS)))
	http.ListenAndServe(":8888", nil)
}

func getFileSystem(useOS bool) http.FileSystem {
	var fsys fs.FS
	if useOS {
		log.Print("using live mode")
		fsys = os.DirFS("static")
	} else {
		log.Print("using embed mode")
		var (
			//go:embed static
			files embed.FS
			err   error
		)
		fsys, err = fs.Sub(files, "static")
		if err != nil {
			panic(err)
		}
	}
	return http.FS(fsys)
}
