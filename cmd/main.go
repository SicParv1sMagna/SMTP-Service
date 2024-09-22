package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SicParv1sMagna/SMTP-service/cmd/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	config := config.MustLoad()

	router := chi.NewRouter()

	err := http.ListenAndServe(fmt.Sprintf("localhost:%s", config.HTTPServer.Port), router)

	if err != nil {
		log.Fatal(err)
	}
}
