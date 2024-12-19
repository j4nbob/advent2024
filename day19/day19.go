// day19 of Advent Of Code 2024 : https://adventofcode.com/2024/day/19
// j4nbob

package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadTextFromFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}
	return string(data)
}

var lookup = make(map[string]int)

// returns ways to get from patterns to target
func findTargetsFromPatterns(patterns []string, desiredTarget string) int {
	if _, exists := lookup[desiredTarget]; exists {
		return lookup[desiredTarget]
	}
	ways := 0
	if desiredTarget == "" {
		ways++
	}
	for _, pattern := range patterns {
		if strings.HasPrefix(desiredTarget, pattern) {
			ways += findTargetsFromPatterns(patterns, desiredTarget[len(pattern):])
		}
	}
	lookup[desiredTarget] = ways
	return ways
}

func main() {
	fmt.Println("Day 19 of Advent Of Code 2024, j4nbob")
	raw := LoadTextFromFile("day19.input")
	patternX := strings.Split(raw, "\n\n")
	patternsX := patternX[0]
	patterns := strings.Split(patternsX, ", ")
	targetsX := patternX[1]
	targets := strings.Split(targetsX, "\n")

	//fmt.Println("patterns:", patterns)
	//fmt.Println("desires:", targets)
	daSum := 0
	for _, target := range targets {
		if findTargetsFromPatterns(patterns, target) > 0 {
			daSum++
		}

	}
	fmt.Println("part1/daSum :", daSum)

	// part 2
	daSum = 0
	for _, target := range targets {
		daSum += findTargetsFromPatterns(patterns, target)
	}
	fmt.Println("part2/daSum :", daSum)

}
