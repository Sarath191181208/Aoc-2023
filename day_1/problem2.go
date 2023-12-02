package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		num1, num2 := findFirstAndLastPosDigit(eachline)
		fmt.Println(num1*10 + num2)
		sum += num1*10 + num2
	}
	fmt.Println(sum)
}

func findFirstAndLastPosDigit(line string) (int, int) {
	firstValue := 0
	secondValue := 0
	numMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	subString := ""
	for i, char := range line {
		if val := getValueIfSubstringInMap(subString, numMap); val != -1 {
			firstValue = val
			break
		}
		if unicode.IsDigit(char) {
			firstValue = int((line[i]) - 48)
			break
		}
		subString += string(char)
	}

	subString = ""
	for i := range line {
		var char rune = rune(line[len(line)-i-1])
		if val := getValueIfSubstringInMap(Reverse(subString), numMap); val != -1 {
			secondValue = val
			break
		}

		if unicode.IsDigit(char) {
			secondValue = int(line[len(line)-i-1]) - 48
			break
		}
		subString += string(char)
	}
	return firstValue, secondValue
}

func getValueIfSubstringInMap(line string, numMap map[string]int) int {
	for key, value := range numMap {
		if strings.Contains(line, key) {
			return value
		}
	}
	return -1
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
