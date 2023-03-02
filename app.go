package main

import (
	"net/http"
	"github.com/auth-server-go/controller"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"log"
	"fmt"
)

const LOGIN_URL = "/login"
const SIGNUP_URL = "/signup"
const PROFILE_URL = "/register"
const SERVER_PORT = ":8085"

func register_service(){

	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"ServiceName":  "Ronit1",
		"ServiceUrl": "localhost" + SERVER_PORT,
	})
	 responseBody := bytes.NewBuffer(postBody)
  	//Leverage Go's HTTP Post function to make request
	 resp, err := http.Post("http://localhost:8081/lb/register", "application/json", responseBody)
  	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
  	//Read the response body
	 body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	 }
	sb := string(body)
	fmt.Printf(sb)
	fmt.Printf("hello registered service")

}




func main() {

	go register_service()
	sign_up_handler := http.HandlerFunc(controller.Handle_sign_up_request)
	login_handler := http.HandlerFunc(controller.Handle_login_request)
	profile_handler := http.HandlerFunc(controller.Handle_Profile_request)
	http.Handle(LOGIN_URL, login_handler)
	http.Handle(SIGNUP_URL, sign_up_handler)
	http.Handle(PROFILE_URL, profile_handler)
	http.ListenAndServe(SERVER_PORT, nil)
}