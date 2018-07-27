package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    "json:id"
	Name  string "json:name"
	Email string "json:email"
	First string "json:first"
	Last  string "json:last"
}

func createUser(w http.ResponseWriter, r *http.Request) {
	NewUser := User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")
	output, err := json.Marshal(NewUser)
	if err != nil {
		fmt.Println("Something Went Wrong")
	}
	fmt.Println(string(output))
	database, err := sql.Open("mysql", "localhost:3306/")
	if err != nil {
		fmt.Println(err)
	}
	query := "INSER INTO users set user_nickname='" + NewUser.Name + "',user_first='" + NewUser.First + "',user_last='" + NewUser.Last + "'user_email='" + NewUser.Email + "'"
	q, err := database.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
}

func main() {
	routes := mux.NewRouter()
	routes.HandleFunc("/api/user/create", createUser).Methods("GET")
	http.ListenAndServe(":8080", nil)
}
