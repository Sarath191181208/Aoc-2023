package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filePath := os.Args[1]
	fileContents := readFileFromPath(filePath)
	solveProblem2(fileContents)
}

func solveProblem1(scratchCards []string) {
	sum := 0
	for _, cardString := range scratchCards {
		splitLine := splitAndStrip(cardString)
		winningNums := toIntArr(splitLine[1])
		ourScratchCards := toIntArr(splitLine[2])
		n := intersectionCountOfArray(winningNums, ourScratchCards)
		if n != 0 {
			sum += pow(2, n-1)
		}
	}
	fmt.Println("Sum: ", sum)
}

func solveProblem2(scratchCards []string) {
	sum := 0
	memoizationMap := make(map[int]int)
	for i := range scratchCards {
		sum += playSingleIterationDFS(scratchCards, i, memoizationMap)
		fmt.Println(i, " |")
	}
	fmt.Println("Sum: ", sum)
}

func playSingleIterationDFS(scratchCards []string, index int, memoizationMap map[int]int) int {
	if index >= len(scratchCards) {
		return 0
	}
	mapVal, isIndexInMap := memoizationMap[index]
	if isIndexInMap {
		return mapVal
	}

	splitLine := splitAndStrip(scratchCards[index])
	winningNums := toIntArr(splitLine[1])
	ourScratchCards := toIntArr(splitLine[2])
	n := intersectionCountOfArray(winningNums, ourScratchCards)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += playSingleIterationDFS(scratchCards, index+i, memoizationMap)
	}
	sum = sum + 1 // for the iteration of this element
	memoizationMap[index] = sum
	return sum
}

func pow(a, b int) int {
	mul := 1
	for i := 0; i < b; i++ {
		mul *= a
	}
	return mul
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

func intersectionCountOfArray(arr1, arr2 []int) int {
	count := 0
	for _, val := range arr1 {
		for _, val2 := range arr2 {
			if val == val2 {
				count++
			}
		}
	}
	return count
}

func toIntArr(s string) []int {
	var resArr []int
	split := regexSplit(s, `\s+`)
	for _, valStr := range split {
		val, err := strconv.Atoi(valStr)
		if err != nil {
			panic("The given string: " + valStr + " Can't be parsed")
		}
		resArr = append(resArr, val)
	}
	return resArr
}

func splitAndStrip(s string) []string {
	// Replace all spaces with an empty string using regexp.ReplaceAllString.
	split := regexSplit(s, `(:|\|)`)
	set := []string{}
	for i := range split {
		set = append(set, strings.TrimSpace(split[i]))
	}
	return set
}

func regexSplit(s string, regex string) []string {
	re := regexp.MustCompile(regex)
	numberOfSplits := -1
	split := re.Split(s, numberOfSplits)
	set := []string{}
	set = append(set, split...)
	return set
}
