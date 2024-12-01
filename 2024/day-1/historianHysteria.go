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
	var distaceBetween []int
	for i := 0; i < len(aList); i++ {
		distance := aList[i] - bList[i]
		if distance < 0 {
			distance = distance * -1
		}
		distaceBetween = append(distaceBetween, distance)
	}

	// calculating total distance
	var totalDistance int = 0
	for _, distaceBetween := range distaceBetween {
		totalDistance += distaceBetween
	}

	fmt.Printf("Part 1 \nTotal Distance: %d", totalDistance)

	// Part 2

	var similarityScore int = 0

	for i := 0; i < len(aList); i++ {
		var numberRepeats int = 0

		for j := 0; j < len(bList); j++ {
			if aList[i] == bList[j] {
				numberRepeats++
			}
		}

		similarityScore += aList[i] * numberRepeats
	}

	fmt.Printf("\nPart 2 \nSimilarity Score: %d", similarityScore)
}