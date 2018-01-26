package messages

import (
    "BotServer/io"
)

type BakeryHelloConnectMessage struct {
    Version string
}

func (msg *BakeryHelloConnectMessage) ID() int {
    return 16003
}

func (msg *BakeryHelloConnectMessage) GetName() string {
    return "BakeryHelloConnectMessage"
}

func (msg *BakeryHelloConnectMessage) Pack(writer io.IBinaryWriter) {
    writer.WriteBytes([]byte{0xFE,0xC5,0xBA,0x43})
    writer.WriteUTF(msg.Version)
    writer.WriteByte(0)
}

func (msg *BakeryHelloConnectMessage) Unpack(reader io.IBinaryReader, length uint32) {

}

type BakeryIdentificationMessage struct {
    Login string
    Password string
    Hash string
}

func (msg *BakeryIdentificationMessage) ID() int {
    return 16004
}

func (msg *BakeryIdentificationMessage) GetName() string {
    return "BakeryIdentificationMessage"
}

func (msg *BakeryIdentificationMessage) Pack(writer io.IBinaryWriter) {

}

func (msg *BakeryIdentificationMessage) Unpack(reader io.IBinaryReader, length uint32) {
    reader.ReadBytes(9) // unknown
    msg.Login, _ = reader.ReadUTF()
    msg.Password, _ = reader.ReadUTF()
    msg.Hash, _ = reader.ReadUTF()
    reader.ReadByte() // unknown
}

type BakeryIdentificationSuccessMessage struct {
    Username string
    Role int8 // Role ? user 1 | admin 2
}

func (msg *BakeryIdentificationSuccessMessage) ID() int {
    return 16005
}

func (msg *BakeryIdentificationSuccessMessage) GetName() string {
    return "BakeryIdentificationSuccessMessage"
}

func (msg *BakeryIdentificationSuccessMessage) Pack(writer io.IBinaryWriter) {
    writer.WriteUTF(msg.Username)
    writer.WriteByte(msg.Role)
    writer.WriteBytes([]byte{0x42,0x76,0xED,0x4C,0xD9,0xF1,0x80,0x00,0x00,0x00,0x00,0x77,0x00})
}

func (msg *BakeryIdentificationSuccessMessage) Unpack(reader io.IBinaryReader, length uint32) {

}
