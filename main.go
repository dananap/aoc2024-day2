package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func removeLevel(report []int, element int) []int {
	newReport := make([]int, len(report));
	_ = copy(newReport, report)
	newReport = append(newReport[:element], newReport[element+1:]...)
	return newReport
}

func checkSafety(report []int) bool {
	increasing := report[0] < report[1]
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if increasing && diff < 0 {
			return false
		} else if !increasing && diff > 0 {
			return false
		}

		if diffAbs := math.Abs(float64(diff)); diffAbs < 1 || diffAbs > 3 {
			return false
		}
	}
	return true
}

func reportSafe(report []int) bool {
	if checkSafety(report) {
		return true
	}
	for i := range report {
		dampenedReport := removeLevel(report, i)
		if checkSafety(dampenedReport) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		reports = append(reports, make([]int, len(numbers)))

		for i, number := range numbers {
			value, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			reports[len(reports)-1][i] = value
		}
	}

	safeReports := 0
	for _, report := range reports {
		if reportSafe(report) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
