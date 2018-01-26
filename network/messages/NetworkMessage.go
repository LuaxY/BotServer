package messages

import "BotServer/io"

type INetworkMessage interface {
    ID() int
    GetName() string
    Pack(writer io.IBinaryWriter)
    Unpack(reader io.IBinaryReader, length uint32)
}