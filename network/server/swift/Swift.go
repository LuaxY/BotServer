package swift

import (
    "net"
    . "github.com/LuaxY/BotServer/utils/log"
    "github.com/LuaxY/BotServer/network/client"
    "github.com/LuaxY/BotServer/network/messages"
)

func SwiftServer(address string) {
    listener, err := net.Listen("tcp", address)

    if err != nil {
        Error.Fatalf("Unable to listen: %s", err)
    }

    defer listener.Close()
    Info.Printf("Start listening Swiftbot on %s", address)

    for {
        conn, err := listener.Accept()

        if err != nil {
            Error.Printf("Accept error: %s", err)
            continue
        }

        Info.Print("New Swift bot")

        c := client.NewClient(client.SWIFTBOT, conn)
        c.Send(&messages.SwiftPingMessage{})
        go c.Receive()
    }
}

