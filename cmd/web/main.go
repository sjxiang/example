package main

import (
	"log"
	"net/http"
)

type App struct {
	SecretKey string
}

func main() {

	app := App{

	}

	srv := &http.Server{
		Addr:    ":9000",
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}