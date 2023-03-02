package controller

import (
	"encoding/json"
	"net/http"
	"github.com/auth-server-go/requests"
	"github.com/auth-server-go/model"
	"github.com/auth-server-go/database"
)

func Handle_login_request(w http.ResponseWriter, r *http.Request) {
	// get username and public key from user
		// extract username and the public key from the request body
	var login requests.Login

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	decoder := json.NewDecoder(r.Body)

    err := decoder.Decode(&login)
    if err != nil {
        panic(err)
    }
	
	// Get the salt of the user from the username
	var user model.User
	
	user = database.Get_user_security_information(login.Username)
	w.Header().Set("Content-Type", "application/json") 
	if user.Password!=login.Password {
		panic("asdasd")
	}
}


