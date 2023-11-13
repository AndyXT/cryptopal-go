package main

import (
	"cryptopal-go/pkg/encoding"
	"cryptopal-go/pkg/hex"
	"cryptopal-go/pkg/util"

	// "encoding/base64"
	hx "encoding/hex"
	"errors"
	"fmt"
)

func main() {
	err := challenge1_1()
	if err != nil {
		fmt.Println("Error: Challenge 1.1")
		fmt.Println(err)
	}
	err = challenge1_2()
	if err != nil {
		fmt.Println("Error: Challenge 1.2")
		fmt.Println(err)
	}
	err = challenge1_3()
	if err != nil {
		fmt.Println("Error: Challenge 1.3")
		fmt.Println(err)
	}
	err = challenge1_4()
	if err != nil {
		fmt.Println("Error: Challenge 1.4")
		fmt.Println(err)
	}
	err = challenge1_5()
	if err != nil {
		fmt.Println("Error: Challenge 1.5")
		fmt.Println(err)
	}
	err = challenge1_6()
	if err != nil {
		fmt.Println("Error: Challenge 1.6")
		fmt.Println(err)
	}
}

func challenge1_1() error {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	b64Str := hex.ToBase64(hexStr)
	fmt.Printf("%s\n", b64Str)
	if b64 != b64Str {
		return errors.New("Error: challenge 1.1")
	}

	return nil
}

func challenge1_2() error {
	hexStr1 := "1c0111001f010100061a024b53535009181c"
	hexStr2 := "686974207468652062756c6c277320657965"
	ansHexStr := "746865206b696420646f6e277420706c6179"

	xorStr := hex.XorStrings(hexStr1, hexStr2)

	if xorStr != ansHexStr {
		fmt.Println("Error: challenge2")
		return errors.New("Error Challenge 1.2")
	}
	fmt.Printf("%s\n", xorStr)
	return nil
}

func challenge1_3() error {
	hexStr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	decryptedBytes, _ := encoding.DecodeXorString(hexStr)

	if string(decryptedBytes) != "Cooking MC's like a pound of bacon" {
		fmt.Println("Error: challenge3")
		return errors.New("Error Challenge 1.3")
	}
	fmt.Printf("%s\n", decryptedBytes)
	return nil
}

func challenge1_4() error {
	path := "4.txt"
	encodedHexStrings, err := util.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var topScore float64
	decodedBytes := make([]byte, 0)

	for i, line := range encodedHexStrings {
		decodedLineBytes, score := encoding.DecodeXorString(line)

		if i == 0 {
			topScore = score
		} else if score > topScore {
			topScore = score
			decodedBytes = decodedLineBytes
		}
	}

	fmt.Println(string(decodedBytes))
	return nil
}

func challenge1_5() error {
	plainString := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	repeatingKey := encoding.RepeatingXorKey(&key, &plainString)

	encodedBytes := encoding.RepeatingXor(repeatingKey, plainString)

	encodedHexStr := hx.EncodeToString(encodedBytes)

	if encodedHexStr != "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f" {
		fmt.Println("Error: challenge5")
		return errors.New("Challenge 1.5 incorrect")
	}

	fmt.Println("Challenge1.5 Passed")
	return nil
}

func challenge1_6() error {
	//testing my hamming distance function
	str1 := "this is a test"
	str2 := "wokka wokka!!!"
	hamDist := encoding.HammingDistanceByteSlice([]byte(str1), []byte(str2))
	if hamDist != 37 {
		fmt.Println("Error: challenge6")
		return errors.New("Challenge 1.6 hamming distance incorrect")
	}

	path := "6.txt"
	b64Str, err := util.ReadFileAll(path)
	if err != nil {
		fmt.Println(err)
	}

	// variable to store best key size
	var keySize int = 0
	var bestHammingDist float64 = 0
	for size := 2; size < 40; size++ {
		// break into blocks of size
		// get hamming distance of each block
		hamDist1 := encoding.HammingDistanceByteSlice([]byte(b64Str[:size]), []byte(b64Str[size:2*size]))
		hamDist2 := encoding.HammingDistanceByteSlice([]byte(b64Str[2*size:3*size]), []byte(b64Str[3*size:4*size]))

		// average hamming distance
		// normalize by dividing by size
		hammingDistNorm := float64(hamDist1 + hamDist2) / float64(2 * size)
		// smallest normalized hamming distance is the key size
		if size == 2 {
			bestHammingDist = hammingDistNorm
			keySize = size
		} else if bestHammingDist > hammingDistNorm {
			bestHammingDist = hammingDistNorm
			keySize = size
		}
	}
	fmt.Println(keySize)

	blockSize := keySize
	blocks := make([]string, 0)
	for block := 0; block < len(b64Str) / int(keySize); block++ {
		// get block
		blocks = append(blocks, b64Str[block*int(blockSize):(block+1)*int(blockSize)])
		// get key
		// xor block with key
		// get score
		// if score is better than best score, replace best score
	}

	blocksT := util.Transpose(blocks)

	fmt.Println(blocks)
	fmt.Println(blocksT)
	return nil
}
