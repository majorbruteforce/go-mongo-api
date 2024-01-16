package router

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)


func Router () *mux.Router {

	r := mux.NewRouter()
	s:= r.Methods("GET").Subrouter() // Prevents writing same matchers repeatedly
	s.HandleFunc("/", handleHome)
	s.HandleFunc("/vars/{variable}", handleRange).Methods("GET","POST")
	s.HandleFunc("/regex/{id:[0-9]+}", handleRange)
	

	return r;
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	message:= "The cake is a lie"
	fmt.Fprintf(w, "Format : %s", message)
}

func handleRange(w http.ResponseWriter, r *http.Request){
	vars:= mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"",vars)
}