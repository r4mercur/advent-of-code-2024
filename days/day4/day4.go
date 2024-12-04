package day4

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Direction int
type Position struct {
	row    int
	column int
}

const (
	Up Direction = iota
	Down
	Left
	Right
	DiagonalUpRight
	DiagonalUpLeft
	DiagonalDownRight
	DiagonalDownLeft
)

func (d Direction) AllDirections() []Direction {
	return []Direction{Up, Down, Left, Right, DiagonalDownLeft, DiagonalDownRight, DiagonalUpRight, DiagonalUpLeft}
}

func (d Direction) DiagonalDirections() []Direction {
	return []Direction{DiagonalDownLeft, DiagonalDownRight, DiagonalUpRight, DiagonalUpLeft}
}

func NewPosition(row, col int) Position {
	return Position{row: row, column: col}
}

func (p Position) Get(matrix [][]rune) (rune, error) {
	if p.row >= len(matrix) || p.row < 0 || p.column < 0 || p.column >= len(matrix[p.row]) {
		return 0, errors.New("position out of bounds")
	}
	return matrix[p.row][p.column], nil
}

func (p Position) Translate(matrix [][]rune, direction Direction) (Position, error) {
	switch direction {
	case Up:
		if p.row-1 < 0 {
			return Position{}, errors.New("underflow")
		}
		return Position{row: p.row - 1, column: p.column}, nil
	case Down:
		if p.row+1 >= len(matrix) {
			return Position{}, errors.New("overflow")
		}
		return Position{row: p.row + 1, column: p.column}, nil
	case Left:
		if p.column-1 < 0 {
			return Position{}, errors.New("underflow")
		}
		return Position{row: p.row, column: p.column - 1}, nil
	case Right:
		if p.column+1 >= len(matrix[0]) {
			return Position{}, errors.New("overflow")
		}
		return Position{row: p.row, column: p.column + 1}, nil
	case DiagonalUpLeft:
		newPos, err := p.Translate(matrix, Up)
		if err != nil {
			return Position{}, err
		}
		return newPos.Translate(matrix, Left)
	case DiagonalUpRight:
		newPos, err := p.Translate(matrix, Up)
		if err != nil {
			return Position{}, err
		}
		return newPos.Translate(matrix, Right)
	case DiagonalDownLeft:
		newPos, err := p.Translate(matrix, Down)
		if err != nil {
			return Position{}, err
		}
		return newPos.Translate(matrix, Left)
	case DiagonalDownRight:
		newPos, err := p.Translate(matrix, Down)
		if err != nil {
			return Position{}, err
		}
		return newPos.Translate(matrix, Right)
	default:
		return Position{}, errors.New("invalid direction")
	}
}

func (p Position) Midpoint(other Position) Position {
	newRow := (p.row + other.row) / 2
	newCol := (p.column + other.column) / 2
	return Position{row: newRow, column: newCol}
}

func CreateMatrix(input string) [][]rune {
	lines := strings.Split(input, "\r\n")
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix
}

func CheckMatrix(matrix [][]rune, targets []rune, position Position, direction Direction) (bool, error) {
	currentPos := position
	for _, target := range targets {
		char, err := currentPos.Get(matrix)
		if err != nil || char != target {
			return false, nil
		}
		currentPos, err = currentPos.Translate(matrix, direction)
		if err != nil {
			return false, nil
		}
	}
	return true, nil
}

func PartOne(input string) int {
	matrix := CreateMatrix(input)
	directions := Direction(0).AllDirections()
	targets := []rune{'X', 'M', 'A', 'S'}
	total := 0

	for rowIdx := 0; rowIdx < len(matrix); rowIdx++ {
		for colIdx := 0; colIdx < len(matrix[rowIdx]); colIdx++ {
			position := NewPosition(rowIdx, colIdx)
			for _, direction := range directions {
				found, err := CheckMatrix(matrix, targets, position, direction)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				if found {
					total++
				}
			}
		}
	}
	return total
}

func RunDay4() {
	input, err := ioutil.ReadFile("days/day4/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	inputStr := string(input)

	fmt.Println("Advent of Code 2024 - Day 4")
	fmt.Println("Part 1:", PartOne(inputStr))
}
