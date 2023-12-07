package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := os.Args[1]
	fileContents := readFileFromPath(filePath)
	solveProblem1(fileContents)
}

func solveProblem1(fileContents []string) {
}

func readSpaceDelimeterdInts(s string) []int {
	strArr := strings.Split(strings.TrimSpace(s), " ")
	var numsArr []int
	for _, val := range strArr {
		if val == "" {
			continue
		}
		num, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			panic("Can't convert" + val + " into num, Failed!")
		}
		numsArr = append(numsArr, num)
	}
	return numsArr
}

func readFileFromPath(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic("Error opening the file: " + filePath + ", File Not Found!")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines
}
