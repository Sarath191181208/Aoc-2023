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
