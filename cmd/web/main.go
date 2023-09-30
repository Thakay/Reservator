package main

import (
	"fmt"
	"github.com/thakay/Reservator/pkg/config"
	"github.com/thakay/Reservator/pkg/handlers"
	"github.com/thakay/Reservator/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template Cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("starting appliaction on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
