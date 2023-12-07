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
	solveProblem2(fileContents)
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

type SeedHolder struct {
	startNum int
	endNum   int
}

func solveProblem2(fileContents []string) {
	groupedSeedsArr := groupSeedsArray(extractSeedsArray(fileContents))
	groupedMapBlocks := splitStringListAsBlocks(fileContents[2:])
	for _, block := range groupedMapBlocks {
		domainConversionData := parseSingleBlock(block[1:])
		tempTransformedRanges := applyRangedDomainTransform(&groupedSeedsArr, domainConversionData)
		groupedSeedsArr = tempTransformedRanges
		// fmt.Println(groupedSeedsArr)
	}
	fmt.Println(groupedSeedsArr.Min().startNum)
}

func groupSeedsArray(arr []int) SeedHolderStack {
	var res SeedHolderStack
	for i := 0; i < len(arr); i += 2 {
		res.Push(SeedHolder{
			startNum: arr[i],
			endNum:   arr[i] + arr[i+1],
		})
	}
	return res
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

func applyRangedDomainTransform(
	stack *SeedHolderStack,
	domainConversionData []ConversionData,
) SeedHolderStack {
	var temp SeedHolderStack
	for !stack.IsEmpty() {
		ele, hasElement := stack.Pop()
		if !hasElement {
			break
		}
		flagIsMatched := false
		for _, data := range domainConversionData {
			start := data.previousDomainStartNumber
			end := data.previousDomainStartNumber + data.stepValue - 1
			overlap_start := max(start, ele.startNum)
			overlap_end := min(ele.endNum, end+1)
			if overlap_start < overlap_end {
				temp.Push(SeedHolder{
					startNum: (overlap_start - data.previousDomainStartNumber) + data.resultingDomainStartNumber,
					endNum:   (overlap_end - data.previousDomainStartNumber) + data.resultingDomainStartNumber,
				})
				if overlap_start > ele.startNum {
					stack.Push(SeedHolder{
						startNum: ele.startNum,
						endNum:   overlap_start,
					})
				}
				if overlap_end < ele.endNum {
					stack.Push(SeedHolder{
						startNum: overlap_end,
						endNum:   ele.endNum,
					})
				}
				flagIsMatched = true
				break
			}
		}
		if !flagIsMatched {
			temp.Push(ele)
		}
	}
	return temp
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

type SeedHolderStack []SeedHolder

func (s *SeedHolderStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *SeedHolderStack) Push(val SeedHolder) {
	*s = append(*s, val) // Simply append the new value to the end of the stack
}

func (s *SeedHolderStack) Pop() (SeedHolder, bool) {
	if s.IsEmpty() {
		return SeedHolder{}, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *SeedHolderStack) Min() SeedHolder {
	smallNum := (*s)[0]
	for _, val := range *s {
		if smallNum.startNum > val.startNum {
			smallNum = val
		}
	}
	return smallNum
}
