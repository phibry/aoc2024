package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// func ReadFileLineByLine(path string) []string {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
//
// 	var output []string
//
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		output = append(output, scanner.Text())
// 	}
//
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	return output
// }

// func ReadFileLineByLine(path string) [][]int {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
//
// 	var data [][]int
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		parts := strings.Fields(line)
//
// 		if len(parts) != 2 {
// 			log.Fatalf("Invalid line format: %s", line)
// 		}
//
// 		n1, err := strconv.Atoi(parts[0])
// 		if err != nil {
// 			log.Fatalf("Error converting first value to int: %s", parts[0])
// 		}
//
// 		n2, err := strconv.Atoi(parts[1])
// 		if err != nil {
// 			log.Fatalf("Error converting first value to int: %s", parts[1])
// 		}
//
// 		data = append(data, []int{n1, n2})
// 	}
//
// 	return data
// }

func ReadFileLineByLine(path string) ([]int, []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var first, second []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) != 2 {
			log.Fatalf("Invalid line format: %s", line)
		}

		n1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Error converting first value to int: %s", parts[0])
		}

		n2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Error converting first value to int: %s", parts[1])
		}

		first = append(first, n1)
		second = append(second, n2)
	}

	return first, second
}

func calculateAbsoluteDifference(n1 int, n2 int) int {
	var difference = n1 - n2
	if difference < 0 {
		difference *= -1
	}

	return difference
}

func calculateSimilarityScore(n int, nList []int) int {
	count := 0
	for _, item := range nList {
		if item == n {
			count++
		}
	}
	return count * n
}

func main() {
	// first, second := ReadFileLineByLine("test.txt")
	first, second := ReadFileLineByLine("input.txt")
	sort.Ints(first)
	sort.Ints(second)

	var sum int
	var score int
	for i := 0; i < len(first); i++ {
		sum += calculateAbsoluteDifference(first[i], second[i])
		score += calculateSimilarityScore(first[i], second)
	}

	// part 1
	fmt.Println(sum)

	// part 2
	fmt.Println(score)

}
