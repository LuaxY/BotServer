package main

import (
    "os"
    "io"
    "io/ioutil"
    "sync"
    "BotServer/network/server/mufi"
    "BotServer/network/server/swift"
    . "BotServer/utils/log"
)

func main() {
    logFile, _ := os.OpenFile("logs/botserver.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr, io.MultiWriter(logFile, os.Stdout))
    //Init(os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stdout)

    Info.Print("Starting server...")

    var wg sync.WaitGroup

    wg.Add(2)

    go func() {
        defer wg.Done()
        mufi.MufiServer("0.0.0.0:6555")
    }()

    go func() {
        defer wg.Done()
        swift.SwiftServer("0.0.0.0:5557")
    }()

    wg.Wait()
}
