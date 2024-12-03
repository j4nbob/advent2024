// day3 of Advent Of Code 2024 : https://adventofcode.com/2024/day/3
// j4nbob
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadTextFromFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}
	return string(data)
}

func breakStringIntoSegments(input, delimiter string) []string {
	return strings.Split(input, delimiter)

}

func returnSumOfMultiplications(contents string) int {
	var daSum int = 0
	segments := breakStringIntoSegments(contents, "mul(")
	for _, segment := range segments {
		//fmt.Println(segment)
		commaPos := strings.Index(segment, ",")
		if commaPos == -1 {
			continue
		}

		segmentString := string(segment[:commaPos])
		segmentInt, err := strconv.Atoi(segmentString)
		if err != nil {
			continue
		}

		restString := string(segment[commaPos+1:])
		closingPos := strings.Index(restString, ")")
		if closingPos == -1 {
			continue
		}

		secondString := string(restString[:closingPos])
		secondSegmentInt, err := strconv.Atoi(secondString)
		if err != nil {
			continue
		}
		daSum += segmentInt * secondSegmentInt
	}
	return daSum
}

func main() {
	fmt.Println("Day 3 of Advent Of Code 2024, j4nbob")
	var contents = LoadTextFromFile("day3.input")

	var daSum int = 0
	daSum = returnSumOfMultiplications(contents)
	fmt.Println("part1/daSum=", daSum)

	//part 2
	daSum = 0
	segments := breakStringIntoSegments(contents, "don't()")
	for index, segment := range segments {
		if index == 0 {
			daSum = returnSumOfMultiplications(segment)
			continue
		}
		resumePos := strings.Index(segment, "do()")
		if resumePos == -1 {
			continue
		}
		restString := string(segment[resumePos+4:])
		daSum = daSum + returnSumOfMultiplications(restString)
	}
	fmt.Println("part2/daSum=", daSum)
}
