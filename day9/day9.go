// day9 of Advent Of Code 2024 : https://adventofcode.com/2024/day/9
// j4nbob
package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type blockInfoType struct {
	id              int
	numbersLeft     int
	originalContent string
	emptySpacesLeft int
	defragCharacter string
	swapDepleted    bool
}

func LoadTextFromFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}
	return string(data)
}

func RepeatIDBuilder(index int, count int) string {
	if count < 1 {
		return ""
	}
	chars := strconv.Itoa(index)
	return strings.Repeat(chars, count)
}

func loadStringintoBlockInfoType(input string) []blockInfoType {
	var blockInfos []blockInfoType
	var index int = 0
	var digitNumber int
	var emptyNumber int

	for i := 0; i < len(input); i += 2 {
		if i+1 < len(input) {
			digitNumber = int(input[i] - '0')
			emptyNumber = int(input[i+1] - '0')
		} else {
			digitNumber = int(input[i] - '0')
			emptyNumber = 0
		}
		//fmt.Println(index, " pair: ", digitNumber, ":", emptyNumber)
		expandContent := RepeatIDBuilder(index, digitNumber)
		blockInfos = append(blockInfos, blockInfoType{id: index, numbersLeft: digitNumber, emptySpacesLeft: emptyNumber, originalContent: expandContent, defragCharacter: "", swapDepleted: false})
		index++

	}
	return blockInfos

}

func ChopOffRight(input string, x int) string {
	if len(input) >= x {
		return input[:len(input)-x]
	}
	return ""
}

func FindIndexFromLeftWithEmptySpacesOfLen(blockInfos *[]blockInfoType, length int) int {
	for i := 0; i < len(*blockInfos); i++ {
		if !(*blockInfos)[i].swapDepleted {
			if (*blockInfos)[i].emptySpacesLeft >= length {
				return i
			}
		}
	}
	return -1
}

func SwapLeftRight(blockInfos *[]blockInfoType, leftIndex int, rightIndex int, rightId int) bool {

	chars := strconv.Itoa(rightId)
	(*blockInfos)[leftIndex].emptySpacesLeft -= len(chars)
	(*blockInfos)[leftIndex].defragCharacter = (*blockInfos)[leftIndex].defragCharacter + chars

	(*blockInfos)[rightIndex].numbersLeft--
	(*blockInfos)[rightIndex].originalContent = ChopOffRight((*blockInfos)[rightIndex].originalContent, len(chars))
	if (*blockInfos)[rightIndex].numbersLeft == 0 {
		(*blockInfos)[rightIndex].swapDepleted = true
	}
	return true

}

// defragmentation
func deFragmentBlocks(blockInfos *[]blockInfoType) bool {
	l := len(*blockInfos)
	for right := l - 1; right >= 0; right-- {
		for {
			rightBlock := (*blockInfos)[right]
			if rightBlock.numbersLeft > 0 {
				chars := strconv.Itoa(rightBlock.id)
				foundIndex := FindIndexFromLeftWithEmptySpacesOfLen(blockInfos, len(chars))
				if foundIndex == -1 {
					//does not fit anywhere
					return true
				}
				if !SwapLeftRight(blockInfos, foundIndex, right, rightBlock.id) {
					return false
				}
			} else {
				break
			}

		}

	}

	return true
}

// takes array of blockInfoType and returns number of iterations
func blockInfosToString(blockInfos []blockInfoType) string {
	var result string
	var l = len(blockInfos)
	for i := 0; i < l; i++ {
		result += blockInfos[i].originalContent + blockInfos[i].defragCharacter
		if blockInfos[i].emptySpacesLeft > 0 {
			sub := strings.Repeat(".", blockInfos[i].emptySpacesLeft)
			result = result + sub
		}
	}
	return result
}

func deFragmentFiles(blockInfos *[]blockInfoType) bool {
	l := len(*blockInfos)
	for right := l - 1; right >= 0; right-- {
		for {
			rightBlock := (*blockInfos)[right]
			if rightBlock.numbersLeft > 0 {
				chars := strconv.Itoa(rightBlock.id)
				foundIndex := FindIndexFromLeftWithEmptySpacesOfLen(blockInfos, len(chars))
				if foundIndex == -1 {
					return true
				}
				if !SwapLeftRight(blockInfos, foundIndex, right, rightBlock.id) {
					return false
				}
			} else {
				break
			}

		}

	}

	return true
}

/*
looks to be incorrect way to do this
func deFragmentString(fragmentedString string) string {
	runes := []rune(fragmentedString)
	n := len(runes)
	result := make([]rune, n)
	copy(result, runes)
	lastSwap := math.MaxInt
	for i := n - 1; i >= 0; i-- {
		if runes[i] != '.' {
			boundaryMax := min(lastSwap, n)
			for j := 0; j < boundaryMax; j++ {
				if result[j] == '.' {
					result[j] = runes[i]
					result[i] = '.'
					lastSwap = i
					break
				}
			}
		}
	}

	return string(result)
}
*/

func SimpleCheckSum(defragString string) *big.Int {
	sum := big.NewInt(0)
	for index, char := range defragString {
		if char != '.' {
			digit := int(char - '0')
			adder := index * digit
			value := big.NewInt(int64(adder))
			sum.Add(sum, value)
		}
	}
	return sum
}

func main() {
	fmt.Println("Day 9 of Advent Of Code 2024, j4nbob")
	var contents = LoadTextFromFile("day9.input")
	var blockInfos = loadStringintoBlockInfoType(contents)
	if !deFragmentBlocks(&blockInfos) {
		os.Exit(3)
	}
	var defragmentedString = blockInfosToString(blockInfos)
	//fmt.Println(defragmentedString)
	fmt.Println("part1/daSum=", SimpleCheckSum(defragmentedString))
	var blockInfos2 = loadStringintoBlockInfoType(contents)
	if !deFragmentFiles(&blockInfos2) {
		os.Exit(4)
	}
	defragmentedString = blockInfosToString(blockInfos2)
	fmt.Println("part2/daSum=", SimpleCheckSum(defragmentedString))

}
