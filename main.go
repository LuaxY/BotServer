package main

import (
    //"log"
    "sync"
    "BotServer/network/server/mufi"
    "BotServer/network/server/swift"
    . "BotServer/utils/log"
    _ "io/ioutil"
    "os"
)

func main() {
    //Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
    Init(os.Stdout, os.Stdout, os.Stdout, os.Stdout)

    Info.Print("Starting server...")

    var wg sync.WaitGroup

    wg.Add(2)

    go func() {
        defer wg.Done()
        mufi.MufiServer()
    }()

    go func() {
        defer wg.Done()
        swift.SwiftServer()
    }()

    wg.Wait()
}
