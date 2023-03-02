package controller

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"math/rand"
)

func Handle_login_request(w http.ResponseWriter, r *http.Request) {
	
	// select the service to send this request to
	var bodyBytes []byte
	var err error
	selected_server := service_registry[rand.Int() % len(service_registry)]
	fmt.Printf(selected_server)

	// send the request

	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Body reading error: %v", err)
			return
		}
		defer r.Body.Close()
	}

	fmt.Printf("Headers: %+v\n", r.Header)

	if len(bodyBytes) > 0 {
		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
			fmt.Printf("JSON parse error: %v", err)
			return
		}
		fmt.Println(string(prettyJSON.Bytes()))
	} else {
		fmt.Printf("Body: No Body Supplied\n")
	}

	http.Redirect(w, r, "http://" + selected_server + "/login", http.StatusPermanentRedirect)




	// return the response

}


