// day13 of Advent Of Code 2024 : https://adventofcode.com/2024/day/13
// j4nbob

package main

import (
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coordinate struct {
	X, Y       int
	Jump1Count int
	Jump2Count int
}

type Task struct {
	aC Coordinate
	bC Coordinate
	pC Coordinate
}

var taskDefinitions []Task

func LoadTextFromFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}
	return string(data)
}
func ExtractTokens(input string) map[string]string {
	lines := strings.Split(input, "\n")
	tokens := make(map[string]string)
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			tokens[key] = value
		}
	}
	return tokens
}
func ExtractNumericalTokens(input string, prize bool) map[string]int {
	var re *regexp.Regexp
	if prize {
		re = regexp.MustCompile(`([XY])\=(\d+)`)
	} else {
		re = regexp.MustCompile(`([XY])\+(\d+)`)
	}
	matches := re.FindAllStringSubmatch(input, -1)
	tokens := make(map[string]int)
	for _, match := range matches {
		if len(match) == 3 {
			key := match[1]
			value, _ := strconv.Atoi(match[2])
			tokens[key] = value
		}
	}
	return tokens
}

func calculateLeastSteps(start, end, jump1, jump2 Coordinate) (int, int, int) {
	queue := list.New()
	queue.PushBack(start)
	visited := make(map[Coordinate]bool)
	visited[start] = true
	steps := 0
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			current := queue.Remove(queue.Front()).(Coordinate)
			if current.X == end.X && current.Y == end.Y {
				return steps, current.Jump1Count, current.Jump2Count
			}
			// MAC DADDY WILL MAKE YOU JUMP JUMP
			next1 := Coordinate{X: current.X + jump1.X, Y: current.Y + jump1.Y, Jump1Count: current.Jump1Count + 1, Jump2Count: current.Jump2Count}
			next2 := Coordinate{X: current.X + jump2.X, Y: current.Y + jump2.Y, Jump1Count: current.Jump1Count, Jump2Count: current.Jump2Count + 1}
			if !visited[next1] && next1.X <= end.X && next1.Y <= end.Y {
				queue.PushBack(next1)
				visited[next1] = true
			}
			if !visited[next2] && next2.X <= end.X && next2.Y <= end.Y {
				queue.PushBack(next2)
				visited[next2] = true
			}
		}
		steps++
	}
	return -1, -1, -1
}

func ProcessTask(task Task, costA int, costB int, part2 bool) int {
	daBeginning := Coordinate{X: 0, Y: 0, Jump1Count: 0, Jump2Count: 0}
	steps, jump1Count, jump2Count := calculateLeastSteps(daBeginning, task.pC, task.aC, task.bC)
	if steps == -1 {
		return 0
	} else {
		if !part2 {
			tokenCost := costA*jump1Count + costB*jump2Count
			return tokenCost
		} else {
			//will do this differently for part 2
			return 0
		}
	}

}

func main() {
	fmt.Println("Day 13 of Advent Of Code 2024, j4nbob")
	raw := LoadTextFromFile("day13.input")
	arr := strings.Split(raw, "\n\n")
	for _, val := range arr {
		tokens := ExtractTokens(val)
		a := ExtractNumericalTokens(tokens["Button A"], false)
		b := ExtractNumericalTokens(tokens["Button B"], false)
		p := ExtractNumericalTokens(tokens["Prize"], true)
		taskDefinitions = append(taskDefinitions, Task{Coordinate{a["X"], a["Y"], 0, 0}, Coordinate{b["X"], b["Y"], 0, 0}, Coordinate{p["X"], p["Y"], 0, 0}})

	}
	var daSum int = 0
	for _, task := range taskDefinitions {
		daSum += ProcessTask(task, 3, 1, false)
	}
	fmt.Println("part1/daSum=", daSum)

	//for part 2 we will use equation solver
	daSum = 0
	for _, task := range taskDefinitions {
		taskCost := 0
		task.pC.X = 10000000000000 + task.pC.X
		task.pC.Y = 10000000000000 + task.pC.Y
		kb := (task.pC.X*task.aC.Y - task.pC.Y*task.aC.X) / (task.bC.X*task.aC.Y - task.aC.X*task.bC.Y)
		ka := (task.pC.X - kb*task.bC.X) / task.aC.X
		if ka*task.aC.X+kb*task.bC.X != task.pC.X {
			taskCost = 0
			continue
		}
		if ka*task.aC.Y+kb*task.bC.Y != task.pC.Y {
			taskCost = 0
			continue
		}
		taskCost = ka*3 + kb
		daSum += taskCost
	}
	fmt.Println("part2/daSum=", daSum)
}
