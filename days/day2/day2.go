package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var counterOfSaveLinesPartOne int = 0
var counterOfSaveLinesPartTwo int = 0

func isSafe(list []int) bool {
	increasing := true
	decreasing := true

	for index := 1; index < len(list); index++ {
		difference := list[index] - list[index-1]

		if difference < -3 || difference > 3 || difference == 0 {
			return false
		}

		if difference > 0 {
			decreasing = false
		}
		if difference < 0 {
			increasing = false
		}
	}
	return increasing || decreasing
}

func makeLineSafeWithOneChange(list []int) bool {
	for index := 0; index < len(list); index++ {
		modList := append([]int{}, list[:index]...)
		modList = append(modList, list[index+1:]...)

		if isSafe(modList) {
			return true
		}
	}
	return false
}

func RunDay2() {
	fmt.Println("Advent of Code 2024 - Day 2")

	file, err := os.Open("days/day2/input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		var ListOfLines []int

		for _, part := range parts {
			level, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting to int")
				return
			}
			ListOfLines = append(ListOfLines, level)
		}

		if isSafe(ListOfLines) {
			counterOfSaveLinesPartOne += 1
		}

		if isSafe(ListOfLines) || makeLineSafeWithOneChange(ListOfLines) {
			counterOfSaveLinesPartTwo += 1
		}
	}

	fmt.Println("Safe lines: ", counterOfSaveLinesPartOne)
	fmt.Println("Safe lines with one change: ", counterOfSaveLinesPartTwo)
}
