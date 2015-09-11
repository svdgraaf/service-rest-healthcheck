package main

import (
    "fmt"
    "html"
    "net/http"
    "log"

    "github.com/gorilla/mux"
    "golang.org/x/sys/windows/svc/mgr"
    "golang.org/x/sys/windows"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", IndexHandler)
    router.HandleFunc("/ping", PingHandler)
    router.HandleFunc("/service/{service}/status/", StatusHandler)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ok")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Pong")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    service_name := vars["service"]
    log.Println("Checking:", service_name)

    m, err := mgr.Connect()
    if err != nil {
      log.Println("Bad Connection")
      w.WriteHeader(http.StatusInternalServerError)
    }
    log.Println("Connected")
    s, err := m.OpenService(service_name)
    if err != nil {
      log.Println("Service not found")
      w.WriteHeader(http.StatusNotFound)
    } else {
      state, err := s.Query()
      if err != nil {
        log.Println("Service is in unknown state")
        w.WriteHeader(http.StatusInternalServerError)
      } else {
        log.Println("Service status:", state)
        if state.State == windows.SERVICE_RUNNING {
          fmt.Fprintf(w, "Service is running: %d", state)
        } else {
          w.WriteHeader(http.StatusInternalServerError)
         	w.Write([]byte("Service is not running"))
        }
      }
    }
    defer m.Disconnect()
}
