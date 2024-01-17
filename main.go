package main

import (
	"github.com/majorbruteforce/go-mongo-api/router"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Codot struct{
	ID primitive.ObjectID `json:"_id, omitempty" bson:"_id,omitempty"`
	Type string `json:"type,omitempty"`
	Defeated bool `json:"defeated,omitempty"`
}


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
