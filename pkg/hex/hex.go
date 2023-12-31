package hex

import (
    "cryptopal-go/pkg/bytes"
    "encoding/base64"
    hx "encoding/hex"
    "fmt"
)

// ToBase64 function    takes a hex string and returns a string representing the base64 value of the hex string.
func ToBase64(hexStr string) string {
  hexToBytes, err := hx.DecodeString(hexStr)
    if (err != nil) {
        fmt.Println(err)
    }

  b64Str := base64.StdEncoding.EncodeToString(hexToBytes)

  return b64Str
}

// ToByteArray function    takes a slice of runes(chars) representing a hex string and returns a slice of bytes representing the byte/ascii value of the string.
func ToByteArray(hexStrSlice []rune) []byte {
  bytesSlice := make([]byte, 0, len(hexStrSlice)/2)

  for i := 0; i < len(hexStrSlice); i += 2 {
    char1 := hexStrSlice[i]
    char2 := hexStrSlice[i+1]
    byteVal := RuneToByte(char1)*16 + RuneToByte(char2)
    bytesSlice = append(bytesSlice, byteVal)
  }

  return bytesSlice
}

// RuneToByte function    takes a rune and returns a byte representing the hex value of the rune.
func RuneToByte(char rune) byte {
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

// XorStrings function    takes two hex strings and returns a hex string representing the XOR of the two strings.
func XorStrings(hexStr1 string, hexStr2 string) string {
    if len(hexStr1) != len(hexStr2) {
        fmt.Println("error: XorStrings")
    }

    hexStr1Bytes, err := hx.DecodeString(hexStr1)
    if (err != nil) {
        fmt.Println(err)
    }
    hexStr2Bytes, err := hx.DecodeString(hexStr2)
    if (err != nil) {
        fmt.Println(err)
    }

    xorBytes, err := bytes.XorBytes(hexStr1Bytes, hexStr2Bytes)
    if (err != nil) {
        fmt.Println(err)
    }

    xorHexStr := hx.EncodeToString(xorBytes)

    return xorHexStr
}
