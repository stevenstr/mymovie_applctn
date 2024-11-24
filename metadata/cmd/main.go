package main

import (
	"log"
	"net/http"

	"cavea-movie.ge/metadata/internal/controller/metadata"
	httphandler "cavea-movie.ge/metadata/internal/handler/http"
	"cavea-movie.ge/metadata/internal/repository/memory"
)

func main() {
	log.Println("Start the movie metadata service...")

	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)

	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
