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
	for i := 1; i < len(report); i++ {
		diff := report[i-1] - report[i]
		if (diff < 0) {
			diff = diff * -1
		}
    if (!(diff >= 1 && diff <= 3)) {
      return false
    }
	}
	return true
}

func problemDampener(report []int) bool {
  fmt.Println("OG", report)
  index := 0
  report = append(report[:index], report[index+1:]...)
  fmt.Println("NO", report)




	// for i := 0; i < len(report); i++ {
	//    fmt.Println("OG", report, i)
	//    slidingSlice := append(report[:i], report[i+1:]...)
	//    // slidingSlice = append(report[:i], report[i+1:]...)
	//    // slidingSlice := make([]int, len(report))
	//    // copy(slidingSlice, report)
	//    // slidingSlice[i]--
	//    fmt.Println("OG", report, i)
	//    fmt.Println("NN", slidingSlice, i)
	//
	// 	if isIncreasing(slidingSlice) || isDecreasing(slidingSlice) {
	// 		if (adjacentNumbersDiffersAtLeastOneAndAtMostThree(slidingSlice)) {
	//        fmt.Println("YES")
	//        return true
	// 		}
	// 	}
	// }

  return false
}

func main() {
	input := ReadFileLineByLine("test.txt")
	// input := ReadFileLineByLine("input.txt")
	safeCounter1 := 0
	// safeCounter2 := 0
	for _, item := range input {
		if isIncreasing(item) || isDecreasing(item) {
			if (adjacentNumbersDiffersAtLeastOneAndAtMostThree(item)) {
				safeCounter1++
				// safeCounter2++
			}
			if (problemDampener(item)) {
				safeCounter1++
			}
		}
	}

  // part 1
	// fmt.Println(safeCounter1)

  // part 2
	fmt.Println(safeCounter1)
}
