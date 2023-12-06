package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filePath := os.Args[1]
	fileContents := readFileFromPath(filePath)
	solvePart2(fileContents)
}

func solvePart1(fileContents []string) {
	sum := 0
	for lineNumber, singleLine := range fileContents {
		startIdxPointer := 0
		for {
			nextIndex, isNumber := iterateTillNextIndex(singleLine, startIdxPointer)
			if isNumber {
				isSymbolAround := checkIsSymbolAround(
					fileContents,
					startIdxPointer,
					nextIndex-1,
					lineNumber,
				)
				numStr := singleLine[startIdxPointer:nextIndex]
				if isSymbolAround {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						panic("Can't convert the value into int num: " + numStr)
					}
					sum += num
				}
			}
			startIdxPointer = nextIndex
			if nextIndex > len(singleLine) {
				break
			}
		}
	}
	fmt.Println("Sum: ", sum)
}

func solvePart2(fileContents []string) {
	sum := 0
	for lineNumber, singleLine := range fileContents {
		fmt.Print(lineNumber, "| ", singleLine, " |: ")
		for i, char := range singleLine {
			if char == '*' {
				num1, num2 := getNumsAroundAsterisk(fileContents, i, lineNumber)
				if num1 == -1 || num2 == -1 {
					continue
				}
				fmt.Print(num1, ", ", num2, "; ")
				sum += num1 * num2
			}
		}
		fmt.Println()
	}
	fmt.Println("Sum: ", sum)
}

func iterateTillNextIndex(engineSchematicLine string, currentIndex int) (int, bool) {
	if currentIndex >= len(engineSchematicLine) {
		return len(engineSchematicLine) + 1, false
	}
	isNumber := isDigit(engineSchematicLine[currentIndex])
	for i, char := range engineSchematicLine[currentIndex:] {
		if isDigit(byte(char)) != isNumber {
			return currentIndex + i, isNumber
		}
	}
	return len(engineSchematicLine), isNumber
}

func checkIsSymbolAround(
	engineSchematicLine []string,
	blockStartIdx int,
	blockEndIdx int,
	lineNumber int,
) bool {
	singleEngineSchematicLine := engineSchematicLine[lineNumber]
	// checking left
	if blockStartIdx-1 > 0 {
		idx := blockStartIdx - 1
		chr := singleEngineSchematicLine[idx]
		if chr != '.' && !isDigit(chr) {
			return true
		}
	}

	// checking right
	if blockEndIdx+1 < len(singleEngineSchematicLine) {
		idx := blockEndIdx + 1
		chr := singleEngineSchematicLine[idx]
		if chr != '.' && !isDigit(chr) {
			return true
		}
	}

	// checking bottom
	isSymbolDown := isSymbolInRange(
		lineNumber+1,
		max(blockStartIdx-1, 0),
		min(blockEndIdx+1, len(singleEngineSchematicLine)),
		engineSchematicLine,
	)

	if isSymbolDown {
		return true
	}

	// checking up
	isSymbolUp := isSymbolInRange(
		lineNumber-1,
		max(blockStartIdx-1, 0),
		min(blockEndIdx+1, len(singleEngineSchematicLine)-1),
		engineSchematicLine,
	)

	return isSymbolUp
}

func getNumsAroundAsterisk(
	engineSchematic []string,
	asteriskIndex int,
	lineNumber int,
) (int, int) {
	line := engineSchematic[lineNumber]
	var resArr []int
	downNumIndex := findNumberIndex(
		lineNumber+1,
		max(asteriskIndex-1, 0),
		min(asteriskIndex+1, len(line)),
		engineSchematic,
	)
	if downNumIndex != -1 {
		num1 := getNumFromStartIdx(engineSchematic[lineNumber+1], downNumIndex)
		resArr = append(resArr, num1)
	}
	newDownIndex := findNumberIndex(
		lineNumber+1,
		min(asteriskIndex+1, len(line)),
		min(asteriskIndex+1, len(line)),
		engineSchematic,
	)
	if newDownIndex != -1 && newDownIndex != downNumIndex {
		num1 := getNumFromStartIdx(engineSchematic[lineNumber+1], newDownIndex)
		resArr = append(resArr, num1)
	}
	upNumIndex := findNumberIndex(
		lineNumber-1,
		max(asteriskIndex-1, 0),
		min(asteriskIndex+1, len(line)),
		engineSchematic,
	)
	if upNumIndex != -1 {
		num1 := getNumFromStartIdx(engineSchematic[lineNumber-1], upNumIndex)
		resArr = append(resArr, num1)
	}
	newUpIndex := findNumberIndex(
		lineNumber-1,
		min(asteriskIndex+1, len(line)),
		min(asteriskIndex+1, len(line)),
		engineSchematic,
	)
	if newUpIndex != -1 && newUpIndex != upNumIndex {
		num1 := getNumFromStartIdx(engineSchematic[lineNumber-1], newUpIndex)
		resArr = append(resArr, num1)
	}

	if asteriskIndex-1 > 0 {
		idx := asteriskIndex - 1
		chr := line[idx]
		if chr != '.' && isDigit(chr) {
			startingIndexOfLeftNum := findNumberIndex(
				lineNumber,
				asteriskIndex-1,
				asteriskIndex,
				engineSchematic,
			)
			if startingIndexOfLeftNum != -1 {
				num1 := getNumFromStartIdx(line, startingIndexOfLeftNum)
				resArr = append(resArr, num1)
			}

		}
	}

	if asteriskIndex+1 < len(line) {
		idx := asteriskIndex + 1
		chr := line[idx]
		if chr != '.' && isDigit(chr) {
			num2 := getNumFromStartIdx(line, idx)
			resArr = append(resArr, num2)
		}
	}

	if len(resArr) < 2 {
		return -1, -1
	}

	return resArr[0], resArr[1]
}

func getNumFromStartIdx(engineSchematicLine string, startIndex int) int {
	nextIndex, isNumber := iterateTillNextIndex(engineSchematicLine, startIndex)
	if !isNumber {
		panic("Invalid arguments to the function")
	}
	numStr := engineSchematicLine[startIndex:nextIndex]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic("Can't convert numStr to number numStr: " + numStr)
	}
	return num
}

func findNumberIndex(
	lineNumber int,
	startIndex int,
	endIndex int,
	engineSchematic []string,
) int {
	if lineNumber < 0 || lineNumber >= len(engineSchematic) {
		return -1
	}
	line := engineSchematic[lineNumber]
	if startIndex < len(line) {
		fmt.Print("(", string(line[startIndex]), ")")
	}
	for i := startIndex; i <= endIndex; i++ {
		isCharDigit := isDigit(byte(line[i]))
		if isCharDigit {
			idx := i
			// find the starting point of the number
			for idx >= 0 && isDigit(byte(line[idx])) {
				idx--
			}
			return idx + 1
		}
	}
	return -1
}

func isSymbolInRange(lineNumber int, startIndex int, endIndex int, engineSchematic []string) bool {
	if lineNumber < 0 || lineNumber >= len(engineSchematic) {
		return false
	}
	line := engineSchematic[lineNumber]

	for i, char := range line {
		isInBounds := i >= startIndex && i <= endIndex
		isCharPeriod := char == '.'
		isCharDigit := isDigit(byte(char))
		if isInBounds && !isCharPeriod && !isCharDigit {
			return true
		}
	}
	return false
}

func readFileFromPath(path string) []string {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println("File not found on path")
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

func isDigit(b byte) bool {
	if b < '0' || b > '9' {
		return false
	}
	return true
}
