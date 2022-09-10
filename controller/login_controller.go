package controller

import (
	"encoding/json"
	"math"
	"net/http"
	"github.com/auth-server-go/requests"
	"github.com/auth-server-go/response"
	"github.com/auth-server-go/model"
	"github.com/auth-server-go/database"
	"github.com/auth-server-go/constant"
	"strconv"
	"fmt"
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

	// send salt back and publich key with a verifier and public key information
		// access the salt according to username from the db
		// calculate the information to be sent
		// send back the information after calculating the session key for future authentication
		var ep_key int64
		verifier_int, err := strconv.Atoi(user.Verifier)
		ep_key = int64(verifier_int*constant.K) + int64(math.Pow(constant.G, constant.B))
		ep_key_str := strconv.FormatInt(int64(ep_key), 10)
		login_response := response.Login {
			Salt: user.Salt,
			Ephemeral_key: ep_key_str,
		} 
		json.NewEncoder(w).Encode(login_response)
	// Calculate session key for future authentication
		// save the session key calculated for future authentication with a deadline(TODO: future scope)
	authentication_stateful_information :=  model.Authentication_User {
		Username: user.Username,
		SessionKey: "asdasdasdas",
	}
	db := database.ConnectDB()
	db.LogMode(true)
	result := db.Create(&authentication_stateful_information)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User name already exists")
		return 
	}
	defer db.Close()

	
	
}


