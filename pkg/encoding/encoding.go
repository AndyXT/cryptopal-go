package encoding

import (
	"cryptopal-go/pkg/hex"
)

func DecodeXorString(hexStr []rune) ([]byte, float64) {
	encryptedBytes := hex.ToByteArray([]rune(hexStr))

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
