package worker

import (
    "BotServer/network/messages"
    "BotServer/network"
    . "BotServer/utils/log"
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
        //ssdcm, _ := msg.(*messages.SelectedServerDataCustomMessage)
        client.Send(&messages.SelectedServerDataAnswerMessage{})
        client.Send(&messages.SwiftPingMessage{})
        return
    case *messages.BakeryAddAccountMessage:
        baam, _ := msg.(*messages.BakeryAddAccountMessage)
        Info.Printf("Account: %s", baam.Account)
        return
    case *messages.BakeryRawDataMessage:
        client.Send(&messages.CheckIntegrityMessage{})
        return
    }
    Debug.Printf("No frame handle message %s (%d)", msg.GetName(), msg.ID())
}
