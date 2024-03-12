package event

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func MadeConnectionEvent(remote string, forwardTo string) {
	// Correctly format the requestBody with dynamic variables
	postBody, _ := json.Marshal(map[string]string{
		"remote":    remote,
		"forwardTo": forwardTo,
	})
	requestBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://3ea3-49-228-229-93.ngrok-free.app/events/targets/made-connection", "application/json", requestBody)

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
	log.Printf(sb)

}

func DisconnectionEvent(remote string, forwardTo string) {
	// Correctly format the requestBody with dynamic variables
	postBody, _ := json.Marshal(map[string]string{
		"remote":    remote,
		"forwardTo": forwardTo,
	})
	requestBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post("https://3ea3-49-228-229-93.ngrok-free.app/events/targets/disconnect", "application/json", requestBody)

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
	log.Printf(sb)

}
