package main

import (
	"fmt"
	"net/http"
)

var (
	username = "abc"
	password = "123"
)

func main() {
	sign_up_handler := http.HandlerFunc(handle_sign_up_request)
	login_handler := http.HandlerFunc(handle_login_request)
	http.Handle("/login", login_handler)
	http.Handle("/signup", sign_up_handler)
	http.ListenAndServe(":8080", nil)
}

func handle_sign_up_request(w http.ResponseWriter, r *http.Request) {

	// Implement sign up function with db
	w.WriteHeader(200)
	return
}

func handle_login_request(w http.ResponseWriter, r *http.Request) {

	user_name, user_password, ok := r.BasicAuth()
	if !ok {
		fmt.Println("Error parsing basic auth")
		w.WriteHeader(401)
		return
	}
	if user_name != username {
		fmt.Printf("Username provided is correct: %s\n", user_name)
		w.WriteHeader(401)
		return
	}
	if user_password != password {
		fmt.Printf("Password provided is correct: %s\n", user_password)
		w.WriteHeader(401)
		return
	}
	fmt.Printf("Username: %s\n", user_name)
	fmt.Printf("Password: %s\n", user_password)
	w.WriteHeader(200)
	return
}
