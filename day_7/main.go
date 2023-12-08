package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CardHand int

const (
	HighCard CardHand = iota + 1
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (c CardHand) Value() int {
	return int(c)
}

func main() {
	filePath := os.Args[1]
	fileContents := readFileFromPath(filePath)
	solveProblem1(fileContents)
}

func solveProblem1(fileContents []string) {
	var roundsList []Round
	for _, line := range fileContents {
		strArr := strings.Split(line, " ")
		cardHandString := strArr[0]
		bidAmount, err := strconv.Atoi(strArr[1])
		if err != nil {
			panic("Can't convert " + strArr[1] + " ,to int")
		}
		// sum += bidAmount * matchMapToCardType(countCharactersInStr(cardHandString)).Value()
		roundsList = append(roundsList, Round{
			handString: cardHandString,
			handType:   matchMapToCardType(countCharactersInStr(cardHandString)),
			bidAmount:  bidAmount,
		})
	}

	sort.Slice(roundsList, func(i, j int) bool {
		a := roundsList[i]
		b := roundsList[j]
		if a.handType.Value() == b.handType.Value() {
			return compareHandStrings(a.handString, b.handString)
		}
		return a.handType.Value() < b.handType.Value()
	})

	sum := 0
	for i, round := range roundsList {
		sum += (i + 1) * round.bidAmount
	}
	fmt.Print("Sum: ", sum)
}

type Round struct {
	handString string
	handType   CardHand
	bidAmount  int
}

func matchMapToCardType(cardCharacterCountMap map[rune]int) CardHand {
	sortedKeysValuesList := sortMap(cardCharacterCountMap)
	maxElement := sortedKeysValuesList[0]
	if maxElement.value == 5 {
		return FiveOfAKind
	}
	nextElement := sortedKeysValuesList[1]
	if maxElement.value == 4 {
		return FourOfAKind
	} else if maxElement.value == 3 && nextElement.value == 2 {
		return FullHouse
	} else if maxElement.value == 3 {
		return ThreeOfKind
	} else if maxElement.value == 2 && nextElement.value == 2 {
		return TwoPair
	} else if maxElement.value == 2 {
		return OnePair
	}
	return HighCard
}

func countCharactersInStr(cardHandString string) map[rune]int {
	countingDict := make(map[rune]int)
	for _, char := range cardHandString {
		currentCount, isCharInDict := countingDict[char]
		if isCharInDict {
			countingDict[char] = currentCount + 1
		} else {
			countingDict[char] = 1
		}
	}
	return countingDict
}

func compareHandStrings(a, b string) bool {
	for i, char := range a {
		val1 := getSingleRuneValue(char)
		val2 := getSingleRuneValue(rune(b[i]))
		if val1 != val2 {
			return val1 < val2
		}
	}
	return false
}

func getSingleRuneValue(r rune) int {
	if r == 'A' {
		return 14
	} else if r == 'K' {
		return 13
	} else if r == 'Q' {
		return 12
	} else if r == 'J' {
		return 11
	} else if r == 'T' {
		return 10
	}
	return int(r - '0')
}

type KeyValue struct {
	key   rune
	value int
}

func sortMap(inputMap map[rune]int) []KeyValue {
	var res []KeyValue
	for k, val := range inputMap {
		res = append(res, KeyValue{k, val})
	}
	sort.Slice(res, func(i, j int) bool {
		a := res[i]
		b := res[j]
		return a.value > b.value
	})
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
