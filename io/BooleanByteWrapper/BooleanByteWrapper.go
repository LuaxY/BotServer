package BooleanByteWrapper

func GetFlag(b uint8, position uint) bool {
    return b&(1<<position) != 0
}

func SetFlag(b uint8, position uint, v bool) uint8 {
    mask := uint8(1 << position)
    if !v {
        mask = ^mask
        return b & mask
    }
    return b | mask
}