package main

import "fmt"

func checkSafeReport(report []int) bool {
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

	if checkSafeReport(report) {
		return true
	}

	fmt.Println(report)

	for i := range len(report) {

		newReport := make([]int, len(report))
		copy(newReport, report)
		newReport = append(newReport[:i], newReport[i+1:]...)
		fmt.Println(i, report, newReport)

		if checkSafeReport(newReport) {
			return true
		}
	}

	return false
}

func main() {
	report := []int{62, 65, 67, 70, 73, 76, 75}

	if problemDampner(report) {
		println("Report is safe")
	} else {
		println("Report is not safe")
	}
}
