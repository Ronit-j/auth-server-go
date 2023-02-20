package main

import (
	"net/http"
	"github.com/auth-server-go/controller"
)

func main() {
	sign_up_handler := http.HandlerFunc(controller.Handle_sign_up_request)
	login_handler := http.HandlerFunc(controller.Handle_login_request)
	// authentication_handler := http.HandlerFunc(controller.Handle_authentication_request)
	profile_handler := http.HandlerFunc(controller.Handle_Profile_request)
	http.Handle("/login", login_handler)
	http.Handle("/signup", sign_up_handler)
	http.Handle("/profile", profile_handler)
	http.ListenAndServe(":8080", nil)
}