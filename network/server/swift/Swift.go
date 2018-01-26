package swift

import (
    "net"
    . "BotServer/utils/log"
    "BotServer/network/client"
)

func SwiftServer() {
    listener, err := net.Listen("tcp", "0.0.0.0:5557")

    if err != nil {
        Error.Fatalf("Unable to listen: %s", err)
    }

    defer listener.Close()
    Info.Print("Start listening Swiftbot")

    for {
        conn, err := listener.Accept()

        if err != nil {
            Error.Printf("Accept error: %s", err)
            continue
        }

        Info.Print("New Swift bot")

        c := client.NewClient(conn)
        go c.Receive()
    }
}

