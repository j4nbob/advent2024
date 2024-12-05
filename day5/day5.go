// day5 of Advent Of Code 2024 : https://adventofcode.com/2024/day/5
// j4nbob
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func readCSVintoArraysOfRows(filePath string, delimiter rune) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = delimiter

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

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Day 5 of Advent Of Code 2024, j4nbob")
	rules, err := readCSVintoArraysOfRows("day5.input", '|')
	if err != nil {
		os.Exit(1)
	}
	//decompose the rules into usable lookup maps
	mapMustBeToTheRight := make(map[int][]int)
	mapMustBeToTheLeft := make(map[int][]int)
	for _, row := range rules {
		left := row[0]
		right := row[1]
		mapMustBeToTheRight[left] = append(mapMustBeToTheRight[left], right)
		mapMustBeToTheLeft[right] = append(mapMustBeToTheLeft[right], left)
	}
	//load up our options set
	options, err := readCSVintoArraysOfRows("day5.input1", ',')
	if err != nil {
		os.Exit(2)
	}

	var daSum int = 0
	var badOptions [][]int //this is for part2

	for _, row := range options {
		disqualified := false
		for i := 0; i < len(row)-1; i++ {
			left := row[i]
			right := row[i+1]
			if rights, exists := mapMustBeToTheRight[left]; exists {
				if !contains(rights, right) {
					disqualified = true
					continue
				}
			} else {
				disqualified = true
				continue
			}

			if lefts, exists := mapMustBeToTheLeft[right]; exists {
				if !contains(lefts, left) {
					disqualified = true
					continue
				}
			} else {
				disqualified = true
				continue
			}

		}
		if !disqualified {
			middleIndex := len(row) / 2
			daSum += row[middleIndex]
		} else {
			badOptions = append(badOptions, row) //this is for part2
		}
	}

	fmt.Println("part1/daSum=", daSum)

	//part 2
	daSum = 0

	for _, row := range badOptions {

		for i := 0; i < len(row)-1; i++ {
			left := row[i]
			right := row[i+1]
			if rights, exists := mapMustBeToTheRight[left]; exists {
				if !contains(rights, right) {
					//swapping the elements
					left, right = right, left
					row[i] = left
					row[i+1] = right
					i = -1 //reset the index to start from the beginning
					continue
				}
			}

			if lefts, exists := mapMustBeToTheLeft[right]; exists {
				if !contains(lefts, left) {
					//swapping the elements
					left, right = right, left
					row[i] = left
					row[i+1] = right
					i = -1 //reset the index to start from the beginning
					continue
				}
			}

		}
		middleIndex := len(row) / 2
		daSum += row[middleIndex]
	}

	fmt.Println("part2/daSum=", daSum)

}
