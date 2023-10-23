package main

import (
	"cryptopal-go/pkg/bytes"
	"cryptopal-go/pkg/encoding"
	"cryptopal-go/pkg/hex"
	"cryptopal-go/pkg/util"
	"fmt"
)

func main() {
	challenge1_1()
	challenge1_2()
	challenge1_3()
	challenge1_4()
	challenge1_5()
}

func challenge1_1() {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	b64Str := hex.ToBase64(hexStr)

	fmt.Printf("%s\n", b64Str)
	if b64 != b64Str {
		fmt.Println("Error: challenge1")
	}
}

func challenge1_2() {
	hexStr1 := "1c0111001f010100061a024b53535009181c"
	hexStr2 := "686974207468652062756c6c277320657965"
	xorHexStr := "746865206b696420646f6e277420706c6179"

	xorStr := hex.XorStrings([]rune(hexStr1), []rune(hexStr2))

	fmt.Printf("%s\n", xorStr)
	if xorStr != xorHexStr {
		fmt.Println("Error: challenge2")
	}
}

func challenge1_3() {
	hexStr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	decryptedBytes, _ := encoding.DecodeXorString([]rune(hexStr))

	fmt.Printf("%s\n", decryptedBytes)
	if string(decryptedBytes) != "Cooking MC's like a pound of bacon" {
		fmt.Println("Error: challenge3")
	}
}

func challenge1_4() {
	path := "4.txt"
	encodedHexStrings := util.ReadFile(path)

	var topScore float64
	decodedBytes := make([]byte, 0)

	for i, line := range encodedHexStrings {
		decodedLineBytes, score := encoding.DecodeXorString([]rune(line))

		if i == 0 {
			topScore = score
		} else if score > topScore {
			topScore = score
			decodedBytes = decodedLineBytes
		}
	}

	fmt.Println(string(decodedBytes))
}

func challenge1_5() {
	plainString := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	repeatingKey := encoding.RepeatingXorKey(&key, &plainString)

	encodedBytes := encoding.RepeatingXor(&repeatingKey, &plainString)

	fmt.Println(len(encodedBytes))

	encodedHexStr := bytes.SliceToHexStr(encodedBytes)

	fmt.Println(len(encodedHexStr))

	if encodedHexStr != "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f" {
		fmt.Println("Error: challenge5")
	}

	fmt.Printf("%d\n", len(plainString))
	fmt.Printf("%d\n", len("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"))
	fmt.Printf("%s\n", "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	fmt.Printf("%s\n", encodedHexStr)
}
