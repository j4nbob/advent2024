// day6 of Advent Of Code 2024 : https://adventofcode.com/2024/day/6
// j4nbob
package main

import (
	"bufio"
	"fmt"
	"os"
)

// #.^
var currentPosX = 0
var currentPosY = 0

func LoadTextFromFile(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	var result [][]int
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		var lineArray []int
		for x, char := range line {
			if char == '#' {
				lineArray = append(lineArray, 1)
			} else if char == '^' {
				currentPosX = x
				currentPosY = y
				lineArray = append(lineArray, 0)
			} else {
				lineArray = append(lineArray, 0)
			}

		}
		result = append(result, lineArray)
		y++
	}
	if err := scanner.Err(); err != nil {
		os.Exit(2)
	}

	return result
}
func makeMapKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

// function moves CurrentX and CurrentY based on the direction
func moveCurrentXandY(a [][]int, m map[string]bool, dirrection int) bool {
	maxX := len(a[0])
	maxY := len(a)

	switch dirrection {
	case 0: //up
		for i := currentPosY; i >= 0; i-- {
			if a[i][currentPosX] == 1 {
				return true
			}
			currentPosY = i
			m[makeMapKey(currentPosX, currentPosY)] = true
		}
	case 1: //right
		for i := currentPosX; i < maxX; i++ {
			if a[currentPosY][i] == 1 {
				return true
			}
			currentPosX = i
			m[makeMapKey(currentPosX, currentPosY)] = true
		}
	case 2: //down
		for i := currentPosY; i < maxY; i++ {
			if a[i][currentPosX] == 1 {
				return true
			}
			currentPosY = i
			m[makeMapKey(currentPosX, currentPosY)] = true
		}
	case 3: //left
		for i := currentPosX; i >= 0; i-- {
			if a[currentPosY][i] == 1 {
				return true
			}
			currentPosX = i
			m[makeMapKey(currentPosX, currentPosY)] = true
		}
	}

	return false
}

func main() {
	fmt.Println("Day 6 of Advent Of Code 2024, j4nbob")
	input := LoadTextFromFile("day6.input")
	var saveCurrentPosX = currentPosX
	var saveCurrentPosY = currentPosY
	var mapOfUniqueVisits = make(map[string]bool)
	mapOfUniqueVisits[makeMapKey(currentPosX, currentPosY)] = true
	spinner := 0
	for {
		direction := spinner % 4
		if !moveCurrentXandY(input, mapOfUniqueVisits, direction) {
			break // we are off the map
		}
		spinner++
	}
	fmt.Println("part1/daSum=", len(mapOfUniqueVisits))

	//part2
	currentPosX = saveCurrentPosX
	currentPosY = saveCurrentPosY
	daSum := 0

	//iterate over input again, but this time we will keep track of the turns, so we can detect looping
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ { //brute forces all the possible obstacles ... this shall run for a while
			if input[i][j] == 0 {
				if j == saveCurrentPosX && i == saveCurrentPosY {
					continue //skip the starting point as we cannot put bricker there
				}
				input[i][j] = 1
				currentPosX = saveCurrentPosX
				currentPosY = saveCurrentPosY
				var p2UniqueVisits = make(map[string]bool)
				p2UniqueVisits[makeMapKey(currentPosX, currentPosY)] = true
				mapOfTurnPoints := make(map[string]int) //map of counter of points we turn, is used to detect looping
				spinner = 0
				for {
					direction := spinner % 4
					if !moveCurrentXandY(input, p2UniqueVisits, direction) {
						break // we are off the map
					}
					if _, exists := mapOfTurnPoints[makeMapKey(currentPosX, currentPosY)]; exists {
						mapOfTurnPoints[makeMapKey(currentPosX, currentPosY)]++
						if mapOfTurnPoints[makeMapKey(currentPosX, currentPosY)] > 2 {
							daSum++
							break
						}
					} else {
						mapOfTurnPoints[makeMapKey(currentPosX, currentPosY)] = 1
					}
					spinner++
				}
				input[i][j] = 0 // flip it back
			}
		}
	}
	fmt.Println("part2/daSum=", daSum)
}
