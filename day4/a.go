package day4

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	common "github.com/MiguelTornero/advent-of-code-2023-golang"
)

func GetLineWinCount(line string) (int, error) {
	count := 0

	err := errors.New("invalid card string")
	parts := regexp.MustCompile(`^Card\s+\d+:`).Split(line, 2)

	if len(parts) != 2 {
		return 0, err
	}

	nums := strings.SplitN(parts[1], "|", 3)
	if len(nums) < 2 {
		return 0, err
	}

	re := regexp.MustCompile(`\s+`)
	winningNumsStr := strings.TrimSpace(nums[0])
	winningNums := re.Split(winningNumsStr, -1)

	obtainedNumsStr := strings.TrimSpace(nums[1])
	obtainedNums := re.Split(obtainedNumsStr, -1)

	winningNumsMap := map[int]bool{}
	for _, winningNumStr := range winningNums {
		num, err := strconv.Atoi(winningNumStr)
		if err != nil {
			return 0, err
		}

		winningNumsMap[num] = true
	}

	for _, obtainedNumStr := range obtainedNums {
		num, err := strconv.Atoi(obtainedNumStr)
		if err != nil {
			return 0, err
		}

		_, ok := winningNumsMap[num]
		if ok {
			fmt.Println("Found an obtained winning number:", num)
			count++
		}
	}

	return count, nil
}

func getLineScore(line string) (int, error) {
	count, err := GetLineWinCount(line)

	if err != nil {
		return 0, err
	}

	if count <= 0 {
		return 0, nil
	}

	score := 1
	// porpusefully starting at 1 so the result is 2^(count - 1)
	for i := 1; i < count; i++ {
		score *= 2
	}

	return score, nil
}

func GetWinningNumbersSum(lines []string) (int, error) {
	output := 0

	for _, line := range lines {
		score, err := getLineScore(line)
		if err != nil {
			return 0, err
		}
		output += score

		fmt.Println("Card score:", score)
	}

	return output, nil
}

func PartA(filename string) {
	lines, err := common.FromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result, err := GetWinningNumbersSum(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESULT:", result)
}
