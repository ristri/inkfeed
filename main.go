package main

import (
	"inkfeed/templates"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := os.MkdirAll("templates", 0755); err != nil {
		log.Fatal(err)
	}

	htmlTemplates, err := templates.LoadTemplates()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/r/", htmlTemplates.HandleSubreddit())
	http.HandleFunc("/comments/", htmlTemplates.HandleComments())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/r/popular", http.StatusTemporaryRedirect)
		} else {
			http.NotFound(w, r)
		}
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
