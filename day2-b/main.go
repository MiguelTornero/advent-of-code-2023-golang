package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	common "github.com/MiguelTornero/advent-of-code-2023-golang"
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

func (g *gameRound) Print() {
	delimiter := "------------------------------------"
	fmt.Println(delimiter)
	fmt.Println("id:", g.id)

	for i, set := range g.sets {
		fmt.Printf("%d: red=%d green=%d blue=%d\n", i, set.red, set.green, set.blue)
	}
	fmt.Println(delimiter)
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

func (s *cubeSet) GetPower() int {
	return s.red * s.green * s.blue
}

func (g *gameRound) GetMinimunCubeSet() *cubeSet {
	output := &cubeSet{
		red:   0,
		green: 0,
		blue:  0,
	}

	for _, set := range g.sets {
		if set.red > output.red {
			output.red = set.red
		}
		if set.green > output.green {
			output.green = set.green
		}
		if set.blue > output.blue {
			output.blue = set.blue
		}
	}

	return output
}

func sumMinimumSetPowers(games []*gameRound) int {
	output := 0
	for _, game := range games {
		minSet := game.GetMinimunCubeSet()
		output += minSet.GetPower()
	}

	return output
}

func main() {

	lines, err := common.FromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	games, err := parseGameLines(lines)
	if err != nil {
		log.Fatal(err)
	}

	result := sumMinimumSetPowers(games)
	fmt.Println("RESULT:", result)
}
