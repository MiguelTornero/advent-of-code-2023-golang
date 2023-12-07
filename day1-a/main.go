package main

import (
	"errors"
	"fmt"
	"log"

	common "github.com/MiguelTornero/advent-of-code-2023-golang"
)

func getLastDigit(s string) (rune, int) {
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		if common.IsDigit(runes[i]) {
			return runes[i], i
		}
	}
	return 0, -1
}

func getFirstDigit(s string) (rune, int) {
	for i, r := range s {
		if common.IsDigit(r) {
			return r, i
		}
	}
	return 0, -1
}

func getIntFromTwoDigits(first rune, second rune) int {
	output := (first-'0')*10 + (second - '0')
	return int(output)
}

func sumLines(lines []string) (int, error) {
	output := 0
	for _, line := range lines {
		first, i := getFirstDigit(line)
		if i < 0 {
			return 0, errors.New("no digit found in string")
		}
		second, _ := getLastDigit(line)
		number := getIntFromTwoDigits(first, second)
		output += number
	}
	return output, nil
}

func main() {
	lines, err := common.FromFile("input.txt")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	output, err := sumLines(lines)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	fmt.Printf("RESULT: %d\n", output)
}
