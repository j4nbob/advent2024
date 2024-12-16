// day16 of Advent Of Code 2024 : https://adventofcode.com/2024/day/16
// j4nbob

package main

import (
	"fmt"
	"os"
	"strings"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
)

type Coordinate struct {
	X, Y int
}

type Point struct {
	co      Coordinate
	lastDir Coordinate
}

type Step struct {
	score   int
	path    map[Coordinate]int
	co      Coordinate
	lastDir Coordinate
}

var Up = Coordinate{0, -1}
var Right = Coordinate{1, 0}
var Down = Coordinate{0, 1}
var Left = Coordinate{-1, 0}

func LoadTextFromFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}
	return string(data)
}

func CreateMaze(mazeRaw string) ([][]int, int, int, int, int) {
	maze := [][]int{}
	initX := 0
	initY := 0
	endX := 0
	endY := 0
	mazeRawX := strings.Split(mazeRaw, "\n")
	for indexRow, row := range mazeRawX {
		mazeRow := make([]int, len(row))
		for indexCol, col := range row {
			point := 0
			if col == 'S' {
				initX = indexCol
				initY = indexRow
				point = 0 // Start
			} else if col == 'E' {
				endX = indexCol
				endY = indexRow
				point = 0 // End
			} else if col == '.' {
				point = 0 // Free
			} else if col == '#' {
				point = 1 // Wall
			}
			mazeRow[indexCol] = point
		}
		maze = append(maze, mazeRow)
	}
	return maze, initX, initY, endX, endY
}

func main() {
	fmt.Println("Day 16 of Advent Of Code 2024, j4nbob")
	raw := LoadTextFromFile("day16.input")
	maze, initX, initY, finishX, finishY := CreateMaze(raw)
	beginning := Coordinate{initX, initY}
	finish := Coordinate{finishX, finishY}
	daPathScore, daBestPaths := findAllShortestPaths(maze, beginning, finish)
	fmt.Println("part1/daSum=", daPathScore)
	tiles := getTileCount(maze, daBestPaths, finish)
	fmt.Println("part2/daSum=", tiles)

}

func findAllShortestPaths(maze [][]int, start, end Coordinate) (int, map[Coordinate]int) {
	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(Step).score - b.(Step).score
	})
	priorityQueue.Enqueue(Step{score: 0, path: make(map[Coordinate]int), co: Coordinate{1, len(maze) - 2}, lastDir: Right})
	v := make(map[Point]struct{})
	for !priorityQueue.Empty() {
		element, _ := priorityQueue.Dequeue()
		c := element.(Step)
		if _, ok := v[Point{c.co, c.lastDir}]; ok {
			continue
		}
		c.path[c.co] = c.score
		if c.co == end {
			return c.score, c.path
		}
		nextSteps := getNextSteps(c, maze, v)
		for _, n := range nextSteps {
			priorityQueue.Enqueue(n)
		}
		v[Point{c.co, c.lastDir}] = struct{}{}
	}
	return -1, make(map[Coordinate]int)
}

func backup(path map[Coordinate]int) map[Coordinate]int {
	backup := make(map[Coordinate]int, len(path))
	for key, value := range path {
		backup[key] = value
	}
	return backup
}

func isValidStep(current Coordinate, input [][]int) bool {
	if current.X < 0 || current.Y < 0 || current.X >= len(input[0]) || current.Y >= len(input) {
		return false
	}
	return true
}

func getDirs(dir Coordinate) []Coordinate {
	switch dir {
	case Up:
		return []Coordinate{Up, Left, Right}
	case Right:
		return []Coordinate{Up, Down, Right}
	case Down:
		return []Coordinate{Down, Left, Right}
	case Left:
		return []Coordinate{Up, Left, Down}
	}
	return []Coordinate{}
}

func getNextSteps(current Step, maze [][]int, v map[Point]struct{}) []Step {
	possibleNext := []Step{}
	for _, dir := range getDirs(current.lastDir) {
		newPosition := Coordinate{current.co.X + dir.X, current.co.Y + dir.Y}
		if !isValidStep(newPosition, maze) {
			continue
		}
		if maze[newPosition.Y][newPosition.X] == 1 {
			continue
		}
		if _, ok := v[Point{newPosition, dir}]; ok {
			continue
		}
		score := current.score + 1
		if dir != current.lastDir {
			score += 1000
		}
		possibleNext = append(possibleNext, Step{co: newPosition, lastDir: dir, score: score, path: backup(current.path)})
	}
	return possibleNext
}

func getTileCount(maze [][]int, path map[Coordinate]int, end Coordinate) int {
	priorityQueue := pq.NewWith(func(a, b interface{}) int {
		return a.(Step).score - b.(Step).score
	})
	priorityQueue.Enqueue(Step{score: 0, path: make(map[Coordinate]int), co: Coordinate{1, len(maze) - 2}, lastDir: Right})
	v := make(map[Point]struct{})
	newCo := make(map[Coordinate]struct{})
	for !priorityQueue.Empty() {
		element, _ := priorityQueue.Dequeue()
		c := element.(Step)
		if score, ok := path[c.co]; ok && score == c.score {
			for point, _ := range c.path {
				if _, ok := path[point]; !ok {
					newCo[point] = struct{}{}
				}
			}
		}
		if _, ok := v[Point{c.co, c.lastDir}]; ok {
			continue
		}
		c.path[c.co] = c.score
		if c.co == end {
			continue
		}
		nextSteps := getNextSteps(c, maze, v)
		for _, n := range nextSteps {
			priorityQueue.Enqueue(n)
		}
		v[Point{c.co, c.lastDir}] = struct{}{}
	}
	return len(path) + len(newCo)
}
