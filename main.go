package main

import "fmt"

func main() {
    fmt.Println("vim-go")

    challenge1_1()
}

func challenge1_1() {
    hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    b64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

    var b64Str string

    b64Str = hexToBase64(hex)

    if b64 != b64Str {
        fmt.Println("Error")
    }
}

func hexToBase64(hexStr string) string {
    var base64 string
    hexStrSize := len(hexStr)
    hexStrSlice := []rune(hexStr)

    if hexStrSize % 2 != 0 {
        fmt.Println("error")
    }

    hexToBytes := hexStrToByteArray(hexStrSlice)

    base64 = bytesToBase64Str(hexToBytes)

    return base64
}

func hexStrToByteArray(hexStrSlice []rune) []byte {
    var bytesArray []byte

    for i := 0; i < len(hexStrSlice); i += 2 {
        char1 := hexStrSlice[i]
        char2 := hexStrSlice[i+1]

    }

    return bytesArray
}

func hexCharToByte(char rune) byte {
    var hexByte byte

    switch {
    case char >= '0' && char <= '9':
        hexByte = byte(char) - byte('0')
    default:
        fmt.Println("error")
    }

    return hexByte
}

func bytesToBase64Str(bytes []byte) string {
    var b64Str string

    return b64Str
}
