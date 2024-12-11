package day5

import (
	"fmt"
	"io/ioutil"
	"sort"
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

func (instance *Sortable) LessThan(other *Sortable, rules map[int]struct{}) bool {
	_, exists := rules[other.Index]
	return !exists
}

func DataParsing(input string) (map[int]struct{}, string) {
	parts := strings.Split(input, "\n\n")

	rulesData := parts[0]
	updates := parts[1]

	rules := make(map[int]struct{})
	for _, line := range strings.Split(rulesData, "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "|")
		key, _ := strconv.Atoi(parts[0])
		rules[key] = struct{}{}
	}

	return rules, updates
}

func splitLines(input string) []string {
	return strings.Split(input, "\n")
}

func SortableFromUpdatesList(updatesList string, rules map[int]struct{}) (int, int) {
	lines := splitLines(updatesList)
	partA, partB := 0, 0

	for _, line := range lines {
		parts := strings.Split(line, ",")
		var unsortedUpdate []*Sortable
		for _, part := range parts {
			index, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println(err)
				return 0, 0
			}
			unsortedUpdate = append(unsortedUpdate, NewSortable(index))
		}

		sortedUpdate := make([]*Sortable, len(unsortedUpdate))
		copy(sortedUpdate, unsortedUpdate)
		sort.Slice(sortedUpdate, func(i, j int) bool {
			return sortedUpdate[i].LessThan(sortedUpdate[j], rules)
		})

		if isEqual(unsortedUpdate, sortedUpdate) {
			partA += sortedUpdate[len(sortedUpdate)/2].Index
		} else {
			partB += sortedUpdate[len(sortedUpdate)/2].Index
		}
	}

	return partA, partB
}

func isEqual(a, b []*Sortable) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}
	return true
}

func RunDay5() {
	input, err := ioutil.ReadFile("days/day5/input.txt")
	if err != nil {
		panic(err)
	}

	inputString := string(input)
	rules, updates := DataParsing(inputString)

	partA, partB := SortableFromUpdatesList(updates, rules)
	fmt.Println(partA, partB)
}
