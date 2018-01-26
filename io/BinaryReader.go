package io

import (
    "io"
    "encoding/binary"
    "errors"
    "bytes"
    "unsafe"
)

type IBinaryReader interface {
    BytesAvailable() uint32
    ReadBool() (bool, error)
    ReadByte() (int8, error)
    ReadUByte() (uint8, error)
    ReadShort() (int16, error)
    ReadUShort() (uint16, error)
    ReadInt() (int32, error)
    ReadUInt() (uint32, error)
    ReadLong() (int64, error)
    ReadULong() (uint64, error)
    ReadFloat() (float32, error)
    ReadDouble() (float64, error)
    ReadUTF() (string, error)
    ReadUTFBytes(len uint32) (string, error)
    ReadBytes(len uint32) ([]byte, error)
    ReadVarShort() (int16, error)
    ReadVarUShort() (uint16, error)
    ReadVarInt() (int32, error)
    ReadVarUInt() (uint32, error)
    ReadVarLong() (int64, error)
    ReadVarULong() (uint64, error)
}

type binaryReader struct {
    reader         io.Reader
    bytesAvailable uint32
}

var ErrReaderMalformedVar = errors.New("malformed variable length integer")

func NewBinaryReader(data []byte, length uint32) IBinaryReader {
    return &binaryReader{bytes.NewReader(data), length}
}

func (r *binaryReader) BytesAvailable() uint32 {
    return r.bytesAvailable
}

func (r *binaryReader) read(x interface{}) error {
    return binary.Read(r.reader, binary.BigEndian, x)
}

func (r *binaryReader) ReadBool() (bool, error) {
    b, err := r.ReadUByte()
    if err != nil {
        return false, err
    }
    return b != 0, nil
}

func (r *binaryReader) ReadByte() (int8, error) {
    b, err := r.ReadUByte()
    if err != nil {
        return 0, err
    }
    return int8(b), nil
}

func (r *binaryReader) ReadUByte() (uint8, error) {
    var b uint8
    if err := r.read(&b); err != nil {
        return 0, err
    }
    r.bytesAvailable -= uint32(unsafe.Sizeof(b))
    return b, nil
}

func (r *binaryReader) ReadShort() (int16, error) {
    var b int16
    if err := r.read(&b); err != nil {
        return 0, err
    }
    r.bytesAvailable -= uint32(unsafe.Sizeof(b))
    return b, nil
}

func (r *binaryReader) ReadUShort() (uint16, error) {
    var b uint16
    if err := r.read(&b); err != nil {
        return 0, err
    }
    r.bytesAvailable -= uint32(unsafe.Sizeof(b))
    return b, nil
}

func (r *binaryReader) ReadInt() (int32, error) {
    var b int32
    if err := r.read(&b); err != nil {
        return 0, err
    }
    r.bytesAvailable -= uint32(unsafe.Sizeof(b))
    return b, nil
}

func (r *binaryReader) ReadUInt() (uint32, error) {
    var b uint32
    if err := r.read(&b); err != nil {
        return 0, err
    }
    r.bytesAvailable -= uint32(unsafe.Sizeof(b))
    return b, nil
}

func (r *binaryReader) ReadLong() (int64, error) {
    var b int64
    if err := r.read(&b); err != nil {
        return 0, err
    }
    r.bytesAvailable -= uint32(unsafe.Sizeof(b))
    return b, nil
}

func (r *binaryReader) ReadULong() (uint64, error) {
    var b uint64
    if err := r.read(&b); err != nil {
        return 0, err
    }
    r.bytesAvailable -= uint32(unsafe.Sizeof(b))
    return b, nil
}

func (r *binaryReader) ReadFloat() (float32, error) {
    var b float32
    if err := r.read(&b); err != nil {
        return 0, err
    }
    r.bytesAvailable -= uint32(unsafe.Sizeof(b))
    return b, nil
}

func (r *binaryReader) ReadDouble() (float64, error) {
    var b float64
    if err := r.read(&b); err != nil {
        return 0, err
    }
    r.bytesAvailable -= uint32(unsafe.Sizeof(b))
    return b, nil
}

func (r *binaryReader) ReadUTF() (string, error) {
    len, err := r.ReadUShort()
    if err != nil {
        return "", err
    }
    return r.ReadUTFBytes(uint32(len))
}

func (r *binaryReader) ReadUTFBytes(len uint32) (string, error) {
    buf := make([]byte, len)
    _, err := io.ReadFull(r.reader, buf)
    if err != nil {
        return "", err
    }
    r.bytesAvailable -= len
    return string(buf), nil
}

func (r *binaryReader) ReadBytes(len uint32) ([]byte, error) {
    buf := make([]byte, len)
    _, err := io.ReadFull(r.reader, buf)
    if err != nil {
        return nil, err
    }
    r.bytesAvailable -= len
    return buf, nil
}

func (r *binaryReader) ReadVarShort() (int16, error) {
    v, err := r.readVar(16)
    if err != nil {
        return 0, err
    }
    return int16(v), err
}

func (r *binaryReader) ReadVarUShort() (uint16, error) {
    v, err := r.readVar(16)
    if err != nil {
        return 0, err
    }
    return uint16(v), err
}

func (r *binaryReader) ReadVarInt() (int32, error) {
    v, err := r.readVar(32)
    if err != nil {
        return 0, err
    }
    return int32(v), err
}

func (r *binaryReader) ReadVarUInt() (uint32, error) {
    v, err := r.readVar(32)
    if err != nil {
        return 0, err
    }
    return uint32(v), err
}

func (r *binaryReader) ReadVarLong() (int64, error) {
    v, err := r.readVar(64)
    if err != nil {
        return 0, err
    }
    return int64(v), err
}

func (r *binaryReader) ReadVarULong() (uint64, error) {
    return r.readVar(64)
}

func (r *binaryReader) readVar(bits uint) (uint64, error) {
    var v uint64
    for n := uint(0); n < bits; n += 7 {
        b, err := r.ReadUByte()
        if err != nil {
            return 0, err
        }

        v |= uint64(b&0x7F) << n

        if b&0x80 == 0 {
            return v, nil
        }
    }
    return 0, ErrReaderMalformedVar
}