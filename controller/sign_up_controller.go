package controller

import (
	"net/http"
)


func Handle_sign_up_request(w http.ResponseWriter, r *http.Request) {

	// Implement sign up function with db
	w.WriteHeader(200)
	return
}
