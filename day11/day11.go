// day11 of Advent Of Code 2024 : https://adventofcode.com/2024/day/11
// j4nbob
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lookup map[string]int

func LoadTextFromFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}
	return string(data)
}

func convertBlinkInputToIntArray(blinkInput string) []int {
	parts := strings.Split(blinkInput, " ")
	intArray := make([]int, len(parts))

	for i, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			os.Exit(2)
		}
		intArray[i] = num
	}

	return intArray
}

func IsEvenNumberOfDigits(input int) bool {
	str := strconv.Itoa(input)
	return (len(str) % 2) == 0
}

func getLeftHalfandRightHalf(input int) (int, int) {
	str := strconv.Itoa(input)
	mid := len(str) / 2
	left, err := strconv.Atoi(str[:mid])
	if err != nil {
		os.Exit(3)
	}
	right, err := strconv.Atoi(str[mid:])
	if err != nil {
		os.Exit(4)
	}
	return left, right
}

func blinkOnce(input []int) []int {
	var output []int
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			output = append(output, 1)
			continue
		}

		if IsEvenNumberOfDigits(input[i]) {
			left, right := getLeftHalfandRightHalf(input[i])
			output = append(output, left)
			output = append(output, right)
			continue
		}

		newNum := input[i] * 2024
		output = append(output, newNum)
	}
	return output
}

// uses lookup map to store intermediate results
func whatLenIsInputAfterSteps(input int, steps int) int {
	key := fmt.Sprintf("%d-%d", input, steps)
	if val, exists := lookup[key]; exists {
		return val
	}

	if steps == 0 {
		lookup[key] = 1
		return 1
	}
	if input == 0 {
		lookup[key] = whatLenIsInputAfterSteps(1, steps-1)
		return lookup[key]
	}
	if IsEvenNumberOfDigits(input) {
		left, right := getLeftHalfandRightHalf(input)
		lookup[key] = whatLenIsInputAfterSteps(left, steps-1) + whatLenIsInputAfterSteps(right, steps-1)
		return lookup[key]
	}

	newNum := input * 2024
	lookup[key] = whatLenIsInputAfterSteps(newNum, steps-1)
	return lookup[key]
}

func main() {
	fmt.Println("Day 11 of Advent Of Code 2024, j4nbob")
	var contents = LoadTextFromFile("day11.input")
	blinkInput := convertBlinkInputToIntArray(contents)
	numOfBlinks := 25 // part 1
	for i := 0; i < numOfBlinks; i++ {
		blinkInput = blinkOnce(blinkInput)
	}
	fmt.Println("part1/daSum=", len(blinkInput))

	//part 2 ---- we gotta be clever and only focus on determining the lenght
	var daSum int = 0
	lookup = make(map[string]int)
	blinkInput = convertBlinkInputToIntArray(contents)
	for _, x := range blinkInput {
		daSum += whatLenIsInputAfterSteps(x, 75)
	}
	fmt.Println("part2/daSum=", daSum)

}
