package web

import (
    "net/http"
    "github.com/gorilla/mux"
    . "github.com/LuaxY/BotServer/utils/log"
)

var version string

func WebServer(address, dir, v string) {
    version = v

    Info.Printf("Start listening Web on %s", address)

    router := mux.NewRouter()
    router.HandleFunc("/version", serverVersion).Methods("GET")
    router.PathPrefix("/").Handler(http.FileServer(http.Dir(dir + "/static/")))
    http.ListenAndServe(address, router)
}

func serverVersion(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(version))
}
