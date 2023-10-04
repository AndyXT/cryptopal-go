package bytes

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
