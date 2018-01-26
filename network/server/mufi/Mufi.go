package mufi

import (
    "crypto/x509"
    "crypto/tls"
    "io/ioutil"
    . "BotServer/utils/log"
    "BotServer/network/client"
    "BotServer/network/messages"
)

func MufiServer() {
    cert, err := loadX509KeyPair("certs/client.crt", "certs/client.key", "Zog1Ri6AWEV9Oe45")

    if err != nil {
        Error.Fatalf("Unable to load client cert/keys: %s", err)
    }

    rootCACert, err := ioutil.ReadFile("certs/ca.crt")

    if err != nil {
        Error.Fatalf("Unable to open CA cert: %s", err)
    }

    rootCertPool := x509.NewCertPool()
    rootCertPool.AppendCertsFromPEM(rootCACert)

    config := &tls.Config{
        Certificates: []tls.Certificate{cert},
        ClientCAs: rootCertPool,
        InsecureSkipVerify: true,
    }

    listener, err := tls.Listen("tcp", "0.0.0.0:6555", config)

    if err != nil {
        Error.Fatalf("Unable to listen: %s", err)
    }

    defer listener.Close()
    Info.Print("Start listening Mufibot")

    for {
        conn, err := listener.Accept()
        if err != nil {
            Error.Printf("Accept error: %s", err)
            continue
        }

        Info.Print("New Mufi bot")

        c := client.NewClient(conn)
        c.Send(&messages.BakeryHelloConnectMessage{"2.45.4.131059311.0"})
        go c.Receive()
    }
}
