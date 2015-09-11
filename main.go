package main

import (
    "fmt"
    "html"
    "log"
    "net/http"

    "github.com/gorilla/mux"

)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", IndexHandler)
    router.HandleFunc("/ping", PingHandler)
    router.HandleFunc("/service/{service}/status/", StatusHandler)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Ok", html.EscapeString(r.URL.Path))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Pong")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    service := vars["service"]
    fmt.Fprintf(w, service)
}
