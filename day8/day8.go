// day8 of Advent Of Code 2024 : https://adventofcode.com/2024/day/8
// j4nbob
package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadTextFromFile(filePath string) ([][]rune, []rune) {
	file, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var result [][]rune
	uniqueRunes := make(map[rune]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineArray []rune
		for _, char := range line {
			if char != '.' {
				//lineArray = append(lineArray, char)
				uniqueRunes[char] = struct{}{}
			}
			lineArray = append(lineArray, char)

		}
		result = append(result, lineArray)
	}

	if err := scanner.Err(); err != nil {
		os.Exit(2)
	}

	var uniqueRuneList []rune
	for r := range uniqueRunes {
		uniqueRuneList = append(uniqueRuneList, r)
	}
	return result, uniqueRuneList
}

func FindRuneLocations(a [][]rune, r rune) [][2]int {
	var locations [][2]int
	for y, row := range a {
		for x, char := range row {
			if char == r {
				locations = append(locations, [2]int{x, y})
			}
		}
	}
	return locations
}

func CalculateSlopeAndDistanceBetweenTwoPoints(x1, y1, x2, y2 int) (int, int) {
	x := x2 - x1
	y := y2 - y1
	return x, y
}

func CreateEmptyRuneArray(rows, cols int) [][]rune {
	emptyArray := make([][]rune, rows)
	for i := range emptyArray {
		emptyArray[i] = make([]rune, cols)
		for j := range emptyArray[i] {
			emptyArray[i][j] = '.'
		}
	}
	return emptyArray
}
func main() {
	fmt.Println("Day 8 of Advent Of Code 2024, j4nbob")
	input, uniqueRuneList := LoadTextFromFile("day8.input")
	emptyArray := CreateEmptyRuneArray(len(input), len(input[0]))
	var daSum int = 0
	for _, r := range uniqueRuneList {
		locations := FindRuneLocations(input, r)
		if len(locations) < 2 {
			continue
		}

		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				//fmt.Printf("Comparing (%d, %d) with (%d, %d)\n", locations[i][0], locations[i][1], locations[j][0], locations[j][1])
				slopeX, slopeY := CalculateSlopeAndDistanceBetweenTwoPoints(locations[i][0], locations[i][1], locations[j][0], locations[j][1])
				upperX := locations[i][0] - slopeX
				upperY := locations[i][1] - slopeY
				lowerX := locations[j][0] + slopeX
				lowerY := locations[j][1] + slopeY
				if upperX >= 0 && upperY >= 0 && upperX < len(input) && upperY < len(input) {
					if emptyArray[upperY][upperX] != '#' {
						//fmt.Println("upper added", upperX, upperY)
						daSum += 1
						emptyArray[upperY][upperX] = '#'
					}

				}
				if lowerX < len(input) && lowerY < len(input) && lowerX >= 0 && lowerY >= 0 {
					if emptyArray[lowerY][lowerX] != '#' {
						//fmt.Println("lower added", lowerX, lowerY)
						daSum += 1
						emptyArray[lowerY][lowerX] = '#'
					}
				}

			}
		}

	}
	fmt.Println("part1/daSum=", daSum)

	//part 2
	emptyArray = CreateEmptyRuneArray(len(input), len(input[0]))
	daSum = 0
	for _, r := range uniqueRuneList {
		locations := FindRuneLocations(input, r)
		if len(locations) < 2 {
			continue
		}

		for i := 0; i < len(locations); i++ {
			if emptyArray[locations[i][1]][locations[i][0]] != '#' {
				daSum += 1
				emptyArray[locations[i][1]][locations[i][0]] = '#'
			}

			for j := i + 1; j < len(locations); j++ {
				if emptyArray[locations[j][1]][locations[j][0]] != '#' {
					daSum += 1
					emptyArray[locations[j][1]][locations[j][0]] = '#'
				}

				//fmt.Printf("Comparing (%d, %d) with (%d, %d)\n", locations[i][0], locations[i][1], locations[j][0], locations[j][1])
				slopeX, slopeY := CalculateSlopeAndDistanceBetweenTwoPoints(locations[i][0], locations[i][1], locations[j][0], locations[j][1])

				for k := 1; k < 50; k++ {
					upperX := locations[i][0] - (slopeX * k)
					upperY := locations[i][1] - (slopeY * k)
					if upperX >= 0 && upperY >= 0 && upperX < len(input) && upperY < len(input) {
						if emptyArray[upperY][upperX] != '#' {
							//fmt.Println("upper added", upperX, upperY)
							daSum += 1
							emptyArray[upperY][upperX] = '#'
						}

					}
				} //k

				for k := 1; k < 50; k++ {
					lowerX := locations[j][0] + (slopeX * k)
					lowerY := locations[j][1] + (slopeY * k)
					if lowerX < len(input) && lowerY < len(input) && lowerX >= 0 && lowerY >= 0 {
						if emptyArray[lowerY][lowerX] != '#' {
							//fmt.Println("lower added", lowerX, lowerY)
							daSum += 1
							emptyArray[lowerY][lowerX] = '#'
						}
					}
				} //k

			}
		}

	}
	fmt.Println("part2/daSum=", daSum)

}

/*
func WriteRuneArrayToFile(array [][]rune, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, row := range array {
		for _, char := range row {
			writer.WriteRune(char)
		}
		writer.WriteRune('\n')
	}
	writer.Flush()
}
*/
