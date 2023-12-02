package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	filepath := os.Args[1]
	readFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	sum := 0
	for _, eachline := range fileLines {
		firstPos, lastPos := getFirstAndLastDigit(eachline)
		num1 := int(eachline[firstPos]) - 48
		num2 := int(eachline[len(eachline)-lastPos-1]) - 48
		sum += num1*10 + num2
	}
	fmt.Println(sum)
}

func getFirstAndLastDigit(line string) (int, int) {
	firstPos := 0
	lastPos := 0
	for i, char := range line {
		// check if the bytes are is digit
		if unicode.IsDigit(char) {
			firstPos = i
			break
		}
	}
	for i := range line {
		var char rune = rune(line[len(line)-i-1])
		if unicode.IsDigit(char) {
			lastPos = i
			break
		}
	}
	return firstPos, lastPos
}
