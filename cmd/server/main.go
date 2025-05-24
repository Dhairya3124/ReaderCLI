package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ReaderCLI Server Running"))
	})
	srv := http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	log.Fatal(srv.ListenAndServe())
}
