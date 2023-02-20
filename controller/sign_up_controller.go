package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"net/http"
	"github.com/auth-server-go/model"
	"github.com/auth-server-go/database"
	"github.com/auth-server-go/utils"
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

	// save the hash of password not the actual password


	
    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err1 := json.NewDecoder(r.Body).Decode(&u)
	u.Password = strconv.FormatUint(uint64(utils.Hash(u.Password)),10)

    if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad request 400")
        return
    }
    
	result := db.Create(&u)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User name already exists")
		return 
	}
	fmt.Fprintf(w, "User: %+v", u)
	defer db.Close()
	return
}
