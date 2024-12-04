// day4 of Advent Of Code 2024 : https://adventofcode.com/2024/day/4
// j4nbob
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadTextFromFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	var result [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineArray []string
		for _, char := range line {
			lineArray = append(lineArray, string(char))
		}
		result = append(result, lineArray)
	}
	if err := scanner.Err(); err != nil {
		os.Exit(2)
	}

	return result
}

func findNonBoundaryA(array [][]string) [][]int {
	var positions [][]int
	for i := 1; i < len(array)-1; i++ {
		for j := 1; j < len(array[i])-1; j++ {
			if array[i][j] == "A" {
				positions = append(positions, []int{i, j})
			}
		}
	}
	return positions
}

func isValidChar(char string) bool {
	return char == "M" || char == "S"
}

func isValidPossibleWords(possibleWords string) bool {
	words := strings.Split(possibleWords, ",")
	for _, word := range words {
		if word == "MAS" || word == "SAM" {
			continue
		} else {
			return false
		}
	}
	return true // we got here via all continues in the loop
}

// the counter of the X -MAS crosses based on the A positions that we detected previously
func findXCournersBasedOnA(array [][]string, positions [][]int) int {
	var daSum int = 0
	for _, position := range positions {
		i, j := position[0], position[1]
		topLeft := string(array[i-1][j-1])
		topRight := string(array[i-1][j+1])
		bottomLeft := string(array[i+1][j-1])
		bottomRight := string(array[i+1][j+1])
		if isValidChar(topLeft) && isValidChar(bottomRight) && isValidChar(topRight) && isValidChar(bottomLeft) {
			possibleWords := fmt.Sprintf("%sA%s,%sA%s,%sA%s,%sA%s", topLeft, bottomRight, bottomRight, topLeft, topRight, bottomLeft, bottomLeft, topRight)
			if isValidPossibleWords(possibleWords) {
				daSum++
			} // the false condition eliminates words like MAM or SAS or others
		}
	}
	return daSum
}

func findStringPatternsInArrays(array [][]string, pattern string, countVertsHoris bool) int {
	var count int = 0
	// directional mess - about to get UGLY. lets take the diagonals first  : \ and /
	// \
	for start := 0; start < len(array); start++ {
		var diagArray []string
		for row, col := start, 0; row < len(array) && col < len(array[0]); row, col = row+1, col+1 {
			diagArray = append(diagArray, array[row][col])
		}
		diagString := strings.Join(diagArray, "")
		count += strings.Count(diagString, pattern)
	}
	for start := 1; start < len(array[0]); start++ {
		var diagArray []string
		for row, col := 0, start; row < len(array) && col < len(array[0]); row, col = row+1, col+1 {
			diagArray = append(diagArray, array[row][col])
		}
		diagString := strings.Join(diagArray, "")
		count += strings.Count(diagString, pattern)
	}
	// /
	for start := 0; start < len(array); start++ {
		var diagArray []string
		for row, col := start, len(array[0])-1; row < len(array) && col >= 0; row, col = row+1, col-1 {
			diagArray = append(diagArray, array[row][col])
		}
		diagString := strings.Join(diagArray, "")
		count += strings.Count(diagString, pattern)
	}
	for start := len(array[0]) - 2; start >= 0; start-- {
		var diagArray []string
		for row, col := 0, start; row < len(array) && col >= 0; row, col = row+1, col-1 {
			diagArray = append(diagArray, array[row][col])
		}
		diagString := strings.Join(diagArray, "")
		count += strings.Count(diagString, pattern)
	}

	// now the vertical and horizontal, the easy ones
	if countVertsHoris {
		// horizontal
		for _, row := range array {
			rowString := strings.Join(row, "")
			count += strings.Count(rowString, pattern)
		}
		// vertical
		for col := 0; col < len(array[0]); col++ {
			var colArray []string
			for row := 0; row < len(array); row++ {
				colArray = append(colArray, array[row][col])
			}
			colString := strings.Join(colArray, "")
			count += strings.Count(colString, pattern)
		}
	}
	return count
}

func main() {
	fmt.Println("Day 4 of Advent Of Code 2024, j4nbob")
	input := LoadTextFromFile("day4.input")
	// part 1
	daSum := findStringPatternsInArrays(input, "XMAS", true) + findStringPatternsInArrays(input, "SAMX", true)
	fmt.Println("part1/daSum=", daSum)
	// part 2
	positions := findNonBoundaryA(input)
	//fmt.Printf("Positions of 'A' not touching boundaries: %v\n", positions)
	daSum = findXCournersBasedOnA(input, positions)
	fmt.Println("part2/daSum=", daSum)

}
