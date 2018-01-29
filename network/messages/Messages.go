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
    reader.ReadBytes(length)
}

///////////////////////////////////////////////

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

///////////////////////////////////////////////

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
    writer.WriteBytes([]byte{0x42,0x76,0xED,0x4C,0xD9,0xF1,0x80,0x00,0x00,0x00,0x00,0x77,0x00}) // date & more
}

func (msg *BakeryIdentificationSuccessMessage) Unpack(reader io.IBinaryReader, length uint32) {
    reader.ReadBytes(length)
}

///////////////////////////////////////////////

type BakeryAddAccountMessage struct {
    Account string
}

func (msg *BakeryAddAccountMessage) ID() int {
    return 16007
}

func (msg *BakeryAddAccountMessage) GetName() string {
    return "BakeryAddAccountMessage"
}

func (msg *BakeryAddAccountMessage) Pack(writer io.IBinaryWriter) {

}

func (msg *BakeryAddAccountMessage) Unpack(reader io.IBinaryReader, length uint32) {
    reader.ReadUInt()
    msg.Account, _ = reader.ReadUTF()
}

///////////////////////////////////////////////

type BakeryRawDataMessage struct {
}

func (msg *BakeryRawDataMessage) ID() int {
    return 16001
}

func (msg *BakeryRawDataMessage) GetName() string {
    return "BakeryRawDataMessage"
}

func (msg *BakeryRawDataMessage) Pack(writer io.IBinaryWriter) {

}

func (msg *BakeryRawDataMessage) Unpack(reader io.IBinaryReader, length uint32) {
    reader.ReadBytes(length)
}

///////////////////////////////////////////////

type SwiftPingMessage struct {
}

func (msg *SwiftPingMessage) ID() int {
    return 999
}

func (msg *SwiftPingMessage) GetName() string {
    return "SwiftPingMessage"
}

func (msg *SwiftPingMessage) Pack(writer io.IBinaryWriter) {

}

func (msg *SwiftPingMessage) Unpack(reader io.IBinaryReader, length uint32) {
    reader.ReadBytes(length)
}

///////////////////////////////////////////////

type SwiftPongMessage struct {
}

func (msg *SwiftPongMessage) ID() int {
    return 998
}

func (msg *SwiftPongMessage) GetName() string {
    return "SwiftPongMessage"
}

func (msg *SwiftPongMessage) Pack(writer io.IBinaryWriter) {

}

func (msg *SwiftPongMessage) Unpack(reader io.IBinaryReader, length uint32) {
    reader.ReadBytes(length)
}

///////////////////////////////////////////////

type SwiftIdentificationMessage struct {
    Login string
    Password string
}

func (msg *SwiftIdentificationMessage) ID() int {
    return 666
}

func (msg *SwiftIdentificationMessage) GetName() string {
    return "SwiftIdentificationMessage"
}

func (msg *SwiftIdentificationMessage) Pack(writer io.IBinaryWriter) {

}

func (msg *SwiftIdentificationMessage) Unpack(reader io.IBinaryReader, length uint32) {
    msg.Login, _ = reader.ReadUTF()
    msg.Password, _ = reader.ReadUTF()
}

///////////////////////////////////////////////

type SwiftIdentificationSuccessMessage struct {
    Nickname string
    FullAccess bool
    Motd string
}

func (msg *SwiftIdentificationSuccessMessage) ID() int {
    return 667
}

func (msg *SwiftIdentificationSuccessMessage) GetName() string {
    return "SwiftIdentificationSuccessMessage"
}

func (msg *SwiftIdentificationSuccessMessage) Pack(writer io.IBinaryWriter) {
    writer.WriteBool(true) // State
    writer.WriteUTF(msg.Nickname) // Pseudo
    writer.WriteInt(0) // Token
    writer.WriteBool(msg.FullAccess) // Full Access
    writer.WriteBool(false) // Instance Full
    writer.WriteUTF(" " + msg.Motd) // Message of the day (number of bots)

    writer.WriteVarInt(11) // Number of options
    writer.WriteUTF("desktopAccess")
    writer.WriteUTF("touchAccess")
    writer.WriteUTF("optiFight")
    writer.WriteUTF("optiControl")
    writer.WriteUTF("craftOther")
    writer.WriteUTF("sellOther")
    writer.WriteUTF("floodOther")
    writer.WriteUTF("secondOther")
    writer.WriteUTF("optiProtection")
    writer.WriteUTF("eliteSupport")
    writer.WriteUTF("mountAcces")

    // Expirations
    writer.WriteVarInt(11) // Number of expiration
    for i := 0; i < 11; i++ {
        writer.WriteLong(0)
    }
}

func (msg *SwiftIdentificationSuccessMessage) Unpack(reader io.IBinaryReader, length uint32) {
    reader.ReadBytes(length)
}

///////////////////////////////////////////////

type SelectedServerDataCustomMessage struct {
    Username string
    Ticket []byte
}

func (msg *SelectedServerDataCustomMessage) ID() int {
    return 747
}

func (msg *SelectedServerDataCustomMessage) GetName() string {
    return "SelectedServerDataCustomMessage"
}

func (msg *SelectedServerDataCustomMessage) Pack(writer io.IBinaryWriter) {

}

func (msg *SelectedServerDataCustomMessage) Unpack(reader io.IBinaryReader, length uint32) {
    msg.Username, _ = reader.ReadUTF()
    ticketLength, _ := reader.ReadVarInt()
    msg.Ticket, _ = reader.ReadBytes(uint32(ticketLength))
}

///////////////////////////////////////////////

type AuthenticationTicketCustomMessage struct {
    Account string
    Ticket string
}

func (msg *AuthenticationTicketCustomMessage) ID() int {
    return 674
}

func (msg *AuthenticationTicketCustomMessage) GetName() string {
    return "AuthenticationTicketCustomMessage"
}

func (msg *AuthenticationTicketCustomMessage) Pack(writer io.IBinaryWriter) {
    writer.WriteUTF(msg.Account)
    writer.WriteUTF(msg.Ticket)
}

func (msg *AuthenticationTicketCustomMessage) Unpack(reader io.IBinaryReader, length uint32) {
    reader.ReadBytes(length)
}

///////////////////////////////////////////////

type SwiftStopBotMessage struct {

}

func (msg *SwiftStopBotMessage) ID() int {
    return 777
}

func (msg *SwiftStopBotMessage) GetName() string {
    return "SwiftStopBotMessage"
}

func (msg *SwiftStopBotMessage) Pack(writer io.IBinaryWriter) {

}

func (msg *SwiftStopBotMessage) Unpack(reader io.IBinaryReader, length uint32) {
    reader.ReadBytes(length)
}

///////////////////////////////////////////////

type CheckIntegrityMessage struct {

}

func (msg *CheckIntegrityMessage) ID() int {
    return 6372
}

func (msg *CheckIntegrityMessage) GetName() string {
    return "CheckIntegrityMessage"
}

func (msg *CheckIntegrityMessage) Pack(writer io.IBinaryWriter) {
    writer.WriteBytes([]byte{0x01,0x00})
}

func (msg *CheckIntegrityMessage) Unpack(reader io.IBinaryReader, length uint32) {
    reader.ReadBytes(length)
}