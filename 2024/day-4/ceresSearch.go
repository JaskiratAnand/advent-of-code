package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var searchDirections = [8][2]int{
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 0},   // Down
	{-1, 0},  // Up
	{1, 1},   // Diagonal Down-Right
	{-1, -1}, // Diagonal Up-Left
	{1, -1},  // Diagonal Down-Left
	{-1, 1},  // Diagonal Up-Right
}

func countXmasOccurances(runeGrid [][]rune, word string) int {
	wordLength := len(word)
	rows := len(runeGrid)
	cols := len(runeGrid[0])
	count := 0

	isValidXmas := func(row, col, direction int) bool {
		for i := 0; i < wordLength; i++ {
			newRow := row + i*searchDirections[direction][0]
			newCol := col + i*searchDirections[direction][1]

			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				return false
			}

			if runeGrid[newRow][newCol] != rune(word[i]) {
				return false
			}
		}
		return true
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for dir := 0; dir < 8; dir++ {
				if isValidXmas(row, col, dir) {
					count++
				}
			}
		}
	}

	return count
}

func countCrossMasOccurances(grid [][]rune) int {

	runeGrid := make([][]rune, len(grid))
	for i := range grid {
		runeGrid[i] = []rune(grid[i])
	}

	rows, cols := len(runeGrid), len(runeGrid[0])
	count := 0

	isDiagonalMatch := func(row, col, dRow, dCol int) bool {
		target := "MAS"
		targetAlt := "SAM"
		rows, cols := len(runeGrid), len(runeGrid[0])
		chars := ""

		for i := 0; i < len(target); i++ {
			newRow, newCol := row+dRow*i, col+dCol*i
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				return false
			}
			chars += string(runeGrid[newRow][newCol])
		}
		return chars == target || chars == targetAlt
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Top-left to bottom-right && Top-right to bottom-left
			if isDiagonalMatch(row-1, col-1, 1, 1) && isDiagonalMatch(row-1, col+1, 1, -1) {
				count++
			}
		}
	}

	return count
}

func main() {
	// Part 1
	// ..X...
	// .SAMX.
	// .A..A.
	// XMAS.S
	// .X....

	data, err := os.ReadFile("input.txt")
	check(err)

	word := "XMAS"

	lines := strings.Split(string(data), "\n")

	runeGrid := make([][]rune, len(lines))
	for i, line := range lines {
		line := strings.TrimSpace(line)
		runeGrid[i] = make([]rune, len(line))
		for j, r := range line {
			runeGrid[i][j] = r
		}
	}

	totalXmasCount := countXmasOccurances(runeGrid, word)

	fmt.Printf("Part 1 \nTotal XMAS count: %v \n", totalXmasCount)

	// Part 2
	// M.S
	// .A.
	// M.S

	totalCrossMasCount := countCrossMasOccurances(runeGrid)
	fmt.Printf("Part 2 \nTotal Cross-MAS count: %v \n", totalCrossMasCount)
}
