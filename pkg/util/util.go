package util

import (
	"os"
	"bufio"
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
// func ReadFile(path string) []string {
// 	readFile, err := os.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanLines)
//
// 	var fileTextLines []string
// 	for fileScanner.Scan() {
// 		fileTextLines = append(fileTextLines, fileScanner.Text())
// 	}
// 	readFile.Close()
//
// 	return fileTextLines
// }
