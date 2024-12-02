package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkSafeReport(report []int) bool {
	/*
		- if list all gradually increase or dec => true
		- atleast 1 or max 3 inc or dec => true

		- if two numbers are same => false
		- if inc then dec => false
		- if dec then inc => false
		- change of more than >3 => false
	*/

	var diff int = 0
	var prevDiff int = report[0] - report[1]

	for i := range len(report) - 1 {
		diff = report[i] - report[i+1]
		if diff > 3 || diff < -3 || diff == 0 {
			return false
		}

		if prevDiff > 0 && diff < 0 {
			return false
		} else if prevDiff < 0 && diff > 0 {
			return false
		}
		prevDiff = diff
	}

	return true
}

func problemDampner(report []int) bool {
	/*
		using checkSafeReport func
		if true return true

		- remove each element from report
		- check if new report is safe
		- if true return true
	*/

	if checkSafeReport(report) {
		return true
	}

	for i := range len(report) {
		// making explicit copy of slice
		newReport := make([]int, len(report))
		copy(newReport, report)

		// remove element at i
		newReport = append(newReport[:i], newReport[i+1:]...)

		// fmt.Println(report, newReport)

		// check if newReport is safe
		if checkSafeReport(newReport) {
			return true
		}
	}

	return false
}

func main() {
	// Part 1 - Red Nosed Report
	data, err := os.ReadFile("input.txt")
	check(err)

	var reports [][]int
	lines := strings.Split(string(data), "\n")
	reports = make([][]int, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		reports[i] = make([]int, len(parts))

		for j, part := range parts {
			num, _ := strconv.Atoi(part)
			reports[i][j] = num
		}
	}

	var totalSafeReports int = 0
	for _, report := range reports {
		if checkSafeReport(report) {
			totalSafeReports++
		}
	}

	fmt.Printf("Part 1 \nTotal Safe Reports: %d \n", totalSafeReports)

	// Part 2 - Problem Dampner
	var totalSafeReportsAfterDampner int = 0
	for _, report := range reports {
		if problemDampner(report) {
			totalSafeReportsAfterDampner++
		}
	}

	fmt.Printf("Part 2 \nTotal Safe Reports w/ Problem Dampner: %d \n", totalSafeReportsAfterDampner)
}
