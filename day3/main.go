package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
	}
	return input
}

func main() {
	// input := ReadFile("test.txt")
	input := ReadFile("input.txt")
	// fmt.Println(input)

	reFindMuls := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := reFindMuls.FindAllString(input, -1)

	reGetDigits := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	result := 0

	for _, match := range matches {
		numbers := reGetDigits.FindStringSubmatch(match)
		n1, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatalf("Error converting first value to int: %s", numbers)
		}

		n2, err := strconv.Atoi(numbers[2])
		if err != nil {
			log.Fatalf("Error converting first value to int: %s", numbers)
		}

		result += n1 * n2
	}

	fmt.Println(result)
}
