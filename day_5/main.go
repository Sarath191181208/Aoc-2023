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
	seedsArr := extractSeedsArray(fileContents)
	groupedMapBlocks := splitStringListAsBlocks(fileContents[2:])
	for _, block := range groupedMapBlocks {
		// name := block[0]
		domainConversionData := parseSingleBlock(block[1:])
		for i, num := range seedsArr {
			seedsArr[i] = applyDomainTransform(num, domainConversionData)
		}
	}
	fmt.Println(findMinInArr(seedsArr))
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

func splitStringListAsBlocks(fileContents []string) [][]string {
	var res [][]string
	var stringContentBuffer []string
	for _, line := range fileContents {
		if line == "" && len(stringContentBuffer) > 0 {
			res = append(res, stringContentBuffer)
			stringContentBuffer = make([]string, 0)
		} else {
			stringContentBuffer = append(stringContentBuffer, line)
		}
	}
	if len(stringContentBuffer) > 0 {
		res = append(res, stringContentBuffer)
	}
	return res
}

type ConversionData struct {
	resultingDomainStartNumber int
	previousDomainStartNumber  int
	stepValue                  int
}

func parseSingleBlock(singleConversionBlockString []string) []ConversionData {
	var res []ConversionData
	for _, line := range singleConversionBlockString {
		singleLineNums := readSpaceDelimeterdInts(line)
		res = append(res, ConversionData{
			resultingDomainStartNumber: singleLineNums[0],
			previousDomainStartNumber:  singleLineNums[1],
			stepValue:                  singleLineNums[2],
		})
	}
	return res
}

func applyDomainTransform(num int, converstionData []ConversionData) int {
	for _, data := range converstionData {
		start := data.previousDomainStartNumber
		end := data.previousDomainStartNumber + data.stepValue - 1
		if num >= start && num <= end {
			diff := num - start
			return data.resultingDomainStartNumber + diff
		}
	}
	return num
}

func findMinInArr(arr []int) int {
	smallNum := arr[0]
	for _, num := range arr[1:] {
		if num < smallNum {
			smallNum = num
		}
	}
	return smallNum
}

func extractSeedsArray(fileContents []string) []int {
	startingLine := fileContents[0]
	res := strings.Split(startingLine, ":")[1]
	res = strings.TrimSpace(res)
	return readSpaceDelimeterdInts(res)
}

func readSpaceDelimeterdInts(s string) []int {
	strArr := strings.Split(s, " ")
	var numsArr []int
	for _, val := range strArr {
		num, err := strconv.Atoi(strings.TrimSpace(val))
		if err != nil {
			panic("Can't convert" + val + " into num, Failed!")
		}
		numsArr = append(numsArr, num)
	}
	return numsArr
}
