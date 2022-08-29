package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/auth-server-go/model"
	"github.com/auth-server-go/database"
)


func Handle_sign_up_request(w http.ResponseWriter, r *http.Request) {
	db := database.ConnectDB()
	db.LogMode(true)
	// Implement sign up function with db
	err := r.ParseForm()
	if err != nil {
		// in case of any error
		fmt.Printf("Error parsing the form")
		return
	}
	var u model.User

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err1 := json.NewDecoder(r.Body).Decode(&u)
    if err1 != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Do something with the Person struct...
    fmt.Fprintf(w, "User: %+v", u)
	fmt.Printf("UserL")
	result := db.Create(&u)
	if result.Error != nil {
		// log.Fatal(result.Error)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	defer db.Close()
	return
}
