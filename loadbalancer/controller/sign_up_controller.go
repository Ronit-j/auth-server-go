package controller

import (
	// "encoding/json"
	"fmt"
	"math/rand"
	// "strconv"
	"net/http"
	// "github.com/auth-server-go/model"
	// "github.com/auth-server-go/database"
	// "github.com/auth-server-go/utils"
)


func Handle_sign_up_request(w http.ResponseWriter, r *http.Request) {
	selected_server := service_registry[rand.Int() % len(service_registry)]
	fmt.Printf(selected_server)
	http.Redirect(w, r, "http://" + selected_server + "/signup", http.StatusPermanentRedirect)
}
