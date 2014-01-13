package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type userProfileReq struct {
	ApiKey string
	Id     string
}

func printError(msg string, rw http.ResponseWriter) {
	rw.WriteHeader(500)
	fmt.Fprintf(rw, "{\"status\":{\"code\":500,\"description\":\"%s\"}}\n", msg)
}

func handleUserProfile(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var u userProfileReq
	err := decoder.Decode(&u)
	if err != nil {
		log.Println(err.Error())
		errorDesc := "Error: unable to parse request"
		printError(errorDesc, res)
		return
	}

	log.Println("apiKey =", u.ApiKey)
	log.Println("id =", u.Id)
	apiKey := u.ApiKey
	id := u.Id
	// TODO: implement isApiKeyValid(apiKey)
	if apiKey != "dd06058e-f32a-4e11-b14b-85a2f98ea523" {
		errorDesc := "Error: invalid apiKey"
		printError(errorDesc, res)
		return
	}

	// TODO: implement lookupUserProfile(id)
	if id != "abcxyz123" {
		errorDesc := "Error: unable to parse request"
		printError(errorDesc, res)
		return
	}

	// no input errors
	buf, err := ioutil.ReadFile("/home/tirto/go/src/github.com/tirto/webserver/userProfile.json")
	if err != nil {
		errorDesc := "Error: unable to lookup userProfile"
		printError(errorDesc, res)
		return
	}
	fmt.Fprintf(res, string(buf))
}

// main function
func main() {
	http.HandleFunc("/AudService/v1/user/profile", handleUserProfile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
