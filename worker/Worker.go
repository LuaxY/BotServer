package worker

import (
    "BotServer/network/messages"
    "BotServer/network"
    . "BotServer/utils/log"
    "crypto/aes"
    "crypto/cipher"
    "encoding/hex"
)

func Process(client network.IClient, msg messages.INetworkMessage) {
    switch msg.(type) {
    case *messages.BakeryIdentificationMessage:
        bim, _ := msg.(*messages.BakeryIdentificationMessage)
        Info.Printf("Login: %s", bim.Login)
        Info.Printf("Password: %s", bim.Password)
        Info.Printf("Hash: %s", bim.Hash[:32])
        Info.Printf("Hash: %s", bim.Hash[32:])
        client.Send(&messages.BakeryIdentificationSuccessMessage{"MufiCrack", 2})
        return
    case *messages.SwiftIdentificationMessage:
        sim, _ := msg.(*messages.SwiftIdentificationMessage)
        Info.Printf("Login: %s", sim.Login)
        Info.Printf("Password: %s", sim.Password)
        client.Send(&messages.SwiftIdentificationSuccessMessage{"SwiftCrack", true, "Swiftbot cracked <3"})
        client.Send(&messages.SwiftPingMessage{})
        return
    case *messages.SelectedServerDataCustomMessage:
        ssdcm, _ := msg.(*messages.SelectedServerDataCustomMessage)
        Info.Printf("Account: %s", ssdcm.Username)

        plaintext := make([]byte, len(ssdcm.Ticket))
        block, _ := aes.NewCipher([]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

        iv := make([]byte, aes.BlockSize)
        ciphertext := ssdcm.Ticket

        /*iv := ciphertext[:aes.BlockSize]
        ciphertext = ciphertext[aes.BlockSize:]*/

        mode := cipher.NewCBCEncrypter(block, iv)
        mode.CryptBlocks(plaintext, ciphertext)

        ticket := hex.EncodeToString(plaintext)

        Debug.Printf("TICKET: %X", ssdcm.Ticket)
        Debug.Printf("IV: %X", iv)
        Debug.Printf("CIPHER: %X", ciphertext)
        Debug.Printf("TICKET: %s", ticket)

        client.Send(&messages.AuthenticationTicketCustomMessage{ticket})
        client.Send(&messages.SwiftPingMessage{})
        return
    case *messages.BakeryAddAccountMessage:
        baam, _ := msg.(*messages.BakeryAddAccountMessage)
        Info.Printf("Account: %s", baam.Account)
        return
    case *messages.BakeryRawDataMessage:
        client.Send(&messages.CheckIntegrityMessage{})
        return

    // Useless messages
    case *messages.SwiftPongMessage:
    case *messages.SwiftStopBotMessage:
        return
    }

    Debug.Printf("No frame handle message %s (%d)", msg.GetName(), msg.ID())
}
