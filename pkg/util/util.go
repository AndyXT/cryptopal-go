package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ReadFile function    Reads a file and returns a slice of strings
func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

//	func ReadFile(path string) []string {
//		readFile, err := os.Open(path)
//		if err != nil {
//			panic(err)
//		}
//		fileScanner := bufio.NewScanner(readFile)
//		fileScanner.Split(bufio.ScanLines)
//
//		var fileTextLines []string
//		for fileScanner.Scan() {
//			fileTextLines = append(fileTextLines, fileScanner.Text())
//		}
//		readFile.Close()
//
//		return fileTextLines
//	}
func ReadFileAll(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}

	return string(content), nil
}
func Transpose(input []string) []string {
	rowCount := len(input)
	columnCount := len(input[0])

	if columnCount == 0 {
		return []string{}
	}

	transposed := make([]string, columnCount)

	for i := 0; i < columnCount; i++ {
		t := ""

		for j := 0; j < rowCount; j++ {
			t += string(input[j][i])
		}

		transposed[i] = t
	}

	return transposed
}
func TransposeBytes(input [][]byte) [][]byte {
	// Check if input is empty
	if len(input) == 0 {
		return [][]byte{}
	}

	// Get the 'column' count - length of the first 'row'
	rowCount := len(input)
	columnCount := len(input[0])

	if columnCount == 0 {
		return [][]byte{}
	}

	// Make a 2D slice for the transposed result
	transposed := make([][]byte, columnCount)
	for i := range transposed {
		transposed[i] = make([]byte, rowCount) // Initialize each 'row'
	}

	// Transpose the input
	for i := 0; i < columnCount; i++ {
		for j := 0; j < rowCount; j++ {
			// Extra check if dimension is irregular matrix (length mismatch)
			if i >= len(input[j]) {
				fmt.Printf("Can't transpose irregular input, skipping value at [%d][%d]", j, i)
				continue
			}
			transposed[i][j] = input[j][i]
		}
	}

	return transposed
}
