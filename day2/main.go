package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFileLineByLine(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		var reports []int
		for _, p := range parts {
			n, err := strconv.Atoi(p)
			if err != nil {
				log.Fatalf("Error converting first value to int: %s", parts)
			}
			reports = append(reports, n)
		}

		data = append(data, reports)
	}

	return data
}

func isIncreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] < report[i-1] {
			return false
		}
	}
	return true
}

func isDecreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1] {
			return false
		}
	}
	return true
}

func adjacentNumbersDiffersAtLeastOneAndAtMostThree(report []int) bool {
	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]
		if diff < 0 {
			diff = diff * -1
		}
		if !(diff >= 1 && diff <= 3) {
			return false
		}
	}
	return true
}

func problemDampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		data1 := append([]int{}, report[:i]...)
		data1 = append(data1, report[i+1:]...)
		if validate(data1) {
			return true
		}
	}

	return false
}

func validate(report []int) bool {
	return (isIncreasing(report) || isDecreasing(report)) && adjacentNumbersDiffersAtLeastOneAndAtMostThree(report)
}

func main() {
	// input := ReadFileLineByLine("test.txt")
	input := ReadFileLineByLine("input.txt")

	safeCounter1 := 0
	for _, item := range input {
		// if validate(item) {
		if validate(item) || problemDampener(item) {
			safeCounter1++
		}
	}

	// part 1 = 624
	// part 2 = 658
	fmt.Println(safeCounter1)
}
