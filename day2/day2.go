// day2 of Advent Of Code 2024 : https://adventofcode.com/2024/day/2
// j4nbob
package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

func diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func isArrayIncreasing(array []int) bool {
	for i := 0; i < len(array); i++ {
		if i+1 < len(array) {
			if array[i] > array[i+1] {
				return false
			}
		}
	}
	return true
}

func isArrayDecreasing(array []int) bool {
	for i := 0; i < len(array); i++ {
		if i+1 < len(array) {
			if array[i] < array[i+1] {
				return false
			}
		}
	}
	return true
}

func verifyElementDifferenceisBetweenOneAndThree(array []int) bool {
	for i := 0; i < len(array)-1; i++ {
		d := diff(array[i+1], array[i])
		if d < 1 || d > 3 {
			return false
		}
	}
	return true
}

func readCSVintoArraysOfRows(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var rows [][]int
	for _, record := range records {
		var row []int
		for _, value := range record {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			row = append(row, intValue)
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func testArrayUntilValidOrExhaustedTries(array []int) bool {
	for i := 0; i < len(array); i++ {
		// Create a copy of the array without the i-th element
		modifiedArray := make([]int, len(array)-1)
		copy(modifiedArray, array[:i])
		copy(modifiedArray[i:], array[i+1:])
		// Run the test function on the modified array
		if isArrayIncreasing(modifiedArray) || isArrayDecreasing(modifiedArray) { // elimite the case where the row is neither increasing nor decreasing
			if verifyElementDifferenceisBetweenOneAndThree(modifiedArray) {
				return true
			}
		}

	}
	return false
}

func main() {
	fmt.Println("Day 2 of Advent Of Code 2024, j4nbob")

	transposed, err := readCSVintoArraysOfRows("day2.input")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//part 1
	var daSum int = 0
	for _, row := range transposed {
		if isArrayIncreasing(row) || isArrayDecreasing(row) { // elimite the case where the row is neither increasing nor decreasing
			if verifyElementDifferenceisBetweenOneAndThree(row) {
				daSum++
			}
		}

	}
	fmt.Println("part1/daSum=", daSum)

	//part 2
	daSum = 0
	for _, row := range transposed {
		if testArrayUntilValidOrExhaustedTries(row) {
			daSum++
		}

	}
	fmt.Println("part2/daSum=", daSum)
}
