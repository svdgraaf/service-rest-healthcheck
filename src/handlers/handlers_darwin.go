package handlers

import (
    "fmt"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ok")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Pong")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "TODO")
}
