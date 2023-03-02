package main

import (
	"net/http"
	"github.com/auth-server-go/loadbalancer/controller"
)

func main() {
	lb_sign_up_handler := http.HandlerFunc(controller.Handle_sign_up_request)
	lb_login_handler := http.HandlerFunc(controller.Handle_login_request)
	service_registry_handler := http.HandlerFunc(controller.Handle_service_registration_request)
	lb_profile_handler := http.HandlerFunc(controller.Handle_Profile_request)
	http.Handle("/lb/login", lb_login_handler)
	http.Handle("/lb/signup", lb_sign_up_handler)
	http.Handle("/lb/profile", lb_profile_handler)
	http.Handle("/lb/register", service_registry_handler)
	http.ListenAndServe(":8081", nil)
}