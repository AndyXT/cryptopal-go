package hex

import (
    "fmt"
    "cryptopal-go/pkg/bytes"
)
func ToBase64(hexStr string) string {
    var base64 string
    hexStrSize := len(hexStr)
    hexStrSlice := []rune(hexStr)

    if hexStrSize % 2 != 0 {
        fmt.Println("error: hexToBase64")
    }

    hexToBytes := ToByteArray(hexStrSlice)

    base64 = bytes.ToBase64Str(hexToBytes)

    return base64
}

func ToByteArray(hexStrSlice []rune) []byte {
    bytesSlice := make([]byte, 0, len(hexStrSlice)/2)

    for i := 0; i < len(hexStrSlice); i += 2 {
        char1 := hexStrSlice[i]
        char2 := hexStrSlice[i+1]
        byteVal := CharToByte(char1) * 16 + CharToByte(char2)
        bytesSlice = append(bytesSlice, byteVal)
    }

    return bytesSlice
}

func CharToByte(char rune) byte {
    var hexByte byte

    switch {
    case char >= '0' && char <= '9':
        hexByte = byte(char) - byte('0')
    case char >= 'a' && char <= 'f':
        hexByte = byte(char) - byte('a') + 10
    case char >= 'A' && char <= 'F':
        hexByte = byte(char) - byte('A') + 10
    default:
        fmt.Println("error: CharToByte")
    }

    return hexByte
}
