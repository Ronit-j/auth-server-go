package controller

import (
	"fmt"
	"net/http"
)

var (
	username = "abc"
	password = "123"
)

func Handle_login_request(w http.ResponseWriter, r *http.Request) {

	user_name, user_password, ok := r.BasicAuth()
	if !ok {
		fmt.Println("Error parsing basic auth")
		w.WriteHeader(401)
		return
	}
	if user_name != username {
		fmt.Printf("Username provided is incorrect: %s\n", user_name)
		w.WriteHeader(401)
		return
	}
	if user_password != password {
		fmt.Printf("Password provided is incorrect: %s\n", user_password)
		w.WriteHeader(401)
		return
	}
	fmt.Printf("Username: %s\n", user_name)
	fmt.Printf("Password: %s\n", user_password)
	w.WriteHeader(200)
	return
}


