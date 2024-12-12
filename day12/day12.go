// day12 of Advent Of Code 2024 : https://adventofcode.com/2024/day/12
// j4nbob
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type boundary struct {
	row, col int
	counted  bool
}

type Coordinate struct {
	row, col int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Shape struct {
	coords     map[Coordinate]bool
	boundaries map[Direction][]boundary
	symbol     byte
	size       int
}

var originalGrid [][]byte

func LoadTextFromFile(filePath string) ([][]byte, []byte) {
	file, err := os.Open(filePath)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	var result [][]byte
	uniquebytes := make(map[byte]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineArray []byte
		for _, char := range line {
			lineArray = append(lineArray, byte(char))
			uniquebytes[byte(char)] = true
		}
		result = append(result, lineArray)
	}

	if err := scanner.Err(); err != nil {
		os.Exit(2)
	}

	var uniquebyteList []byte
	for r := range uniquebytes {
		uniquebyteList = append(uniquebyteList, r)
	}
	return result, uniquebyteList
}

func DepthFirstSearch(g [][]byte, i, j int, marker byte) {
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	if i < 0 || j < 0 || i >= len(g) || j >= len(g[0]) {
		return
	}
	if g[i][j] == 1 {
		g[i][j] = marker
		for _, d := range dirs {
			DepthFirstSearch(g, i+d[0], j+d[1], marker)
		}
	}
}

func CopyGridRowByRow(g [][]byte) [][]byte {
	rows := len(g)
	cols := len(g[0])
	newGrid := make([][]byte, rows)
	for i := range newGrid {
		newGrid[i] = make([]byte, cols)
		copy(newGrid[i], g[i])
	}
	return newGrid
}

func detectAllShapes(g [][]byte) (int, [][]byte) {
	cnt := 0
	var marker byte = 2
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			//gCurrent := CopyGridRowByRow(gBackup)
			if g[i][j] == 1 {
				DepthFirstSearch(g, i, j, marker)
				cnt++
				//shapeGatherer[marker] = gCurrent
				marker++

			}
		}
	}

	return cnt, g
}

func CopyOriginalGrid(g [][]byte, char byte) [][]byte {
	rows := len(g)
	cols := len(g[0])
	newGrid := make([][]byte, rows)
	for i := range newGrid {
		newGrid[i] = make([]byte, cols)
		for j := range g[i] {
			if g[i][j] == char {
				newGrid[i][j] = 1
			} else {
				newGrid[i][j] = 0
			}
		}
	}
	return newGrid
}

func FindPerimeterLength(grid [][]byte, marker byte) int {
	rows := len(grid)
	cols := len(grid[0])
	perimeter := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == marker {
				// Check up
				if i == 0 || grid[i-1][j] == 0 {
					perimeter++
				}
				// Check down
				if i == rows-1 || grid[i+1][j] == 0 {
					perimeter++
				}
				// Check left
				if j == 0 || grid[i][j-1] == 0 {
					perimeter++
				}
				// Check right
				if j == cols-1 || grid[i][j+1] == 0 {
					perimeter++
				}
			}
		}
	}

	return perimeter
}

func FindArea(g [][]byte, marker byte) int {
	area := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			if g[i][j] == marker {
				area++
			}
		}
	}
	return area
}

func detectAllShapes2(grid [][]byte) []Shape {
	s := make(map[Coordinate]bool)
	var gs []Shape

	for row := range grid {
		for col := range grid[row] {
			coord := Coordinate{row, col}
			if !s[coord] {
				if g := exploreShape(grid, coord, s); g.size > 0 {
					gs = append(gs, g)
				}
			}
		}
	}

	return gs
}

func isValidCoord(grid [][]byte, c Coordinate) bool {
	return c.row >= 0 && c.row < len(grid) && c.col >= 0 && c.col < len(grid[0])
}

func AdjacentCoord(c Coordinate) []Coordinate {
	return []Coordinate{
		{c.row - 1, c.col},
		{c.row + 1, c.col},
		{c.row, c.col + 1},
		{c.row, c.col - 1},
	}
}

func exploreShape(grid [][]byte, start Coordinate, seen map[Coordinate]bool) Shape {
	if seen[start] || !isValidCoord(grid, start) {
		return Shape{}
	}

	symbol := grid[start.row][start.col]
	Shape := Shape{
		symbol:     symbol,
		coords:     make(map[Coordinate]bool),
		boundaries: make(map[Direction][]boundary),
	}

	queue := []Coordinate{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if seen[curr] || !isValidCoord(grid, curr) || grid[curr.row][curr.col] != symbol {
			continue
		}

		seen[curr] = true
		Shape.coords[curr] = true
		Shape.size++

		isBoundary(grid, curr, symbol, North, &Shape)
		isBoundary(grid, curr, symbol, South, &Shape)
		isBoundary(grid, curr, symbol, East, &Shape)
		isBoundary(grid, curr, symbol, West, &Shape)

		for _, next := range AdjacentCoord(curr) {
			if isValidCoord(grid, next) && grid[next.row][next.col] == symbol {
				queue = append(queue, next)
			}
		}
	}

	return Shape
}

func isBoundary(grid [][]byte, pos Coordinate, symbol byte, dir Direction, Shape *Shape) {
	var isEdge bool
	switch dir {
	case North:
		isEdge = pos.row == 0 || grid[pos.row-1][pos.col] != symbol
	case South:
		isEdge = pos.row == len(grid)-1 || grid[pos.row+1][pos.col] != symbol
	case East:
		isEdge = pos.col == len(grid[0])-1 || grid[pos.row][pos.col+1] != symbol
	case West:
		isEdge = pos.col == 0 || grid[pos.row][pos.col-1] != symbol
	}
	if isEdge {
		Shape.boundaries[dir] = append(Shape.boundaries[dir], boundary{pos.row, pos.col, true})
	}
}

func pruneAxis(bounds []boundary, sortByCol bool) {
	if sortByCol {
		sort.Slice(bounds, func(i, j int) bool { return bounds[i].col < bounds[j].col })
	} else {
		sort.Slice(bounds, func(i, j int) bool { return bounds[i].row < bounds[j].row })
	}

	for i := range bounds {
		var next int
		if sortByCol {
			next = bounds[i].col
		} else {
			next = bounds[i].row
		}

		for {
			next++
			found := false
			for j := range bounds {
				var matches bool
				if sortByCol {
					matches = bounds[j].col == next && bounds[j].row == bounds[i].row
				} else {
					matches = bounds[j].row == next && bounds[j].col == bounds[i].col
				}
				if matches {
					bounds[j].counted = false
					found = true
					break
				}
			}
			if !found {
				break
			}
		}
	}
}

func prune(Shape *Shape) {
	for dir := range Shape.boundaries {
		if dir == North || dir == South {
			pruneAxis(Shape.boundaries[dir], true)
		} else {
			pruneAxis(Shape.boundaries[dir], false)
		}
	}
}

func FindSides(Shape Shape) int {
	count := 0
	for _, bounds := range Shape.boundaries {
		for _, b := range bounds {
			if b.counted {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println("Day 12 of Advent Of Code 2024, j4nbob")
	var uniquebyteList []byte
	originalGrid, uniquebyteList = LoadTextFromFile("day12.input")

	var daSum int = 0
	for _, letter := range uniquebyteList {
		workingGrid := CopyOriginalGrid(originalGrid, letter)
		numOfShapes, g := detectAllShapes(workingGrid)
		for marker := byte(2); marker < byte(2+numOfShapes); marker++ {
			perimeter := FindPerimeterLength(g, marker)
			area := FindArea(g, marker)
			daSum += perimeter * area
		}

	}
	fmt.Println("part1/daSum=", daSum)

	// Part 2
	daSum = 0
	shapes := detectAllShapes2(originalGrid) //taking a different approach here
	for _, shape := range shapes {
		prune(&shape)
		area := shape.size
		sides := FindSides(shape)
		daSum += area * sides
	}
	fmt.Println("part2/daSum=", daSum)
}
