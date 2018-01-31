package web

import (
    "net/http"
    "github.com/gorilla/mux"
    . "BotServer/utils/log"
)

var version string

func WebServer(address, v string) {
    version = v

    Info.Printf("Start listening Web on %s", address)

    router := mux.NewRouter()
    router.HandleFunc("/", home).Methods("GET")
    http.ListenAndServe(address, router)
}

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello (" + version + ")"))
}
