package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/noogler-eng/go-lang/tree/main/mysql/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.BookRouter(r);
	// func http.Handle(pattern string, handler http.Handler)
	http.Handle("/", r);
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("error", err);
	}
}