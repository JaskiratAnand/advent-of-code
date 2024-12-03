package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getMul(inputStr string) int {
	// returns multiple of mul(a,b)
	// eg mul(2,3) -> 6

	trimmedString := strings.TrimSuffix(strings.TrimPrefix(inputStr, "mul("), ")")

	nums := strings.Split(trimmedString, ",")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])

	return num1 * num2
}

func main() {
	// Part 1
	data, err := os.ReadFile("input.txt")
	check(err)

	mulReg := `mul\(\d{1,3},\d{1,3}\)`
	mulRegExp := regexp.MustCompile(mulReg)

	matches := mulRegExp.FindAllString(string(data), -1)

	var totalSumOfMultiples int = 0
	for _, match := range matches {
		totalSumOfMultiples += getMul(match)
	}

	fmt.Printf("Part 1 \nTotal Sum of Multiples: %v \n", totalSumOfMultiples)

	// Part 2
	doDontReg := `do\(\)|don't\(\)`
	doDontRegExp := regexp.MustCompile(doDontReg)

	isEnabled := true

	splitByDoDont := doDontRegExp.Split(string(data), -1)
	matchDoDont := doDontRegExp.FindAllString(string(data), -1)

	var totalDoDontSum int = 0
	for i, split := range splitByDoDont {
		mulMatches := mulRegExp.FindAllString(split, -1)
		for _, match := range mulMatches {
			if isEnabled {
				totalDoDontSum += getMul(match)
			}
		}

		if i < len(matchDoDont) {
			if matchDoDont[i] == "do()" {
				isEnabled = true
			} else {
				isEnabled = false
			}
		}
	}

	fmt.Printf("Part 2 \nTotal Sum of Multiples: %v \n", totalDoDontSum)
}
