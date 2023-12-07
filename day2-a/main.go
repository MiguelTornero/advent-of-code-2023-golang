package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	adventofcode2023golang "github.com/MiguelTornero/advent-of-code-2023-golang"
)

type cubeSet struct {
	red   int
	green int
	blue  int
}

type gameRound struct {
	id   int
	sets []*cubeSet
}

func parseSetString(s string) (*cubeSet, error) {
	s = strings.TrimSpace(s)
	split := strings.Split(s, ",")
	itemRegex, _ := regexp.Compile(`^(\d+)\s+([A-Za-z]+)$`)
	numbers := map[string]int{
		"red":   0,
		"blue":  0,
		"green": 0,
	}

	for _, item := range split {
		item := strings.TrimSpace(item)
		groups := itemRegex.FindStringSubmatch(item)

		if groups == nil {
			return nil, errors.New("invalid set string: " + item)
		}

		itemNumStr := groups[1]
		itemNum, _ := strconv.Atoi(itemNumStr)
		itemColorStr := groups[2]

		n, ok := numbers[itemColorStr]
		if !ok {
			continue // ignore invalid colors
		}
		numbers[itemColorStr] = n + itemNum
	}

	outputSet := &cubeSet{
		red:   numbers["red"],
		blue:  numbers["blue"],
		green: numbers["green"],
	}

	return outputSet, nil
}

func parseSetLineString(s string) ([]*cubeSet, error) {
	s = strings.TrimSpace(s)
	split := strings.Split(s, ";")
	output := []*cubeSet{}

	for _, item := range split {
		set, err := parseSetString(item)
		if err != nil {
			return nil, err
		}
		output = append(output, set)
	}

	return output, nil
}

func parseGameString(s string) (*gameRound, error) {
	s = strings.TrimSpace(s)
	r, _ := regexp.Compile(`^Game\s+(\d+):\s+(.*)$`)

	groups := r.FindStringSubmatch(s)

	if groups == nil {
		return nil, errors.New("invalid game string: " + s)
	}

	gameIdStr := groups[1]
	setsStr := groups[2]

	gameId, _ := strconv.Atoi(gameIdStr)
	sets, err := parseSetLineString(setsStr)

	if err != nil {
		return nil, err
	}

	return &gameRound{
		id:   gameId,
		sets: sets,
	}, nil
}

func parseGameLines(lines []string) ([]*gameRound, error) {
	output := []*gameRound{}

	for _, line := range lines {
		game, err := parseGameString(line)
		if err != nil {
			return nil, err
		}

		output = append(output, game)
	}

	return output, nil
}

func (g *gameRound) isPossible(maxRed int, maxGreen int, maxBlue int) bool {
	red, green, blue := 0, 0, 0
	for _, set := range g.sets {
		red += set.red
		green += set.green
		blue += set.blue
	}

	if red > maxRed {
		return false
	}
	if green > maxGreen {
		return false
	}
	if blue > maxBlue {
		return false
	}

	return true
}

func getPossibleGamesSum(games []*gameRound, maxRed int, maxGreen int, maxBlue int) int {
	output := 0

	for _, game := range games {
		if game.isPossible(maxRed, maxGreen, maxBlue) {
			output += game.id
		}
	}

	return output
}

func main() {
	const MAX_RED = 12
	const MAX_GREEN = 13
	const MAX_BLUE = 14

	lines, err := adventofcode2023golang.FromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	games, err := parseGameLines(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sum of possible game IDs with %d reds, %d greens, and %d blues:\n", MAX_RED, MAX_GREEN, MAX_BLUE)

	impossibleGames := getPossibleGamesSum(games, MAX_RED, MAX_GREEN, MAX_BLUE)
	fmt.Println(impossibleGames)
}
