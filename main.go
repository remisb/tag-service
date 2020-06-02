package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/remisb/tag-service/tag"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const port = 8000

func main() {
	startServer()
}

func startServer() {
	r := chi.NewRouter()

	r.Route("/api/v1/", func(r chi.Router) {
		r.Mount("/tags", tag.InitRouter())
	})

	address := fmt.Sprintf(":%d", port)
	log.Infof("Server started at %s", address)
	http.ListenAndServe(address, r)
}
