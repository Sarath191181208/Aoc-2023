package main

import (
	"bufio"
	"fmt"
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
	times := readSpaceDelimeterdInts(strings.Split(fileContents[0], ":")[1])
	distances := readSpaceDelimeterdInts(strings.Split(fileContents[1], ":")[1])
	mul := 1
	for i, time := range times {
		mul *= calcNumRecordBreaks(time, distances[i])
	}
	fmt.Print(mul)
}

func calcNumRecordBreaks(time int, distance int) int {
	sum := 0
	for i := 0; i < time; i++ {
		num := i * (time - i)
		if num > distance {
			sum += 1
		}
	}
	return sum
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
