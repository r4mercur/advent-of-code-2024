package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var leftList []int
var rightList []int

func Total(listA []int, listB []int) int {
	total := 0
	for i := 0; i < len(listA) && i < len(listB); i++ {
		total += int(math.Abs(float64(listA[i] - listB[i])))
	}
	return total
}

func countOccurrences(list []int) map[int]int {
	counts := make(map[int]int)
	for _, value := range list {
		counts[value]++
	}
	return counts
}

func similarityScore(listA []int, listB []int) int {
	score := 0
	rightCount := countOccurrences(listB)

	for _, value := range listA {
		countInB, ok := rightCount[value]
		if ok {
			score += value * countInB
		}
	}

	return score
}

func RunDay1() {
	fmt.Println("Advent of Code 2024 - Day 1")

	file, err := os.Open("days/day1/input.txt")
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
		if len(parts) == 2 {
			left, err1 := strconv.Atoi(parts[0])
			right, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				leftList = append(leftList, left)
				rightList = append(rightList, right)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	result := Total(leftList, rightList)
	fmt.Println("Result: ", result)

	score := similarityScore(leftList, rightList)
	fmt.Println("Similarity score: ", score)
}
