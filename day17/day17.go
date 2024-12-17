// day17 of Advent Of Code 2024 : https://adventofcode.com/2024/day/17
// j4nbob

package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

// programStringToArray converts a string to an array of strings based on spaces.
func programStringToArray(input string) []int {
	// Split the input string by spaces
	result := strings.Split(input, ",")
	intArray := make([]int, len(result))
	for i, str := range result {
		num, err := strconv.Atoi(str)
		if err != nil {
			os.Exit(1)
		}
		intArray[i] = num
	}
	return intArray
}

func Combo(input, a, b, c int) int {
	switch input {
	case 0, 1, 2, 3:
		return input
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		os.Exit(2)
	}
	return 0
}

func runProgram(Ainput int, programArray []int, part2 bool) []int {
	var out []int = make([]int, 0)
	a := Ainput
	b := 0
	c := 0

	ip := 0
	for ip < len(programArray) {
		switch programArray[ip] {
		case 0:
			a = a / int(math.Pow(2, float64(Combo(programArray[ip+1], a, b, c))))
			ip += 2
		case 1:
			b = b ^ programArray[ip+1]
			ip += 2
		case 2:
			b = Combo(programArray[ip+1], a, b, c) % 8
			ip += 2
		case 3:
			if a != 0 {
				ip = programArray[ip+1]
			} else {
				ip += 2
			}
		case 4:
			b = b ^ c
			ip += 2
		case 5:
			out = append(out, Combo(programArray[ip+1], a, b, c)%8)
			if part2 && out[len(out)-1] != programArray[len(out)-1] {
				return out
			}
			ip += 2
		case 6:
			c = a / int(math.Pow(2, float64(Combo(programArray[ip+1], a, b, c))))
			ip += 2
		case 7:
			c = a / int(math.Pow(2, float64(Combo(programArray[ip+1], a, b, c))))
			ip += 2
		default:
			os.Exit(3)
		}
	}

	return out
}

func main() {
	fmt.Println("Day 17 of Advent Of Code 2024, j4nbob")
	//input is simple today, will save time by prodiving it as a string
	var programString = "2,4,1,3,7,5,0,3,1,5,4,1,5,5,3,0"
	var programArray []int = programStringToArray(programString)
	//part1
	out := runProgram(63687530, programArray, false)
	outStr := make([]string, len(out))
	for i, num := range out {
		outStr[i] = strconv.Itoa(num)
	}
	daSum := strings.Join(outStr, ",")
	fmt.Println("part1/daSum=", daSum)

	// part2
	candidates := []int{0}
	//backwards processor
	for i := len(programArray) - 1; i >= 0; i-- {
		original := slices.Clone(candidates)
		candidates = []int{}
		for _, originalSpot := range original {
			originalSpot = originalSpot << 3
			for moduloEight := range 8 {
				a := moduloEight + originalSpot
				out := runProgram(a, programArray, true)
				if out[0] == programArray[i] {
					candidates = append(candidates, a)
				}
			}
		}

	}
	for _, candidate := range candidates {
		out := runProgram(candidate, programArray, true)
		equal := true
		for i := range out {
			if out[i] != programArray[i] {
				equal = false
				break
			}
		}
		if equal {
			//gots a winner, can bail after the first one
			fmt.Println("part2/daSum=", candidate)
			os.Exit(0)
		}
	}

}
