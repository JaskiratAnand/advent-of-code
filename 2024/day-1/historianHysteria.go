package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Part 1
	data, err := os.ReadFile("input.txt")
	check(err)

	var aList []int
	var bList []int

	// reading input file and storing values in aList and bList
	for _, line := range strings.Split(string(data), "\n") {
		parts := strings.Fields(line)

		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])

		aList = append(aList, num1)
		bList = append(bList, num2)
	}

	// sorting lists
	slices.Sort(aList)
	slices.Sort(bList)

	// calculating distance between each element of aList and bList
	var distanceBetween []int
	for i := 0; i < len(aList); i++ {
		distanceBetween = append(distanceBetween, abs(aList[i]-bList[i]))
	}

	// calculating total distance
	var totalDistance int = 0
	for _, distaceBetween := range distanceBetween {
		totalDistance += distaceBetween
	}

	fmt.Printf("Part 1 \nTotal Distance: %d", totalDistance)

	// Part 2

	var similarityScore int = 0

	for i := 0; i < len(aList); i++ {
		var numberRepeats int = 0

		// checking for repeated numbers in bList
		for j := 0; j < len(bList); j++ {
			if aList[i] == bList[j] {
				numberRepeats++
			}
		}

		// calculating similarity score => aList[i] * number of repeats
		similarityScore += aList[i] * numberRepeats
	}

	fmt.Printf("\nPart 2 \nSimilarity Score: %d", similarityScore)
}
