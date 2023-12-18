package day5

import (
	"errors"
	"math"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func FromAlmanac(lines []string) ([]int, *Graph[*RangeCollection], *Mapper, int, []int, error) {
	mapper := NewMapper()
	seeds, graph, err := ParseAlmanac(lines, 8, mapper)

	if err != nil {
		return nil, nil, nil, 0, nil, err
	}

	start, path := GetPathFromSeedToLocation(graph, mapper)

	if path == nil {
		return nil, nil, nil, 0, nil, errors.New("no path from seeds to location")
	}

	return seeds, graph, mapper, start, path, nil
}

func GetMinLocationFromSeedRanges(seedRanges []int, graph *Graph[*RangeCollection], start int, path []int) (int, error) {
	if len(seedRanges)%2 != 0 {
		return 0, errors.New("invalid seed ranges slice")
	}

	currMin := math.MaxInt

	for i := 0; i < len(seedRanges); i += 2 {
		rangeStart := seedRanges[i]
		rangeEnd := rangeStart + seedRanges[i+1]

		for seed := rangeStart; seed < rangeEnd; i++ {
			loc, err := GetLocation(seed, graph, start, path)
			if err != nil {
				return 0, err
			}

			currMin = min(currMin, loc)
		}
	}

	return currMin, nil
}
