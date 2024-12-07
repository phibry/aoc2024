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
	return true;
}

func adjacentNumbersDiffersAtLeastOneAndAtMostThree(report []int) bool {
	bools := false
	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]
		if (diff < 0) {
			diff = diff * -1
		}
		bools = diff >= 1 && diff <= 3
	}
	return bools
}

func main() {
	input := ReadFileLineByLine("test.txt")

	// [7 6 4 2 1]
	safeCounter := 0
	for _, item := range input {
		if isIncreasing(item) || isDecreasing(item) {
			if (adjacentNumbersDiffersAtLeastOneAndAtMostThree(item)) {
				fmt.Println(item)
				safeCounter++
			}
		}
		// for i:= 0; i < len(item); i++ {
		//
		// }
	}

	fmt.Println(safeCounter)

}
