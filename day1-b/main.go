package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode/utf8"

	common "github.com/MiguelTornero/advent-of-code-2023-golang"
)

func getTextMatch(s string, m map[string]int) (int, int) {
	for key, value := range m {
		if strings.HasPrefix(s, key) {
			return value, len(key)
		}
	}

	return 0, -1
}

func getFirstNumber(s string, m map[string]int) (int, int) {
	length := len(s)
	for i := 0; i < length; i++ {
		r, _ := utf8.DecodeRuneInString(s[i:])
		if common.IsDigit(r) {
			num := int(r - '0')
			return num, i + 1
		}

		number, matchSize := getTextMatch(s[i:], m)

		if matchSize >= 0 {
			return number, i + matchSize
		}
	}

	return 0, -1
}

func getLastNumber(s string, m map[string]int) (int, int) {
	number, matchIndex := 0, -1
	for i, r := range s {
		if common.IsDigit(r) {
			number, matchIndex = int(r-'0'), i
			continue
		}

		matchedNumber, matchSize := getTextMatch(s[i:], m)

		if matchSize >= 0 {
			number, matchIndex = matchedNumber, i+matchSize
		}
	}

	return number, matchIndex
}

func sumLines(lines []string, m map[string]int) (int, error) {
	output := 0

	for _, line := range lines {
		first, i := getFirstNumber(line, m)
		if i < 0 {
			return 0, errors.New("no number found in string")
		}
		second, ii := getLastNumber(line[i:], m)
		if ii < 0 {
			second = first
		}

		output += first*10 + second
	}

	return output, nil
}

func main() {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	lines, err := common.FromFile("input.txt")

	if err != nil {
		log.Fatalf("%v", err)
	}

	result, err := sumLines(lines, numbers)

	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("RESULT: %d\n", result)
}
