package bytes

// ToBase64Str function    takes a slice of bytes and returns a string representing the base64 value of the bytes.
func ToBase64Str(bytes []byte) string {
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

// ToHexRunes function    takes a byte and returns a slice of 2 runes representing the hex value of the byte.
func ToHexRunes(hexStrByte byte) []rune {
  hexAlphabet := [16]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',  'a', 'b', 'c', 'd', 'e', 'f'}
  hexStr := make([]rune, 0, 2)

  zeroPlace := hexStrByte % 16
  sixteenPlace := (hexStrByte - zeroPlace) / 16

  hexStr = append(hexStr, hexAlphabet[sixteenPlace])
  hexStr = append(hexStr, hexAlphabet[zeroPlace])

  return hexStr
}

// XorBytes function    takes 2 slices of bytes and returns a slice of bytes representing the xor of the 2 slices.
func XorBytes(bytes1 []byte, bytes2 []byte) []byte {
  xorBytes := make([]byte, 0, len(bytes1))

  for i := 0; i < len(bytes1); i++ {
    xorBytes = append(xorBytes, bytes1[i]^bytes2[i])
  }

  return xorBytes
}

// SliceToHexStr function    takes a slice of bytes and returns a string representing the hex value of the bytes.
func SliceToHexStr(bytes []byte) string {
	hexStr := make([]rune, 0, len(bytes)*2)

	for _, b := range bytes {
		hexStr = append(hexStr, ToHexRunes(b)...)
	}

	return string(hexStr)
}
