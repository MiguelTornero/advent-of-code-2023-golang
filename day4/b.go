package day4

import (
	"fmt"
	"log"

	common "github.com/MiguelTornero/advent-of-code-2023-golang"
)

func GetCardsTotal(lines []string) (int, error) {
	cardCountList := make([]int, len(lines))

	// initializing as 1's
	for i := range cardCountList {
		cardCountList[i] = 1
	}

	for i, line := range lines {
		fmt.Println(cardCountList)
		winCount, err := GetLineWinCount(line)

		if err != nil {
			return 0, err
		}

		cardCount := cardCountList[i]
		for j := 1; j <= winCount; j++ {
			cardCountList[i+j] += cardCount
		}
	}
	return common.Sum(cardCountList), nil
}

func PartB(filename string) {
	lines, err := common.FromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result, err := GetCardsTotal(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESULT:", result)
}
