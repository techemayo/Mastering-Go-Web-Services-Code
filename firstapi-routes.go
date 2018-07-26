package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/drone/routes"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()
	name := urlParams.Get(":name")
	helloMessage := "Hello, " + name
	message := API{helloMessage}
	output, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Something went wrong")
	}
	fmt.Fprintf(w, string(output))
}

type API struct {
	Message string "json:message"
}

func main() {
	mux := routes.New()
	mux.Get("/api/:name", Hello)
	http.Handle("/", mux)
	http.ListenAndServe(":8080", nil)
}
