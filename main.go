package main

import (
    "fmt"
)

func main() {
    fmt.Println("vim-go")

    challenge1_1()
}

func challenge1_1() {
    hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    b64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

    var b64Str string

    b64Str = hexToBase64(hex)

    fmt.Printf("%s", b64Str)
    if b64 != b64Str {
        fmt.Println("Error: challenge1")
    }
}

func hexToBase64(hexStr string) string {
    var base64 string
    hexStrSize := len(hexStr)
    hexStrSlice := []rune(hexStr)

    if hexStrSize % 2 != 0 {
        fmt.Println("error: hexToBase64")
    }

    hexToBytes := hexStrToByteArray(hexStrSlice)

    base64 = bytesToBase64Str(hexToBytes)

    return base64
}

func hexStrToByteArray(hexStrSlice []rune) []byte {
    bytesSlice := make([]byte, 0, len(hexStrSlice)/2)

    for i := 0; i < len(hexStrSlice); i += 2 {
        char1 := hexStrSlice[i]
        char2 := hexStrSlice[i+1]
        byteVal := hexCharToByte(char1) * 16 + hexCharToByte(char2)
        bytesSlice = append(bytesSlice, byteVal)
    }

    return bytesSlice
}

func hexCharToByte(char rune) byte {
    var hexByte byte

    switch {
    case char >= '0' && char <= '9':
        hexByte = byte(char) - byte('0')
    case char >= 'a' && char <= 'f':
        hexByte = byte(char) - byte('a') + 10
    case char >= 'A' && char <= 'F':
        hexByte = byte(char) - byte('A') + 10
    default:
        fmt.Println("error: hexCharToByte")
}

    return hexByte
}

func bytesToBase64Str(bytes []byte) string {
    base64Alphabet := [64]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '/'}
    b64Str := make([]rune, 0, len(bytes)*4/3)

    u32Slice := make([]uint32, 0, len(bytes)/3)
    for i := 0; i < len(bytes); i += 3 {
        val := uint32(bytes[i]) << 16
        val |= uint32(bytes[i+1]) << 8
        val |= uint32(bytes[i+2])
        u32Slice = append(u32Slice, val)
    }

    for i := 0; i < len(u32Slice); i++ {
        for j := 0; j < 4; j++ {
            var idx byte = byte((u32Slice[i] >> (18 - j * 6)) & 0x3F)
            b64Str = append(b64Str, base64Alphabet[idx])
        }
    }

    return string(b64Str)
}
