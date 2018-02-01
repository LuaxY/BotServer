package network

import "github.com/LuaxY/BotServer/network/messages"

type IClient interface {
    Send(messages.INetworkMessage)
    Receive()
    GetIP() string
}
