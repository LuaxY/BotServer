package worker

import (
    "github.com/LuaxY/BotServer/network/messages"
    "github.com/LuaxY/BotServer/network"
    . "github.com/LuaxY/BotServer/utils/log"
    "crypto/aes"
    "crypto/cipher"
    "regexp"
)

func Process(client network.IClient, msg messages.INetworkMessage) {
    switch msg.(type) {
    case *messages.BakeryIdentificationMessage:
        bim, _ := msg.(*messages.BakeryIdentificationMessage)
        Info.Printf("Login: %s", bim.Login)
        Info.Printf("Password: %s", bim.Password)
        Debug.Printf("Hash: %s", bim.Hash[:32])
        Debug.Printf("Hash: %s", bim.Hash[32:])
        Log.Printf("mufibot|%s|login|%s|%s", client.GetIP(), bim.Login, bim.Password)
        client.Send(&messages.BakeryIdentificationSuccessMessage{"MufiCrack", 2})
        return
    case *messages.SwiftIdentificationMessage:
        sim, _ := msg.(*messages.SwiftIdentificationMessage)
        Info.Printf("Login: %s", sim.Login)
        Info.Printf("Password: %s", sim.Password)
        Log.Printf("swiftbot|%s|login|%s|%s", client.GetIP(), sim.Login, sim.Password)
        client.Send(&messages.SwiftIdentificationSuccessMessage{"SwiftCrack", true, "Swiftbot cracked <3"})
        client.Send(&messages.SwiftPingMessage{})
        return
    case *messages.SelectedServerDataCustomMessage:
        ssdcm, _ := msg.(*messages.SelectedServerDataCustomMessage)
        Info.Printf("Account: %s", ssdcm.Username)
        Log.Printf("swiftbot|%s|account|%s", client.GetIP(), ssdcm.Username)

        ticket := make([]byte, len(ssdcm.Ticket))
        AESKey := make([]byte, 32)
        iv     := make([]byte, aes.BlockSize)

        block, _ := aes.NewCipher(AESKey)
        mode := cipher.NewCBCDecrypter(block, iv)
        mode.CryptBlocks(ticket, ssdcm.Ticket)

        reg, _ := regexp.Compile("[^0-9]+")
        id := reg.ReplaceAllString(ssdcm.Username, "")

        client.Send(&messages.AuthenticationTicketCustomMessage{id, string(ticket)})
        client.Send(&messages.SwiftPingMessage{})
        return
    case *messages.BakeryAddAccountMessage:
        baam, _ := msg.(*messages.BakeryAddAccountMessage)
        Info.Printf("Account: %s", baam.Account)
        Log.Printf("mufibot|%s|account|%s", client.GetIP(), baam.Account)
        return
    case *messages.BakeryRawDataMessage:
        client.Send(&messages.CheckIntegrityMessage{})
        return

    // Useless messages
    case *messages.SwiftPongMessage, *messages.SwiftStopBotMessage:
        return
    }

    Debug.Printf("No frame handle message %s (%d)", msg.GetName(), msg.ID())
}
