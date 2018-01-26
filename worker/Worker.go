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
    }

    Debug.Printf("No frame handle message %s (%d)", msg.GetName(), msg.ID())
}
