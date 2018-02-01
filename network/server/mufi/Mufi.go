package mufi

import (
    "crypto/x509"
    "crypto/tls"
    "io/ioutil"
    . "github.com/LuaxY/BotServer/utils/log"
    "github.com/LuaxY/BotServer/network/client"
    "github.com/LuaxY/BotServer/network/messages"
)

func MufiServer(address, dir string) {
    passphrase, err := ioutil.ReadFile(dir + "/certs/passphrase.txt")

    if err != nil {
        Error.Fatalf("Unable to load passphrase file: %s", err)
    }

    cert, err := loadX509KeyPair(dir + "/certs/client.crt", dir + "/certs/client.key", string(passphrase))

    if err != nil {
        Error.Fatalf("Unable to load client cert/keys: %s", err)
    }

    rootCACert, err := ioutil.ReadFile(dir + "/certs/ca.crt")

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

    listener, err := tls.Listen("tcp", address, config)

    if err != nil {
        Error.Fatalf("Unable to listen: %s", err)
    }

    defer listener.Close()
    Info.Printf("Start listening Mufibot on %s", address)

    for {
        conn, err := listener.Accept()
        if err != nil {
            Error.Printf("Accept error: %s", err)
            continue
        }

        Info.Print("New Mufi bot")

        c := client.NewClient(client.MUFIBOT, conn)
        c.Send(&messages.BakeryHelloConnectMessage{"2.45.4.131059311.0"})
        go c.Receive()
    }
}
