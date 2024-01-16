package router

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)


func Router () *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/", handleHome).Methods("GET")

	return r;
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	message:= "The cake is a lie"
	fmt.Fprintf(w, "Format strings here: %s", message)
}