package main

import (
	"cryptopal-go/pkg/hex"
	"fmt"
)

func main() {
	fmt.Println("vim-go")

	challenge1_1()
}

func challenge1_1() {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	b64Str := hex.ToBase64(hexStr)

	fmt.Printf("%s", b64Str)
	if b64 != b64Str {
		fmt.Println("Error: challenge1")
	}
}
