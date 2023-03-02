package controller

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/auth-server-go/loadbalancer/requests"
)
var servers int = 0
var service_registry []string

func Handle_service_registration_request(w http.ResponseWriter, r *http.Request) {
	// get username and public key from user
		// extract username and the public key from the request body
	var registration requests.Registration

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	decoder := json.NewDecoder(r.Body)

    err := decoder.Decode(&registration)
    if err != nil {
        panic(err)
    }
	// Increment counter
	servers += 1
	service_registry = append(service_registry, registration.ServiceURL)
	fmt.Printf(service_registry[0])
}


