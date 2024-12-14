// day14 of Advent Of Code 2024 : https://adventofcode.com/2024/day/14
// j4nbob

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y int
}
type Robot struct {
	pos   Coordinate
	xStep int
	yStep int
}

var robots []Robot

const maxX = 101 //11
const maxY = 103 //7
const maxClicks = 100

func LoadTextFromFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}
	return string(data)
}

func parseLine(line string) (Coordinate, Coordinate) {
	parts := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(line, "p=", ""), "v=", ""))
	pParts := strings.Split(parts[0], ",")
	vParts := strings.Split(parts[1], ",")
	pX, _ := strconv.Atoi(pParts[0])
	pY, _ := strconv.Atoi(pParts[1])
	vX, _ := strconv.Atoi(vParts[0])
	vY, _ := strconv.Atoi(vParts[1])
	p := Coordinate{X: pX, Y: pY}
	v := Coordinate{X: vX, Y: vY}
	return p, v
}

func MoveRobotPosByXandY(r *Robot) {
	r.pos.X += r.xStep
	r.pos.Y += r.yStep

	if r.pos.X < 0 {
		r.pos.X = maxX + r.pos.X
	}
	if r.pos.Y < 0 {
		r.pos.Y = maxY + r.pos.Y
	}
	if r.pos.X > maxX-1 {
		r.pos.X = r.pos.X - maxX
	}
	if r.pos.Y > maxY-1 {
		r.pos.Y = r.pos.Y - maxY
	}
}

// must be odd numbers
func findMidpointCoordinates(x int, y int) Coordinate {
	midX := x / 2
	midY := y / 2
	return Coordinate{X: midX, Y: midY}
}

func CountRobotsInFourQuadrants(maxX, maxY, midX, midY int) (int, int, int, int) {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		if robot.pos.X != midX && robot.pos.Y != midY {
			if robot.pos.X < midX && robot.pos.Y < midY {
				q1++
			} else if robot.pos.X > midX && robot.pos.Y < midY {
				q2++
			} else if robot.pos.X < midX && robot.pos.Y > midY {
				q3++
			} else {
				q4++
			}
		}
	}

	return q1, q2, q3, q4
}

// for verification only / display for the guess work for part 2
func renderRobotGrid() [][]int {
	grid := make([][]int, maxY)
	for i := range grid {
		grid[i] = make([]int, maxX)
	}
	for _, robot := range robots {
		grid[robot.pos.Y][robot.pos.X] = 1
	}
	return grid
}

// hate guess work
func doesGridContainTree(grid *[][]int) bool {
	rows := len(*grid)
	if rows == 0 {
		return false
	}
	cols := len((*grid)[0])
	foundColumn := false
	// Check rows for 20 consecutive 1s
outerLoop:
	for _, row := range *grid {
		count := 0
		for _, cell := range row {
			if cell == 1 {
				count++
				if count >= 20 {
					//return true
					foundColumn = true
					break outerLoop
				}
			} else {
				count = 0
			}
		}
	}

	// Check columns for 20 consecutive 1s
	for col := 0; col < cols; col++ {
		count := 0
		for row := 0; row < rows; row++ {
			if (*grid)[row][col] == 1 {
				count++
				if count >= 20 {
					if foundColumn {
						return true
					}
				}
			} else {
				count = 0
			}
		}
	}

	return false
}

func main() {
	fmt.Println("Day 14 of Advent Of Code 2024, j4nbob")
	raw := LoadTextFromFile("day14.input")
	arr := strings.Split(raw, "\n")
	for _, val := range arr {
		p, v := parseLine(val)
		robots = append(robots, Robot{pos: p, xStep: v.X, yStep: v.Y})
	}

	midPoint := findMidpointCoordinates(maxX, maxY)
	for rindex, _ := range robots {
		for i := 0; i < maxClicks; i++ {
			MoveRobotPosByXandY(&robots[rindex])
		}
	}
	a, b, c, d := CountRobotsInFourQuadrants(maxX, maxY, midPoint.X, midPoint.Y)
	safetyFactor := a * b * c * d
	fmt.Println("part1/daSum=", safetyFactor)

	//part 2 we go  10000 steps and guess what a tree looks like
	//reset to beginning
	robots = nil
	for _, val := range arr {
		p, v := parseLine(val)
		robots = append(robots, Robot{pos: p, xStep: v.X, yStep: v.Y})
	}

	for i := 0; i < maxClicks+10000; i++ {
		for rindex, _ := range robots {
			MoveRobotPosByXandY(&robots[rindex])
		}
		previewGrid := renderRobotGrid()
		if doesGridContainTree(&previewGrid) {
			fmt.Println("part2/daSum=", i+1) // number of seconds ticked
			break
		}

	}

}
