package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

type userProfileLookupReq struct {
    Api_key string
    Verbose bool
    Lookup_by struct {
        Interests struct {
            Interest []string
        }
        Keywords struct {
            Keyword []string
        }
        Categories struct {
            Category []string
        }
        Tags struct {
            Tag map[string]string
        }
    }
}

type userProfileReq struct {
    ApiKey string
    Id     string
}

func printError(msg string, rw http.ResponseWriter) {
    rw.WriteHeader(500)
    fmt.Fprintf(rw, "{\"status\":{\"code\":500,\"description\":\"%s\"}}\n", msg)
}

func handleUserProfileLookup(res http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var u userProfileLookupReq
    err := decoder.Decode(&u)
    if err != nil {
        log.Println(err.Error())
        errorDesc := "Error: unable to parse request"
        printError(errorDesc, res)
        return
    }

    api_key := u.Api_key
    verbose := u.Verbose
    lookup_by := u.Lookup_by
    interests := lookup_by.Interests
    categories := lookup_by.Categories
    tags := lookup_by.Tags
    keywords := lookup_by.Keywords
    log.Println("api_key = ",api_key)
    log.Println("verbose = ",verbose)
    log.Println("lookup_by = ",lookup_by)
    log.Println("interests=", interests)
    log.Println("categories=", categories)
    log.Println("tags=", tags)
    log.Println("keywords=", keywords)



    // no input errors
    filename := "/home/tirto/go/src/github.com/tirto/webserver/userProfileLookup.json"
    if verbose {
        filename = "/home/tirto/go/src/github.com/tirto/webserver/userProfileLookupVerbose.json"
    }

    // TODO: implement isApiKeyValid(api_key)
    if api_key != "dd06058e-f32a-4e11-b14b-85a2f98ea523" {
        errorDesc := "Error: invalid apiKey"
        printError(errorDesc, res)
        return
    }

    buf, err := ioutil.ReadFile(filename)
    if err != nil {
        errorDesc := "Error: unable to lookup userProfile"
        printError(errorDesc, res)
        return
    }
    fmt.Fprintf(res, string(buf))
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

    log.Println("api_key =", u.Api_key)
    log.Println("id =", u.Id)
    api_key := u.Api_key
    id := u.Id
    // TODO: implement isApiKeyValid(api_key)
    if api_key != "dd06058e-f32a-4e11-b14b-85a2f98ea523" {
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
    http.HandleFunc("/AudService/v1/user/profile/lookup", handleUserProfileLookup)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
