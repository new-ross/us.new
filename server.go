package main

import (
	_ "fmt"
	"log"
	"net/http"
)


func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			fs.ServeHTTP(w, r)
			return
		}
		http.ServeFile(w, r, "./public/index.html")
	})

	server := &http.Server{
		Addr: ":80",
		Handler: mux,
	}

	log.Printf("listening on http://localhost%v\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}