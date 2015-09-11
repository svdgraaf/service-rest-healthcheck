package main

import (
    "net/http"
    "log"

    "github.com/gorilla/mux"
    "handlers"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", handlers.IndexHandler)
    router.HandleFunc("/ping", handlers.PingHandler)
    router.HandleFunc("/service/{service}/status/", handlers.StatusHandler)
    log.Fatal(http.ListenAndServe(":8080", router))
}
