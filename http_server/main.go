package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"
)

func sendIndex(assetPath string, serve http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexPage := path.Join(assetPath, "index.html")
		requestedPage := path.Join(assetPath, r.URL.Path)

		_, err := os.Stat(requestedPage)

		if err != nil {
			http.ServeFile(w, r, indexPage)
			return
		}

		serve.ServeHTTP(w, r)
	}
}

func main() {
	var (
		port     = flag.String("port", "8080", "Server port")
		contents = flag.String("contents", ".", "Folder to serve")
	)

	flag.Parse()

	folder := flag.Arg(0)

	if folder != "" {
		*contents = folder
	}

	cwd, _ := os.Getwd()
	contentPath := path.Join(cwd, *contents)

	http.HandleFunc("/", sendIndex(contentPath, http.FileServer(http.Dir(contentPath))))

	log.Println("Listening on port", *port)

	err := http.ListenAndServe(":"+*port, nil)

	if err != nil {
		panic(err)
	}
}
