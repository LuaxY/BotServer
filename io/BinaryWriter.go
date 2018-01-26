package io

import (
    "io"
    "encoding/binary"
    "bytes"
)

type IBinaryWriter interface {
    Data() []byte
    WriteBool(bool) error
    WriteByte(int8) error
    WriteUByte(uint8) error
    WriteShort(int16) error
    WriteUShort(uint16) error
    WriteInt(int32) error
    WriteUInt(uint32) error
    WriteLong(int64) error
    WriteULong(uint64) error
    WriteFloat(float32) error
    WriteDouble(float64) error
    WriteBytes([]byte) error
    WriteUTF(string) error
    WriteUTFBytes(string) error
    WriteVarShort(int16) error
    WriteVarUShort(uint16) error
    WriteVarInt(int32) error
    WriteVarUInt(uint32) error
    WriteVarLong(int64) error
    WriteVarULong(uint64) error
}

type binaryWriter struct {
    w io.Writer
}

func NewBinaryWriter() IBinaryWriter {
    return &binaryWriter{new(bytes.Buffer)}
}

func (w *binaryWriter) Data() []byte {
    buffer, _ := w.w.(*bytes.Buffer)
    return buffer.Bytes()
}

func (w *binaryWriter) write(x interface{}) error {
    return binary.Write(w.w, binary.BigEndian, x)
}

func (w *binaryWriter) writeVar(x uint64) error {
    for x != 0 {
        b := uint8(x & 0x7f)
        x >>= 7
        if x != 0 {
            b |= 0x80
        }
        if err := w.WriteUByte(b); err != nil {
            return err
        }
    }
    return nil
}

func (w *binaryWriter) WriteBool(x bool) error {
    var b uint8
    if x {
        b = 1
    } else {
        b = 0
    }
    return w.write(b)
}

func (w *binaryWriter) WriteByte(x int8) error {
    return w.write(x)
}

func (w *binaryWriter) WriteUByte(x uint8) error {
    return w.write(x)
}

func (w *binaryWriter) WriteShort(x int16) error {
    return w.write(x)
}

func (w *binaryWriter) WriteUShort(x uint16) error {
    return w.write(x)
}

func (w *binaryWriter) WriteInt(x int32) error {
    return w.write(x)
}

func (w *binaryWriter) WriteUInt(x uint32) error {
    return w.write(x)
}

func (w *binaryWriter) WriteLong(x int64) error {
    return w.write(x)
}

func (w *binaryWriter) WriteULong(x uint64) error {
    return w.write(x)
}

func (w *binaryWriter) WriteFloat(x float32) error {
    return w.write(x)
}

func (w *binaryWriter) WriteDouble(x float64) error {
    return w.write(x)
}

func (w *binaryWriter) WriteBytes(x []byte) error {
    return w.write(x)
}

func (w *binaryWriter) WriteUTF(x string) error {
    if err := w.WriteUShort(uint16(len(x))); err != nil {
        return err
    }
    return w.write([]byte(x))
}

func (w *binaryWriter) WriteUTFBytes(x string) error {
    return w.write([]byte(x))
}

func (w *binaryWriter) WriteVarShort(x int16) error {
    return w.writeVar(uint64(x))
}

func (w *binaryWriter) WriteVarUShort(x uint16) error {
    return w.writeVar(uint64(x))
}

func (w *binaryWriter) WriteVarInt(x int32) error {
    return w.writeVar(uint64(x))
}

func (w *binaryWriter) WriteVarUInt(x uint32) error {
    return w.writeVar(uint64(x))
}

func (w *binaryWriter) WriteVarLong(x int64) error {
    return w.writeVar(uint64(x))
}

func (w *binaryWriter) WriteVarULong(x uint64) error {
    return w.writeVar(x)
}