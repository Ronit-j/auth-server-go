package main

import (
	"net/http"
	"github.com/auth-server-go/controller"
	"github.com/auth-server-go/database"
	
	
	
)

func main() {
	// start db connection
	database.ConnectDB()
	sign_up_handler := http.HandlerFunc(controller.Handle_sign_up_request)
	login_handler := http.HandlerFunc(controller.Handle_login_request)
	http.Handle("/login", login_handler)
	http.Handle("/signup", sign_up_handler)
	http.ListenAndServe(":8080", nil)
}