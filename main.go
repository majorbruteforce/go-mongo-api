package main

import (
	"github.com/majorbruteforce/go-mongo-api/router"
	"github.com/majorbruteforce/go-mongo-api/controller"
	"log"
	"net/http"
	"time"
	"fmt"
)

func main() {
	r := router.Router()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	

}
