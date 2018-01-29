package network

import "BotServer/network/messages"

type IClient interface {
    Send(messages.INetworkMessage)
    Receive()
    GetIP() string
}
