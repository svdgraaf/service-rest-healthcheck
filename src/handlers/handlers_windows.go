package handlers

import (
    "fmt"
    "net/http"
    "log"

    "github.com/gorilla/mux"

    "golang.org/x/sys/windows"
    "golang.org/x/sys/windows/svc/mgr"
)

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
      defer m.Disconnect()
      return
    }

    log.Println("Connected")
    s, err := m.OpenService(service_name)
    if err != nil {
      log.Println("Service not found")
      w.WriteHeader(http.StatusNotFound)
      defer m.Disconnect()
      return
    }

    state, err := s.Query()
    if err != nil {
      log.Println("Service is in unknown state")
      w.WriteHeader(http.StatusInternalServerError)
      defer m.Disconnect()
      return
    }

    log.Println("Service status:", state)
    if state.State == windows.SERVICE_RUNNING {
      fmt.Fprintf(w, "Service is running: %d", state)
    } else {
      w.WriteHeader(http.StatusInternalServerError)
     	w.Write([]byte("Service is not running"))
    }
}
