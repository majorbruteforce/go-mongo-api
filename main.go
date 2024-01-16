package main

import (
	"net/http"
	"time"
	"log"
	"github.com/majorbruteforce/go-mongo-api/router"
)

func main() {
	r:= router.Router();
	serverConfig := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(serverConfig.ListenAndServe())

}


