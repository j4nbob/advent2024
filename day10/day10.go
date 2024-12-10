// day10 of Advent Of Code 2024 : https://adventofcode.com/2024/day/10
// j4nbob
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var topo [][]int

type Coordinate struct {
	x, y int
}

func LoadTextFromFile(filePath string) ([][]int, []int) {
	file, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var result [][]int
	uniqueRunes := make(map[rune]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineArray []int
		for _, char := range line {
			if char != '.' {
				//lineArray = append(lineArray, char)
				uniqueRunes[char] = struct{}{}
			}
			intValue, err := strconv.Atoi(string(char))
			if err != nil {
				os.Exit(3)
			}
			lineArray = append(lineArray, intValue)

		}
		result = append(result, lineArray)
	}

	if err := scanner.Err(); err != nil {
		os.Exit(2)
	}

	var uniqueRuneList []int
	for r := range uniqueRunes {
		intValue, err := strconv.Atoi(string(r))
		if err != nil {
			os.Exit(3)
		}
		uniqueRuneList = append(uniqueRuneList, intValue)
	}
	return result, uniqueRuneList
}

func FindRuneLocations(a [][]int, r int, mustBeEdge bool) [][2]int {
	var locations [][2]int
	for y, row := range a {
		for x, char := range row {
			if char == r {
				if mustBeEdge {
					if x == 0 || y == 0 || x == len(row)-1 || y == len(a)-1 {
						locations = append(locations, [2]int{x, y})
					}
				} else {
					locations = append(locations, [2]int{x, y})
				}

			}
		}
	}
	return locations
}

func GoUpTheTopo(elevation, x, y int, uniqueTopoTops *map[Coordinate]bool, trackUniqueTops bool) int {
	if elevation == 9 {
		if trackUniqueTops {
			coord := Coordinate{x, y}
			if _, exists := (*uniqueTopoTops)[coord]; !exists {
				(*uniqueTopoTops)[coord] = true
				//fmt.Println("Found a new top at ", x, y)
				return 1
			} else {
				return 0 //don't count it more than once
			}
		} else {
			return 1
		}
	} else {
		counter := 0
		nextElevation := elevation + 1
		//calculate indices up, right, down, left
		//if the indices are within bounds, and the value is greater than the elevation, then return goupthetopo
		deadEnd := true
		if y > 0 && topo[y-1][x] == nextElevation {
			counter += GoUpTheTopo(nextElevation, x, y-1, uniqueTopoTops, trackUniqueTops)
			deadEnd = false
		}
		if x < len(topo[y])-1 && topo[y][x+1] == nextElevation {
			counter += GoUpTheTopo(nextElevation, x+1, y, uniqueTopoTops, trackUniqueTops)
			deadEnd = false
		}
		if y < len(topo)-1 && topo[y+1][x] == nextElevation {
			counter += GoUpTheTopo(nextElevation, x, y+1, uniqueTopoTops, trackUniqueTops)
			deadEnd = false
		}
		if x > 0 && topo[y][x-1] == nextElevation {
			counter += GoUpTheTopo(nextElevation, x-1, y, uniqueTopoTops, trackUniqueTops)
			deadEnd = false
		}
		if deadEnd {
			return 0
		} else {
			return counter
		}
	}
}

func main() {
	fmt.Println("Day 10 of Advent Of Code 2024, j4nbob")
	var uniqueRuneList []int
	topo, uniqueRuneList = LoadTextFromFile("day10.input")
	//fmt.Println(topo)
	//fmt.Println(uniqueRuneList)
	if len(uniqueRuneList) != 10 {
		//topo does not have 10 unique runes
		os.Exit(4)
	}
	trailheads := FindRuneLocations(topo, 0, false)
	if len(trailheads) == 0 {
		//we aint got no trailheads
		os.Exit(5)
	}
	//fmt.Println(trailheads)

	var daSum int = 0
	for _, trailhead := range trailheads {
		//fmt.Println(trailheadsIndex, trailhead)
		uniqueTopoTops := make(map[Coordinate]bool)
		trailWorthyness := GoUpTheTopo(0, trailhead[0], trailhead[1], &uniqueTopoTops, true)
		daSum += trailWorthyness
		//fmt.Println("Trailworthyness: ", trailWorthyness)
	}
	fmt.Println("part1/daSum=", daSum)
	daSum = 0
	for _, trailhead := range trailheads {
		//fmt.Println(trailheadsIndex, trailhead)
		uniqueTopoTops := make(map[Coordinate]bool)
		trailWorthyness := GoUpTheTopo(0, trailhead[0], trailhead[1], &uniqueTopoTops, false)
		daSum += trailWorthyness
		//fmt.Println("Trailworthyness: ", trailWorthyness)
	}
	fmt.Println("part2/daSum=", daSum)

}
