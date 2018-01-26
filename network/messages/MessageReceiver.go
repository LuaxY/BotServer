package messages

import (
    "reflect"
    "BotServer/io"
)

type MessageReceiver struct {
    messagesTypes map[int]reflect.Type
}

func NewMessageReceiver() *MessageReceiver {
    mr := MessageReceiver{}
    mr.register()
    return &mr
}

func (mr *MessageReceiver) register() {
    mr.messagesTypes = make(map[int]reflect.Type)

    mr.messagesTypes[666] = reflect.TypeOf(SwiftIdentificationMessage{})
    mr.messagesTypes[747] = reflect.TypeOf(SelectedServerDataCustomMessage{})
    mr.messagesTypes[777] = reflect.TypeOf(SwiftStopBotMessage{})
    mr.messagesTypes[998] = reflect.TypeOf(SwiftPongMessage{})
    mr.messagesTypes[16001] = reflect.TypeOf(BakeryRawDataMessage{})
    mr.messagesTypes[16004] = reflect.TypeOf(BakeryIdentificationMessage{})
    mr.messagesTypes[16007] = reflect.TypeOf(BakeryAddAccountMessage{})
}

func (mr *MessageReceiver) Parse(reader io.IBinaryReader, id uint16, length uint32) INetworkMessage {
    messageType := mr.messagesTypes[int(id)]

    if messageType == nil {
        reader.ReadBytes(length)
        return nil
    }

    obj := reflect.New(messageType)
    msg := obj.Interface().(INetworkMessage)
    msg.Unpack(reader, length)

    return msg
}