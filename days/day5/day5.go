package day5

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Sortable struct {
	Index int
}

func NewSortable(index int) *Sortable {
	return &Sortable{Index: index}
}

func (instance *Sortable) Equals(other *Sortable) bool {
	return instance.Index == other.Index
}

func (instance *Sortable) LessThan(other *Sortable, rules map[int]int) bool {
	_, exists := rules[instance.Index]
	return !exists
}

func DataParsing(input string) (map[int]int, string) {
	parts := strings.Split(input, "\n\n")

	rulesData := parts[0]
	updates := parts[1]

	rules := make(map[int]int)
	for _, line := range strings.Split(rulesData, "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		key, _ := strconv.Atoi(parts[0])
		value, _ := strconv.Atoi(parts[1])
		rules[key] = value
	}

	return rules, updates
}

func splitLines(input string) []string {
	return strings.Split(input, "\n")
}

func SortableFromUpdatesList(updatesList string) {
	lines := splitLines(updatesList)
	for _, line := range lines {
		parts := strings.Split(line, ", ")
		var unsortedUpdate []*Sortable
		for _, part := range parts {
			index, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println(err)
				return
			}
			unsortedUpdate = append(unsortedUpdate, NewSortable(index))
		}

		// TODO: sort here the unsortedUpdate and check cases (length / 2) & also return value
	}
}

func RunDay5() {
	input, err := ioutil.ReadFile("days/day5/input.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(input)
	rules, updates := DataParsing(inputString)

	fmt.Println(rules)
	fmt.Println(updates)

	SortableFromUpdatesList(updates)
}
