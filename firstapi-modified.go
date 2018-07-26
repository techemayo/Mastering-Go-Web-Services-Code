package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {
	urlParam := mux.Vars(r)
	name := urlParam["user"]
	helloMessage := "Hello, " + name
	message := API{helloMessage}
	output, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	fmt.Fprintf(w, string(output))
}

func main() {
	gorrilaRoute := mux.NewRouter()
	gorrilaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)
	http.Handle("/", gorrilaRoute)
	http.ListenAndServe(":8080", nil)
}
