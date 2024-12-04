package day3

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func RunDay3() {
	input, err := ioutil.ReadFile("days/day3/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	inputStr := string(input)

	part1, part2 := 0, 0
	enabled := true

	re := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)
	matches := re.FindAllString(inputStr, -1)

	for _, s := range matches {
		if strings.HasPrefix(s, "d") {
			enabled = s == "do()"
		} else {
			nums := strings.Split(s[4:len(s)-1], ",")
			product1, _ := strconv.Atoi(nums[0])
			product2, _ := strconv.Atoi(nums[1])
			part1 += product1 * product2
			if enabled {
				part2 += product1 * product2
			}
		}
	}

	fmt.Printf("part1: %d, part2: %d\n", part1, part2)
}
