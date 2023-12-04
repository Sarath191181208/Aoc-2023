package main

import (
	"bufio"
	"fmt"
	"os"
)

type Color int

const (
	Red Color = iota
	Green
	Blue
)

func main() {
	filePath := os.Args[1]
	fileLines := readFile(filePath)
}

func readFile(filePath string) []string {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file")
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines
}
