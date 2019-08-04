package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/johan-lejdung/go-microservice-middleware-guide/mw"
	"github.com/urfave/negroni"
)

func main() {
	router := mux.NewRouter()
	router.
		Methods("GET").
		Path("/").
		HandlerFunc(endpointAction)

	n := negroni.New()
	n.Use(&mw.Logger{})
	n.UseHandler(router)

	err := http.ListenAndServe(":8080", n)
	if err != nil {
		panic(err)
	}
}

func endpointAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint finished")
}
