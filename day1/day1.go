// day1 of Advent Of Code 2024 : https://adventofcode.com/2024/day/1
// j4nbob
package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func main() {
	fmt.Println("Day 1 of Advent Of Code 2024, j4nbob")

	// we start off quick and dirty, Bad Santa style
	file, err := os.Open("day1.input")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var array1, array2 []int
	for _, record := range records {
		if len(record) >= 2 {
			leftValue, err := strconv.Atoi(record[0])
			if err != nil {
				os.Exit(2)
			}
			rightValue, err := strconv.Atoi(record[1])
			if err != nil {
				os.Exit(2)
			}
			array1 = append(array1, leftValue)
			array2 = append(array2, rightValue)
		}
	}
	var leftSorted, rightSorted []int

	leftSorted = make([]int, len(array1))
	copy(leftSorted, array1)
	sort.Ints(leftSorted)
	rightSorted = make([]int, len(array2))
	copy(rightSorted, array2)
	sort.Ints(rightSorted)

	var daSum int
	daSum = 0
	for i := 0; i < len(leftSorted); i++ {
		distance := diff(leftSorted[i], rightSorted[i])
		daSum += distance
	}
	fmt.Println("part1/daSum=", daSum)

	//part 2
	counter := make(map[int]int)
	for _, item := range rightSorted {
		counter[item]++
	}
	daSum = 0
	for _, daElement := range leftSorted {
		if value, exists := counter[daElement]; exists {
			daSum = daSum + (value * daElement)
		}
	}
	fmt.Println("part2/daSum=", daSum)
}
