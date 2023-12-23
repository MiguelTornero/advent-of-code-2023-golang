package day5_test

import (
	"strings"
	"testing"

	"github.com/MiguelTornero/advent-of-code-2023-golang/day5"
	"github.com/stretchr/testify/assert"
)

func TestSeedRanges(t *testing.T) {
	lines := strings.Split(almanacStr, "\n")

	seedRanges, graph, _, start, path, err := day5.FromAlmanac(lines)

	assert.Nil(t, err)

	result, err := day5.GetMinLocationFromSeedRanges(seedRanges, graph, start, path)

	assert.Nil(t, err)

	assert.Equal(t, 46, result)
}

func TestRangeSplitter(t *testing.T) {
	splitStart, splitEnd := 5, 13

	part1Start, part1End, part2Start, part2End, part3Start, part3End := day5.SplitRange(0, 3, splitStart, splitEnd)
	assert.Equal(t, 0, part1Start)
	assert.Equal(t, 3, part1End)
	assert.Equal(t, 5, part2Start)
	assert.Equal(t, 4, part2End)
	assert.Equal(t, 14, part3Start)
	assert.Equal(t, 13, part3End)

	part1Start, part1End, part2Start, part2End, part3Start, part3End = day5.SplitRange(6, 9, splitStart, splitEnd)
	assert.Equal(t, 5, part1Start)
	assert.Equal(t, 4, part1End)
	assert.Equal(t, 6, part2Start)
	assert.Equal(t, 9, part2End)
	assert.Equal(t, 14, part3Start)
	assert.Equal(t, 13, part3End)

	part1Start, part1End, part2Start, part2End, part3Start, part3End = day5.SplitRange(15, 18, splitStart, splitEnd)
	assert.Equal(t, 5, part1Start)
	assert.Equal(t, 4, part1End)
	assert.Equal(t, 14, part2Start)
	assert.Equal(t, 13, part2End)
	assert.Equal(t, 15, part3Start)
	assert.Equal(t, 18, part3End)

	part1Start, part1End, part2Start, part2End, part3Start, part3End = day5.SplitRange(0, 6, splitStart, splitEnd)
	assert.Equal(t, 0, part1Start)
	assert.Equal(t, 4, part1End)
	assert.Equal(t, 5, part2Start)
	assert.Equal(t, 6, part2End)
	assert.Equal(t, 14, part3Start)
	assert.Equal(t, 13, part3End)

	part1Start, part1End, part2Start, part2End, part3Start, part3End = day5.SplitRange(9, 18, splitStart, splitEnd)
	assert.Equal(t, 5, part1Start)
	assert.Equal(t, 4, part1End)
	assert.Equal(t, 9, part2Start)
	assert.Equal(t, 13, part2End)
	assert.Equal(t, 14, part3Start)
	assert.Equal(t, 18, part3End)

	part1Start, part1End, part2Start, part2End, part3Start, part3End = day5.SplitRange(0, 18, splitStart, splitEnd)
	assert.Equal(t, 0, part1Start)
	assert.Equal(t, 4, part1End)
	assert.Equal(t, 5, part2Start)
	assert.Equal(t, 13, part2End)
	assert.Equal(t, 14, part3Start)
	assert.Equal(t, 18, part3End)
}
