package encoding

import (
	hx "encoding/hex"
	"fmt"
)

func DecodeXorString(hexStr string) ([]byte, float64) {
	// encryptedBytes := hex.ToByteArray([]rune(hexStr))
	encryptedBytes, err := hx.DecodeString(hexStr)
    if (err != nil) {
        fmt.Println(err)
    }

	bestScore, key := getScore(encryptedBytes)

	decryptedBytes := singleByteXorCipher(key, encryptedBytes)

	return decryptedBytes, bestScore
}

func singleByteXorCipher(key byte, encryptedBytes []byte) []byte {
	decryptedBytes := make([]byte, len(encryptedBytes))
	for i := 0; i < len(encryptedBytes); i++ {
		decryptedBytes[i] = encryptedBytes[i] ^ key
	}

	return decryptedBytes
}

func getScore(encryptedBytes []byte) (float64, byte) {

	occurenceInEnglish := [26]float64{
		8.2389258, 1.5051398, 2.8065007, 4.2904556, 12.813865, 2.2476217, 2.0327458, 6.1476691,
		6.1476691, 0.1543474, 0.7787989, 4.0604477, 2.4271893, 6.8084376, 7.5731132, 1.9459884,
		0.0958366, 6.0397268, 6.3827211, 9.1357551, 2.7822893, 0.9866131, 2.3807842, 0.1513210,
		1.9913847, 0.0746517,
	}

	var bestScore float64
	var bestKey byte
	for i := 0; i < 256; i++ {
		var score float64
		for _, b := range encryptedBytes {
			decodedByte := b ^ byte(i)

			if decodedByte >= 'a' && decodedByte <= 'z' {
				score += occurenceInEnglish[decodedByte-'a']
			} else if decodedByte >= 'A' && decodedByte <= 'Z' {
				score += occurenceInEnglish[decodedByte-'A']
			} else if decodedByte == ' ' {
				score += 15.0
			}
		}

		if i == 0 || score > bestScore {
			bestScore = score
			bestKey = byte(i)
		}
	}

	return bestScore, bestKey
}
func RepeatingXorKey(key *string, plainString *string) string {
	keyLength := len(*key)
	stringLength := len(*plainString)

	repeatingKey := make([]rune, 0, stringLength) // initialized with 0 length but with capacity of stringLength
	keyRunes := []rune(*key)

	for i := 0; i < stringLength; i++ {
		repeatingKey = append(repeatingKey, keyRunes[i%keyLength])
	}

	return string(repeatingKey)
}

// RepeatingXor function    returns a byte array representing the XOR of a key and a plain text string.
func RepeatingXor(key string, plainString string) []byte {
	plainBytes := []byte(plainString)
	keyBytes := []byte(key)

	encryptedBytes := make([]byte, len(plainBytes))

	for i := 0; i < len(plainBytes); i++ {
		encryptedBytes[i] = plainBytes[i] ^ keyBytes[i]
	}

	return encryptedBytes
}

func HammingDistanceByteSlice(bytes1 []byte, bytes2 []byte) (int, error) {
	if len(bytes1) != len(bytes2) {
		return 0, fmt.Errorf("byte slices are not of equal length")
	}
	
	var hamDist int
	for i := 0; i < len(bytes1); i++ {
		hamDist += hammingDistanceByte(bytes1[i], bytes2[i])
	}
	return hamDist, nil
}
// func HammingDistanceByteSlice(bytes1 []byte, bytes2 []byte) int {
// 	var hamDist int = 0
// 	for i := 0; i < len(bytes1); i++ {
// 		hamDist += hammingDistanceByte(bytes1[i], bytes2[i])
// 	}
// 	return hamDist
// }

func hammingDistanceByte(byte1 byte, byte2 byte) int {
	var distance int
	xor := byte1 ^ byte2
	for xor > 0 {
		distance += int(xor & 1)
		xor >>= 1
	}
	return distance
}
