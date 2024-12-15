// day15 of Advent Of Code 2024 : https://adventofcode.com/2024/day/15
// j4nbob

package main

import (
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	X, Y int
}

type TerrainType byte
type MoveType byte

const (
	Wall      TerrainType = 35
	Box       TerrainType = 79
	Free      TerrainType = 46
	Prankster TerrainType = 64
	BoxL      TerrainType = 91
	BoxR      TerrainType = 93
)

const (
	Up    MoveType = 94
	Down  MoveType = 118
	Right MoveType = 62
	Left  MoveType = 60
)

func LoadTextFromFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}
	return string(data)
}

var robotPos Coordinate
var maze [][]byte

func CreateMaze(mazeRaw string) ([][]byte, int, int) {
	maze := [][]byte{}
	initX := 0
	initY := 0
	mazeRawX := strings.Split(mazeRaw, "\n")
	for indexRow, row := range mazeRawX {
		mazeRow := []byte(row)
		maze = append(maze, mazeRow)
		for indexCol, col := range mazeRow {
			if col == 64 {
				initX = indexCol
				initY = indexRow
			}
		}
	}
	return maze, initX, initY
}

func CreateMoves(movesRaw string) []MoveType {
	moves := []MoveType{}
	for _, move := range movesRaw {
		switch MoveType(move) {
		case Up, Down, Right, Left:
			moves = append(moves, MoveType(move))
		default:
			//fmt.Println("SKIPPING")
		}
	}
	return moves
}

func getCoordinatesOfBoxesGPSSum(maze *[][]byte, matchOnByte byte) int {
	boxes := []Coordinate{}
	sumCoordinates := 0
	for y, row := range *maze {
		for x, col := range row {
			if col == matchOnByte {
				boxes = append(boxes, Coordinate{X: x, Y: y})
			}
		}
	}
	for _, box := range boxes {
		sumCoordinates += (box.X + (100 * box.Y))
	}
	return sumCoordinates
}

func AttemptMoveByPushingBoxes(maze *[][]byte, inDirection MoveType, currentPosition Coordinate, attemptedPostion Coordinate) bool {
	//simple case
	if (*maze)[attemptedPostion.Y][attemptedPostion.X] == byte(TerrainType(Free)) {
		return true
	}
	//box case
	if (*maze)[attemptedPostion.Y][attemptedPostion.X] == byte(TerrainType(Box)) {
		currentPosition = attemptedPostion
		switch inDirection {
		case Up:
			attemptedPostion = Coordinate{X: currentPosition.X, Y: currentPosition.Y - 1}
		case Down:
			attemptedPostion = Coordinate{X: currentPosition.X, Y: currentPosition.Y + 1}
		case Right:
			attemptedPostion = Coordinate{X: currentPosition.X + 1, Y: currentPosition.Y}
		case Left:
			attemptedPostion = Coordinate{X: currentPosition.X - 1, Y: currentPosition.Y}
		}

		if AttemptMoveByPushingBoxes(maze, inDirection, currentPosition, attemptedPostion) {
			(*maze)[currentPosition.Y][currentPosition.X] = byte(TerrainType(Free))
			(*maze)[attemptedPostion.Y][attemptedPostion.X] = byte(TerrainType(Box))
			return true
		}
	}

	if (*maze)[attemptedPostion.Y][attemptedPostion.X] == byte(TerrainType(Wall)) {
		return false // wall case
	}

	return false
}

func widenMaze(board [][]byte) [][]byte {
	newMaze := make([][]byte, 0, len(board))
	for i, line := range board {
		newMaze = append(newMaze, make([]byte, 0, 2*len(line)))
		for _, terrain := range line {
			switch terrain {
			case '.':
				newMaze[i] = append(newMaze[i], '.', '.')
			case '@':
				newMaze[i] = append(newMaze[i], '@', '.')
			case '#':
				newMaze[i] = append(newMaze[i], '#', '#')
			case 'O':
				newMaze[i] = append(newMaze[i], '[', ']')
			}
		}
	}
	return newMaze
}

// part 2 moving is more complex
func AttemptMoveByPushingWiderBoxes(maze [][]byte, x, y, directionextX, directionextY int, attemptMoveOnly bool) bool {
	nextX, nextY := x+directionextX, y+directionextY
	switch true {
	case (maze[nextX][nextY] == byte(TerrainType(BoxL)) || maze[nextX][nextY] == byte(TerrainType(BoxR))) && directionextY != 0:
		valid := AttemptMoveByPushingWiderBoxes(maze, nextX, nextY, directionextX, directionextY, attemptMoveOnly)
		if valid && !attemptMoveOnly {
			maze[nextX][nextY] = maze[x][y]
			maze[x][y] = byte(TerrainType(Free))
		}
		return valid
	case maze[nextX][nextY] == byte(TerrainType(BoxL)) && directionextX != 0:
		valid := AttemptMoveByPushingWiderBoxes(maze, nextX, nextY, directionextX, directionextY, true) && AttemptMoveByPushingWiderBoxes(maze, nextX, nextY+1, directionextX, directionextY, true)
		if valid && !attemptMoveOnly {
			AttemptMoveByPushingWiderBoxes(maze, nextX, nextY, directionextX, directionextY, false)
			AttemptMoveByPushingWiderBoxes(maze, nextX, nextY+1, directionextX, directionextY, false)
			maze[nextX][nextY] = maze[x][y]
			maze[x][y] = byte(TerrainType(Free))
		}
		return valid
	case maze[nextX][nextY] == byte(TerrainType(BoxR)) && directionextX != 0:
		valid := AttemptMoveByPushingWiderBoxes(maze, nextX, nextY, directionextX, directionextY, true) && AttemptMoveByPushingWiderBoxes(maze, nextX, nextY-1, directionextX, directionextY, true)
		if valid && !attemptMoveOnly {
			AttemptMoveByPushingWiderBoxes(maze, nextX, nextY, directionextX, directionextY, false)
			AttemptMoveByPushingWiderBoxes(maze, nextX, nextY-1, directionextX, directionextY, false)
			maze[nextX][nextY] = maze[x][y]
			maze[x][y] = byte(TerrainType(Free))
		}
		return valid

	case maze[nextX][nextY] == byte(TerrainType(Wall)):
		return false
	case maze[nextX][nextY] == byte(TerrainType(Free)):
		if !attemptMoveOnly {
			maze[nextX][nextY] = maze[x][y]
			maze[x][y] = byte(TerrainType(Free))
		}
		return true
	default:
		return false
	}
}

func main() {
	fmt.Println("Day 15 of Advent Of Code 2024, j4nbob")
	raw := LoadTextFromFile("day15.input")
	mazeRawX := strings.Split(raw, "\n\n")
	mazeRaw := mazeRawX[0]
	movesRaw := mazeRawX[1]
	robotX, robotY := 0, 0
	maze, robotX, robotY = CreateMaze(mazeRaw)
	robotPos = Coordinate{X: robotX, Y: robotY}
	robotMoves := CreateMoves(movesRaw)
	fmt.Println("Part 1 Input Maze : ")
	visualiseMaze(maze)
	for _, move := range robotMoves {
		moveFlag := false
		switch MoveType(move) {
		case Up:
			if AttemptMoveByPushingBoxes(&maze, Up, robotPos, Coordinate{X: robotPos.X, Y: robotPos.Y - 1}) {
				moveFlag = true
			}
			if moveFlag {
				maze[robotPos.Y][robotPos.X] = byte(TerrainType(Free))
				robotPos.Y--
				maze[robotPos.Y][robotPos.X] = byte(TerrainType(Prankster))
			}
		case Down:
			if AttemptMoveByPushingBoxes(&maze, Down, robotPos, Coordinate{X: robotPos.X, Y: robotPos.Y + 1}) {
				moveFlag = true
			}
			if moveFlag {
				maze[robotPos.Y][robotPos.X] = byte(TerrainType(Free))
				robotPos.Y++
				maze[robotPos.Y][robotPos.X] = byte(TerrainType(Prankster))
			}
		case Right:
			if AttemptMoveByPushingBoxes(&maze, Right, robotPos, Coordinate{X: robotPos.X + 1, Y: robotPos.Y}) {
				moveFlag = true
			}
			if moveFlag {
				maze[robotPos.Y][robotPos.X] = byte(TerrainType(Free))
				robotPos.X++
				maze[robotPos.Y][robotPos.X] = byte(TerrainType(Prankster))
			}
		case Left:
			if AttemptMoveByPushingBoxes(&maze, Left, robotPos, Coordinate{X: robotPos.X - 1, Y: robotPos.Y}) {
				moveFlag = true
			}
			if moveFlag {
				maze[robotPos.Y][robotPos.X] = byte(TerrainType(Free))
				robotPos.X--
				maze[robotPos.Y][robotPos.X] = byte(TerrainType(Prankster))
			}

		default:
			fmt.Println("Invalid move proposed. Panic, panic, panic!")
			os.Exit(1)
		}
		//fmt.Println("moveIndex=", moveIndex, "Moved=", moveFlag, "robotPos=", robotPos)
	}

	fmt.Println("Part 1 Output Maze : ")
	visualiseMaze(maze)
	daSum := getCoordinatesOfBoxesGPSSum(&maze, byte(TerrainType(Box)))
	fmt.Println("part1/daSum=", daSum)

	daSum = 0
	maze, _, _ = CreateMaze(mazeRaw)
	maze = widenMaze(maze)
	fmt.Println("Part 2 Input Maze : ")
	visualiseMaze(maze)

	var position []int
	for i, row := range maze {
		for j, col := range row {
			if col == '@' {
				position = []int{i, j}
				break
			}
			if len(position) == 2 {
				break
			}
		}
	}
	for _, move := range robotMoves {
		switch MoveType(move) {
		case Up:
			if AttemptMoveByPushingWiderBoxes(maze, position[0], position[1], -1, 0, false) {
				position[0]--
			}
		case Down:
			if AttemptMoveByPushingWiderBoxes(maze, position[0], position[1], 1, 0, false) {
				position[0]++
			}
		case Right:
			if AttemptMoveByPushingWiderBoxes(maze, position[0], position[1], 0, 1, false) {
				position[1]++
			}
		case Left:
			if AttemptMoveByPushingWiderBoxes(maze, position[0], position[1], 0, -1, false) {
				position[1]--
			}
		default:
			fmt.Println("Invalid move proposed. Panic, panic, panic!")
			os.Exit(1)
		}
		//fmt.Println("moveIndex=", moveIndex, "robotPos=", position)

	}
	fmt.Println("Part 2 Output Maze : ")
	visualiseMaze(maze)

	daSum = getCoordinatesOfBoxesGPSSum(&maze, byte(TerrainType(BoxL)))
	fmt.Println("part2/daSum=", daSum)
}

func visualiseMaze(maze [][]byte) {
	for _, row := range maze {
		for _, col := range row {
			if col == byte(TerrainType(Free)) {
				fmt.Print(".")
			} else if col == byte(TerrainType(Wall)) {
				fmt.Print("#")
			} else if col == byte(TerrainType(Box)) {
				fmt.Print("O")
			} else if col == byte(TerrainType(Prankster)) {
				fmt.Print("@")
			} else if col == byte(TerrainType(BoxL)) {
				fmt.Print("[")
			} else if col == byte(TerrainType(BoxR)) {
				fmt.Print("]")
			}
		}
		fmt.Println()
	}
}
