package main

import (
	"log"
	"net/http"
)

type Config struct {
	
}

func main() {

	app := Config{
	
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