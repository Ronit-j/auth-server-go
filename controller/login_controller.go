package controller

import (
	"encoding/json"
	"net/http"
	// "bytes"
	"strconv"
	// "fmt"
	// "io/ioutil"
	"github.com/auth-server-go/requests"
	"github.com/auth-server-go/model"
	"github.com/auth-server-go/database"
	"github.com/auth-server-go/utils"
)

func Handle_login_request(w http.ResponseWriter, r *http.Request) {

	// var bodyBytes []byte
	// var err error
	// get username and public key from user
		// extract username and the public key from the request body
	var login requests.Login
	// if r.Body != nil {
	// 	bodyBytes, err = ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		fmt.Printf("Body reading error: %v", err)
	// 		return
	// 	}
	// 	defer r.Body.Close()
	// }

	// fmt.Printf("Headers: %+v\n", r.Header)

	// if len(bodyBytes) > 0 {
	// 	var prettyJSON bytes.Buffer
	// 	if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
	// 		fmt.Printf("JSON parse error: %v", err)
	// 		return
	// 	}
	// 	fmt.Println(string(prettyJSON.Bytes()))
	// } else {
	// 	fmt.Printf("Body: No Body Supplied\n")
	// }

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	decoder := json.NewDecoder(r.Body)
	// fmt.Printf(decoder)
    err1 := decoder.Decode(&login)
    if err1 != nil {
		// fmt.Printf(err1)
        panic(err1)
    }
	
	// Get the password hash of the user from the username
	var user model.User
	user = database.Get_user_security_information(login.Username)
	w.Header().Set("Content-Type", "application/json")
	if strconv.FormatUint(uint64(utils.Hash(login.Password)),10)!=user.Password {
		panic("Password Incorrect")
	}
}


