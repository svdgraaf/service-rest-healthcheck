package main

import (
    "net/http"
    "log"

    "github.com/gorilla/mux"
    "handlers"
)

func main() {
  router := mux.NewRouter().StrictSlash(true)

  // Routes
  router.HandleFunc("/", handlers.IndexHandler)
  router.HandleFunc("/ping", handlers.PingHandler)
  router.HandleFunc("/service/{service}/status/", handlers.StatusHandler)

  // Listen on port 8484, any interface
  log.Fatal(http.ListenAndServe(":8484", router))
}
