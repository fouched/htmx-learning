package main

import (
	"fmt"
	"github.com/fouched/htmx-learning/internal/config"
	"github.com/fouched/htmx-learning/internal/handlers"
	"github.com/fouched/htmx-learning/internal/render"
	"log"
	"net/http"
)

const port = ":9080"

var appConfig config.AppConfig

func main() {

	run()

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}
	fmt.Println(fmt.Sprintf("Starting application on %s", port))

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func run() {
	hc := handlers.NewConfig(&appConfig)
	handlers.NewHandlers(hc)
	render.NewRenderer(&appConfig)
}
