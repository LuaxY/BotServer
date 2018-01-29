package client

import (
    "net"
    "BotServer/io"
    "BotServer/worker"
    "BotServer/network"
    "BotServer/network/messages"
    . "BotServer/utils/log"

    _ "fmt"
    _ "encoding/hex"
    "time"
)

const (
    MUFIBOT = 1 << iota
    SWIFTBOT
)

type Client struct {
    botType              int
    socket               net.Conn
    parser               *messages.MessageReceiver
    inputBuffer          []byte
    mufiBotID            uint32
    splittedPacket       bool
    splittedPacketHeader uint16
    splittedPacketID     uint16
    splittedPacketLength uint32
}

func NewClient(botType int, conn net.Conn) network.IClient {
    var c Client
    c.botType = botType
    c.socket = conn
    c.parser = messages.NewMessageReceiver()
    return &c
}

func (c *Client) Send(msg messages.INetworkMessage) {
    writer := io.NewBinaryWriter()
    msg.Pack(writer)

    data := writer.Data()

    writer = io.NewBinaryWriter()

    if c.botType == MUFIBOT {
        writer.WriteUInt(c.mufiBotID) // Unknown data for Bakery
    }

    typeLen := computeTypeLen(len(data))
    writer.WriteShort(subComputeStaticHeader(msg.ID(), typeLen))

    switch typeLen {
    case 1:
        writer.WriteByte(int8(len(data)))
    case 2:
        writer.WriteShort(int16(len(data)))
    case 3:
        high := (len(data) >> 16) & 255
        low := len(data) & 65535
        writer.WriteByte(int8(high))
        writer.WriteShort(int16(low))
    }

    if len(data) > 0 {
        writer.WriteBytes(data)
    }

    //fmt.Printf("%s", hex.Dump(writer.Data()))

    time.Sleep(500 * time.Millisecond)
    n, err := c.socket.Write(writer.Data())

    if err != nil {
        Error.Printf("%s", err)
    }

    Debug.Printf("[SND] %s (%d) %d bytes", msg.GetName(), msg.ID(), n)
}

func (c *Client) Receive() {
    buffer := make([]byte, 4096)

    for {
        length, err := c.socket.Read(buffer)

        //fmt.Printf("%s", hex.Dump(buffer[:length]))

        if err != nil {
            //Error.Printf("Read error: %s", err)
            c.socket.Close()
            return
        }

        Debug.Printf("Receive: %d bytes", length)

        reader := io.NewBinaryReader(buffer, uint32(length))

        for reader.BytesAvailable() > 0 {
            msg := c.lowReceive(reader)

            if msg == nil {
                break
            }

            worker.Process(c, msg)
        }
    }
}

func (c *Client) lowReceive(reader io.IBinaryReader) messages.INetworkMessage {
    if !c.splittedPacket {
        if c.botType == MUFIBOT {
            if reader.BytesAvailable() < 4 {
                return nil
            }

            c.mufiBotID, _ = reader.ReadUInt()
        }

        if reader.BytesAvailable() < 2 {
            return nil
        }

        header, _ := reader.ReadUShort()

        id := getMessageID(header)

        Debug.Printf("Message ID: %d", id)

        if reader.BytesAvailable() < uint32(header&3) {
            c.splittedPacketHeader = header
            c.splittedPacketID = id
            c.splittedPacket = true
            c.inputBuffer, _ = reader.ReadBytes(reader.BytesAvailable())
            return nil
        }

        length := readMessageLength(header, reader)

        Debug.Printf("Message length: %d", length)

        if reader.BytesAvailable() < length {
            c.splittedPacketID = id
            c.splittedPacketLength = length
            c.splittedPacket = true
            c.inputBuffer, _ = reader.ReadBytes(reader.BytesAvailable())
            return nil
        }

        msg := c.parser.Parse(reader, id, length)

        if msg != nil {
            Debug.Printf("[RCV] %s (%d) %d bytes", msg.GetName(), msg.ID(), length)
        }

        return msg
    } else {
        if c.botType == MUFIBOT {
            // TODO: read unknown header for Bakery
        }

        if c.splittedPacketHeader != 0 {
            c.splittedPacketLength = readMessageLength(c.splittedPacketHeader, reader)
            c.splittedPacketHeader = 0
        }

        if reader.BytesAvailable()+uint32(len(c.inputBuffer)) >= c.splittedPacketLength {
            data, _ := reader.ReadBytes(reader.BytesAvailable())
            c.inputBuffer = append(c.inputBuffer, data...)
            reader := io.NewBinaryReader(c.inputBuffer, uint32(len(c.inputBuffer)))

            msg := c.parser.Parse(reader, c.splittedPacketID, c.splittedPacketLength)

            if msg != nil {
                Debug.Printf("[RCV] %s (%d) %d bytes", msg.GetName(), msg.ID(), c.splittedPacketLength)
            }

            c.splittedPacketHeader = 0
            c.splittedPacketID = 0
            c.splittedPacketLength = 0
            c.splittedPacket = false
            c.inputBuffer = []byte{}

            return msg
        } else {
            data, _ := reader.ReadBytes(reader.BytesAvailable())
            c.inputBuffer = append(c.inputBuffer, data...)
            return nil
        }
    }
}

func getMessageID(header uint16) uint16 {
    return header >> 2
}

func readMessageLength(header uint16, reader io.IBinaryReader) uint32 {
    byteLenDynamicHeader := header & 3
    messageLength := uint32(0)

    switch byteLenDynamicHeader {
    case 1:
        length, _ := reader.ReadUByte()
        messageLength = uint32(length)
    case 2:
        length, _ := reader.ReadUShort()
        messageLength = uint32(length)
    case 3:
        p1, _ := reader.ReadByte()
        p2, _ := reader.ReadByte()
        p3, _ := reader.ReadByte()
        messageLength = uint32(((int(p1) & 0xFF) << 16) + ((int(p2) & 0xFF) << 8) + (int(p3) & 0xFF))
    }

    return messageLength
}

func subComputeStaticHeader(msgId int, typeLen int) int16 {
    return int16((msgId << 2) | typeLen)
}

func computeTypeLen(length int) int {
    if length > 65535 {
        return 3
    }
    if length > 255 {
        return 2
    }
    if length > 0 {
        return 1
    }
    return 0
}