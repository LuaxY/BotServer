package messages

import "github.com/LuaxY/BotServer/io"

type INetworkMessage interface {
    ID() int
    GetName() string
    Pack(writer io.IBinaryWriter)
    Unpack(reader io.IBinaryReader, length uint32)
}