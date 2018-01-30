package main

import (
    "os"
    "io"
    "io/ioutil"
    "sync"
    "net/http"
    "github.com/gorilla/mux"
    "BotServer/network/server/mufi"
    "BotServer/network/server/swift"
    . "BotServer/utils/log"
)

var version = "untagged"

func main() {
    logFile, _ := os.OpenFile("logs/botserver.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr, io.MultiWriter(logFile, os.Stdout))
    //Init(os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stdout)

    Info.Printf("Starting server (V: %s)", version)

    var wg sync.WaitGroup

    wg.Add(2)

    go func() {
        defer wg.Done()
        mufi.MufiServer("0.0.0.0:6555", os.Args[1])
    }()

    go func() {
        defer wg.Done()
        swift.SwiftServer("0.0.0.0:5557")
    }()

    Info.Printf("Start listening Web on %s", "0.0.0.0:80")

    router := mux.NewRouter()
    router.HandleFunc("/", home).Methods("GET")
    http.ListenAndServe("0.0.0.0:80", router)

    wg.Wait()
}

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello (" + version + ")"))
}
