package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	RED_CUBES_LIMIT   = 12
	GREEN_CUBES_LIMIT = 13
	BLUE_CUBES_LIMIT  = 14
)

type ColorCounts struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	filePath := os.Args[1]
	fileLines := readFile(filePath)

	sum := 0
	for _, line := range fileLines {
		colorCounts := readFileStr(line)
		score := calculateScore(colorCounts)
		sum += score
	}
	fmt.Println(sum)
}

func calculateScore(colorCounts []ColorCounts) int {
	maxColorCounts := ColorCounts{0, 0, 0}
	for _, colorCount := range colorCounts {
		maxColorCounts.Red = max(maxColorCounts.Red, colorCount.Red)
		maxColorCounts.Green = max(maxColorCounts.Green, colorCount.Green)
		maxColorCounts.Blue = max(maxColorCounts.Blue, colorCount.Blue)
	}
	return maxColorCounts.Red * maxColorCounts.Green * maxColorCounts.Blue
}

func checkIfValid(colorCounts ColorCounts) bool {
	if colorCounts.Red > RED_CUBES_LIMIT {
		return false
	} else if colorCounts.Green > GREEN_CUBES_LIMIT {
		return false
	} else if colorCounts.Blue > BLUE_CUBES_LIMIT {
		return false
	}
	return true
}

func readFile(filePath string) []string {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file")
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

func readFileStr(line string) []ColorCounts {
	// Game 1: 10 red, 7 green, 3 blue; 5 blue, 3 red, 10 green; 4 blue, 14 green, 7 red; 1 red, 11 green; 6 blue, 17 green, 15 red; 18 green, 7 red, 5 blue
	splitLine := strings.Split(line, ":")
	gameStates := strings.Split(splitLine[1], ";") // ["10 red, 7 green, 3 blue", ...]
	colorCountsVec := make([]ColorCounts, len(gameStates))
	for _, gameState := range gameStates {
		colorCounts := ColorCounts{0, 0, 0}
		singleRound := strings.Split(gameState, ",") // [ ... ,"10 red"]
		for _, color := range singleRound {
			countAndColor := strings.Split(strings.TrimSpace(color), " ")
			countStr, color := countAndColor[0], countAndColor[1]
			count, err := strconv.Atoi(countStr)
			if err != nil {
				fmt.Println("Error converting string to int")
				panic(err)
			}
			switch color {
			case "red":
				colorCounts.Red += count
			case "green":
				colorCounts.Green += count
			case "blue":
				colorCounts.Blue += count
			}
		}
		colorCountsVec = append(colorCountsVec, colorCounts)
	}
	return colorCountsVec
}
