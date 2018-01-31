package main

import (
    "os"
    "io"
    "io/ioutil"
    "sync"
    "BotServer/network/server/mufi"
    "BotServer/network/server/swift"
    "BotServer/network/server/web"
    . "BotServer/utils/log"
)

var version = "untagged"

func main() {
    logFile, _ := os.OpenFile("logs/botserver.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr, io.MultiWriter(logFile, os.Stdout))
    //Init(os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stdout)

    Info.Printf("Starting server (V: %s)", version)

    var wg sync.WaitGroup

    wg.Add(3)

    go func() {
        defer wg.Done()
        mufi.MufiServer("0.0.0.0:6555", os.Args[1])
    }()

    go func() {
        defer wg.Done()
        swift.SwiftServer("0.0.0.0:5557")
    }()

    go func() {
        defer wg.Done()
        web.WebServer("0.0.0.0:80", version)
    }()

    wg.Wait()
}

