package day4

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GetDirections(grid []string, row, column int) []string {
	var directions []string
	for horizontalPosition := -1; horizontalPosition <= 1; horizontalPosition++ {
		for verticalPosition := -1; verticalPosition <= 1; verticalPosition++ {
			if horizontalPosition == 0 && verticalPosition == 0 {
				continue
			}
			if row+3*horizontalPosition >= 0 && row+3*horizontalPosition <
				len(grid) && column+3*verticalPosition >= 0 && column+3*verticalPosition < len(grid[row]) {

				var direction string
				valid := true
				for index := 0; index < 4; index++ {
					newHorizontalPosition := row + index*horizontalPosition
					newVerticalPosition := column + index*verticalPosition

					if newHorizontalPosition < 0 || newHorizontalPosition >= len(grid) ||
						newVerticalPosition < 0 || newVerticalPosition >= len(grid[newHorizontalPosition]) {
						valid = false
						break
					}

					direction += string(grid[newHorizontalPosition][newVerticalPosition])
				}
				if valid {
					directions = append(directions, direction)
				}
			}
		}
	}
	return directions
}

func IsXmas(grid []string, horizontal, vertical int) bool {
	if horizontal-1 >= 0 && horizontal+1 < len(grid) && vertical-1 >= 0 && vertical+1 < len(grid[horizontal]) {
		pattern1 := string(grid[horizontal-1][vertical-1]) + string(grid[horizontal+1][vertical+1])
		pattern2 := string(grid[horizontal-1][vertical+1]) + string(grid[horizontal+1][vertical-1])
		if (pattern1 == "MS" || pattern1 == "SM") && (pattern2 == "MS" || pattern2 == "SM") {
			return true
		}
	}
	return false
}

func RunDay4() {
	input, err := ioutil.ReadFile("days/day4/input.txt")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	grid := strings.Split(string(input), "\n")

	partA := 0
	for horizontalElement, row := range grid {
		for verticalElement := range row {
			for _, direction := range GetDirections(grid, horizontalElement, verticalElement) {
				if direction == "XMAS" {
					partA++
				}
			}
		}
	}
	fmt.Println("Part A:", partA)

	partB := 0
	for horizontalElement, row := range grid {
		for verticalElement, direction := range row {
			if direction == 'A' && IsXmas(grid, horizontalElement, verticalElement) {
				partB++
			}
		}
	}
	fmt.Println("Part B:", partB)
}
