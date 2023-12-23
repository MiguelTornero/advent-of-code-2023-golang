package day5

import (
	"errors"
	"fmt"
	"log"
	"math"

	common "github.com/MiguelTornero/advent-of-code-2023-golang"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func clamp(n int, min int, max int) int {
	if n < min {
		return min
	}
	if n > max {
		return max
	}

	return n
}

func TestClamp(a int, b int, c int) int {
	return clamp(a, b, c)
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
		rangeEnd := rangeStart + seedRanges[i+1] - 1

		fmt.Println("TESTING RANGE:", rangeStart, "-", rangeEnd)

		lowerLim, _, err := WalkPathGetRange(rangeStart, rangeEnd, graph, start, path)
		if err != nil {
			return 0, err
		}

		currMin = min(currMin, lowerLim)
	}

	return currMin, nil
}

func (r *Range) Overlaps(start int, end int) bool {
	return !(end < r.start || start > r.end)
}

func (rc *RangeCollection) GetOverlappingRanges(start int, end int) []*Range {
	output := []*Range{}
	for _, r := range rc.ranges {
		if r.Overlaps(start, end) {
			output = append(output, r)
		}
	}

	return output
}

func (rc *RangeCollection) GetLimitsFromRange(start int, end int) (int, int) {
	innerLower, innerUpper := math.MaxInt, math.MinInt

	outerLower, outerUpper := start, end

	if len(rc.ranges) < 1 {
		return start, end
	}

	for _, r := range rc.ranges {
		beforelower, beforeupper, rlower, rupper, afterlower, afterupper := SplitRange(start, end, r.start, r.end)

		var rOuterLower, rOuterUpper int
		if beforelower <= beforeupper {
			// range before exists
			rOuterLower = beforelower
		} else {
			// range before does not exist, go up to after range
			rOuterLower = afterlower
		}

		if afterlower <= afterupper {
			// range after exists
			rOuterUpper = afterlower
		} else {
			// range after does not exist, lower to before range
			rOuterUpper = beforeupper
		}

		// update current inner limits
		if rlower <= rupper {
			innerLower = min(innerLower, r.Transform(rlower))
			innerUpper = max(innerUpper, r.Transform(rupper))

			// opdate current outerlimits with intersection of current outer and this outer
			_, _, outerLower, outerUpper, _, _ = SplitRange(outerLower, outerUpper, rOuterLower, rOuterUpper)
		}
		fmt.Println("Limits of transform with start", r.start, ", end", r.end, "and offset", r.offset, "are", r.Transform(rlower), "-", r.Transform(rupper), "given the range", start, "-", end, "(", rlower <= rupper, ")")
		fmt.Println("Acumulated outer limits:", outerLower, "-", outerUpper, "(outer),", innerLower, "-", innerUpper, "(inner)")
	}

	fmt.Println("Outer limits of tranform collection is", outerLower, "-", outerUpper, "given the range", start, "-", end)

	outputLower := innerLower
	outputUpper := innerUpper
	if outerLower <= outerUpper {
		// outer range exists
		outputLower = min(outerLower, innerLower)
		outputUpper = max(outerUpper, innerUpper)
	}

	fmt.Println("Limits are", outputLower, "-", outputUpper)

	return outputLower, outputUpper

}

func SplitRange(start int, end int, splitStart int, splitEnd int) (int, int, int, int, int, int) {
	part1Start := min(start, splitStart)
	part1End := min(end, splitStart-1)

	part2Start := clamp(start, splitStart, splitEnd+1)
	part2End := clamp(end, splitStart-1, splitEnd)

	part3Start := max(start, splitEnd+1)
	part3End := max(end, splitEnd)

	return part1Start, part1End, part2Start, part2End, part3Start, part3End
}

func (r *Range) GetLimitsFromRange(start int, end int) (int, int) {
	lowerLimit, upperLimit := math.MaxInt, math.MinInt

	part1Start, part1End, part2Start, part2End, part3Start, part3End := SplitRange(start, end, r.start, r.end)

	if part1Start <= part1End {
		// part 1 exists
		lowerLimit = min(lowerLimit, r.Transform(part1Start))
		upperLimit = max(upperLimit, r.Transform(part1End))
	}

	if part2Start <= part2End {
		// part 2 exists
		lowerLimit = min(lowerLimit, r.Transform(part2Start))
		upperLimit = max(upperLimit, r.Transform(part2End))
	}

	if part3Start <= part3End {
		// part 3 exists
		lowerLimit = min(lowerLimit, r.Transform(part3Start))
		upperLimit = max(upperLimit, r.Transform(part3End))
	}

	return lowerLimit, upperLimit
}

func WalkPathGetRange(start int, end int, graph *Graph[*RangeCollection], startIndex int, path []int) (int, int, error) {
	if len(path) < 1 {
		return start, end, nil
	}

	rangeSplits := []int{start, end}
	processedRanges := []int{}

	lastIndex := startIndex
	for _, node := range path {
		rc, _ := graph.GetEdge(lastIndex, node)
		lastIndex = node

		if rc == nil {
			continue
		}

		for i := 0; i < len(rangeSplits); i += 2 {
			splitStart := rangeSplits[i]
			splitEnd := rangeSplits[i+1]

			fmt.Println("Testing ranges", rangeSplits[i:])
			fmt.Println("PROCESSED", processedRanges)

			r := rc.GetFirstOverlappingRange(splitStart, splitEnd)

			if r == nil {
				processedRanges = append(processedRanges, splitStart, splitEnd)
				fmt.Println("Transform not found")
				continue
			}

			beforeStart, beforeEnd, innerStart, innerEnd, afterStart, afterEnd := SplitRange(splitStart, splitEnd, r.start, r.end)
			fmt.Println("Found transfrom with range,", r.start, r.end)
			fmt.Println("Splits:", beforeStart, beforeEnd, innerStart, innerEnd, afterStart, afterEnd)

			if beforeStart <= beforeEnd {
				rangeSplits = append(rangeSplits, beforeStart, beforeEnd)
				fmt.Println("Adding range", beforeStart, beforeEnd)
			}
			if afterStart <= afterEnd {
				rangeSplits = append(rangeSplits, afterStart, afterEnd)
				fmt.Println("Adding range", afterStart, afterEnd)
			}

			if innerStart <= innerEnd { // should always be true, I think?
				processedRanges = append(processedRanges, r.Transform(innerStart), r.Transform(innerEnd))
				fmt.Println("added processed:", r.Transform(innerStart), r.Transform(innerEnd))
			}

		}

		fmt.Println("------------------------")
		rangeSplits = rangeSplits[:0]
		rangeSplits = append(rangeSplits, processedRanges...)
		processedRanges = processedRanges[:0]
	}
	fmt.Println("*********", rangeSplits, "************")

	outputLower, outputUpper := common.Min(rangeSplits), common.Max(rangeSplits)

	return outputLower, outputUpper, nil
}

func (rc *RangeCollection) GetFirstOverlappingRange(start int, end int) *Range {
	if end < start {
		return nil
	}

	for _, r := range rc.ranges {
		if r.Overlaps(start, end) {
			return r
		}
	}

	return nil
}

func PartB(filename string) {
	lines, err := common.FromFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	seedRanges, graph, _, start, path, err := FromAlmanac(lines)
	if err != nil {
		log.Fatal(err)
	}

	result, err := GetMinLocationFromSeedRanges(seedRanges, graph, start, path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESULT:", result)
}
