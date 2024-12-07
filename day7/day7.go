// day7 of Advent Of Code 2024 : https://adventofcode.com/2024/day/7
// j4nbob
package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func generateCombinations(chars []rune, length int) []string {
	if length == 0 {
		return []string{""}
	}
	smallerCombinations := generateCombinations(chars, length-1)
	var combinations []string
	for _, combination := range smallerCombinations {
		for _, char := range chars {
			combinations = append(combinations, combination+string(char))
		}
	}
	return combinations
}

func main() {
	fmt.Println("Day 7 of Advent Of Code 2024, j4nbob")

	file, err := os.Open("day7.input")
	if err != nil {
		os.Exit(1)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			os.Exit(2)
			continue
		}
		key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			os.Exit(3)
			continue
		}
		valuesStr := strings.Fields(parts[1])
		var values []int
		for _, valueStr := range valuesStr {
			value, err := strconv.Atoi(valueStr)
			if err != nil {
				os.Exit(4)
				continue
			}
			values = append(values, value)
		}
		data[key] = values
	}

	if err := scanner.Err(); err != nil {
		os.Exit(5)
		return
	}

	if true {
		addSolutions := big.NewInt(0)
		counter := 0
		//fmt.Println("total work=", len(data))
		for key, values := range data {
			//fmt.Println("key=", key, "values=", values)
			elements := len(values) - 1
			chars := []rune{'+', '*'}
			combined := generateCombinations(chars, elements)
			counter++
			//fmt.Println("counter=", counter, "key=", key, "values=", values, "uniqueValues=", len(combined))

			for _, uniqueKey := range combined {
				runes := []rune(uniqueKey)
				daSum := 0

				for i := 0; i <= elements; i++ {
					if i == 0 {
						daSum = values[i]
					} else {
						var operator rune = runes[i-1]
						switch operator {
						case '+':
							daSum = daSum + values[i]
						case '*':
							daSum = daSum * values[i]
						}
					}

				}
				if daSum == key {
					addSolutions = new(big.Int).Add(addSolutions, big.NewInt(int64(daSum)))
					//addSolutions += daSum
					//fmt.Println("Found Solution daSum=,", daSum, ",adder=", addSolutions, "attempt=", i)
					break
				}

			}

		}
		fmt.Println("part1/daSum=", addSolutions)
	}

	//part2
	if true {
		addSolutions := big.NewInt(0)
		counter := 0
		//fmt.Println("total work=", len(data))
		for key, values := range data {
			//fmt.Println("key=", key, "values=", values)
			elements := len(values) - 1
			chars := []rune{'+', '*', '|'}
			combined := generateCombinations(chars, elements)
			counter++
			//fmt.Println("counter=", counter, "key=", key, "values=", values, "uniqueValues=", len(combined))

			for _, uniqueKey := range combined {
				runes := []rune(uniqueKey)
				daSum := 0
				daSaved := 0
				daSavePlusOperation := false
				for i := 0; i <= elements; i++ {
					if i == 0 {
						daSum = values[i]
					} else {
						var operator rune = runes[i-1]
						switch operator {
						case '+':
							daSaved = daSum
							daSavePlusOperation = true
							daSum = daSum + values[i]
						case '*':
							daSaved = daSum
							daSavePlusOperation = false
							daSum = daSum * values[i]
						case '|':

							daSum = daSaved
							var prevNum = values[i-1]
							var curNum = values[i]
							var concatString = fmt.Sprintf("%d%d", prevNum, curNum)
							concatNum, err := strconv.Atoi(concatString)
							if err != nil {
								os.Exit(6)
							}
							//fmt.Println("concatString=", concatString, "concatNum=", concatNum)
							if daSum == 0 {
								daSum = concatNum
							} else {
								if daSavePlusOperation {
									daSum = daSum + concatNum
								} else {
									daSum = daSum * concatNum
								}
							}
							//fmt.Println("concatString=", concatString, "concatNum=", concatNum, "daSum=", daSum)
							daSaved = daSum
						}
					}

				}
				if daSum == key {
					addSolutions = new(big.Int).Add(addSolutions, big.NewInt(int64(daSum)))
					// 7290: 6 8 6 15 can be made true using 6 * 8 || 6 * 15
					// ^^^to me this is wrong it should be^^^^
					// 7740: 6 8 6 15 can be made true using  6 * 8 || 6 * 15 , 6*86*15=7740
					break
				}

			}

		}
		fmt.Println("part2/daSum=", addSolutions)

	}

}
