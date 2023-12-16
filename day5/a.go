package day5

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Mapper struct {
	counter int
	items   map[string]int
}

func (m *Mapper) GetItem(name string) int {
	val, ok := m.items[name]
	if ok {
		return val
	}

	old := m.counter
	m.counter++
	m.items[name] = old

	return old
}

func NewMapper() *Mapper {
	m := &Mapper{
		counter: 0,
		items:   map[string]int{},
	}

	return m
}

type Graph[T any] struct {
	size   int
	matrix []T
}

func NewGraph[T any](elements int) *Graph[T] {
	if elements < 0 {
		panic(errors.New("invalid graph size"))
	}
	mat := make([]T, elements*elements)
	return &Graph[T]{
		size:   elements,
		matrix: mat,
	}
}

func (g *Graph[T]) GetSize() int {
	return g.size
}

func (g *Graph[T]) GetEdge(from int, to int) (T, error) {
	var a T
	if from < 0 || from >= g.size || to < 0 || to >= g.size {
		return a, errors.New("index out of range")
	}

	index := from*g.size + to
	a = g.matrix[index]

	return a, nil
}

func (g *Graph[T]) SetEdge(from int, to int, value T) (T, error) {
	var a T
	if from < 0 || from >= g.size || to < 0 || to >= g.size {
		return a, errors.New("index out of range")
	}

	index := from*g.size + to
	g.matrix[index] = value

	return value, nil
}

type Range struct {
	start  int
	end    int
	offset int
}

func NewRange(start int, end int, offset int) *Range {
	if end < start {
		panic(errors.New("invalid range"))
	}

	return &Range{
		start:  start,
		end:    end,
		offset: offset,
	}
}

func (r *Range) InRange(n int) bool {
	return n >= r.start && n <= r.end
}

func (r *Range) Transform(n int) int {
	if r.InRange(n) {
		return n + r.offset
	}

	return n
}

type RangeCollection struct {
	ranges []*Range
}

func NewRangeCollection() *RangeCollection {
	return &RangeCollection{
		ranges: []*Range{},
	}
}

func (rc *RangeCollection) InRange(n int) bool {
	for _, r := range rc.ranges {
		if r.InRange(n) {
			return true
		}
	}

	return false
}

func (rc *RangeCollection) GetSize() int {
	return len(rc.ranges)
}

func (rc *RangeCollection) AddRage(start int, end int, offset int) *RangeCollection {
	r := NewRange(start, end, offset)
	rc.ranges = append(rc.ranges, r)

	return rc
}

func (rc *RangeCollection) Transform(n int) int {
	for _, r := range rc.ranges {
		if r.InRange(n) {
			return r.Transform(n)
		}
	}

	return n
}

func (rc *RangeCollection) GetFirstMatchingRange(n int) *Range {
	for _, r := range rc.ranges {
		if r.InRange(n) {
			return r
		}
	}

	return nil
}

func ParseMapNums(lines []string, m *Mapper) (*RangeCollection, error) {
	output := NewRangeCollection()
	r := regexp.MustCompile(`\s+`)

	FIELDS := 3 // a map string has three numbers

	for _, line := range lines {
		line = strings.TrimSpace(line)
		lineErr := errors.New("invalid map line: " + line)

		if line == "" {
			return output, nil
		}

		numsStr := r.Split(line, FIELDS+1)

		if len(numsStr) != FIELDS {
			return nil, lineErr
		}

		nums := make([]int, FIELDS)
		for i, numStr := range numsStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, lineErr
			}
			nums[i] = num
		}

		start := nums[1]
		end := nums[1] + nums[2] - 1
		offset := nums[0]

		output.AddRage(start, end, offset)

	}

	return output, nil
}

func ParseAlmanac(lines []string, maxElems int, m *Mapper) error {
	seedsRegex := regexp.MustCompile(`^seeds:(.*)$`)
	mapRegex := regexp.MustCompile(`^([A-Za-z]+)-to-([A-Za-z]+)\s+map:$`)

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue // ignore empty lines
		}

		seedsMatch := seedsRegex.FindStringSubmatch(line)
		if len(seedsMatch) > 0 {
			// parse seed here
			continue
		}

		mapMatch := mapRegex.FindStringSubmatch(line)
		if len(mapMatch) > 0 {
			rc, err := ParseMapNums(lines[i+1:], m)
			if err != nil {
				return err
			}
			size := rc.GetSize()
			i = i + size

			continue
		}

		return errors.New("invalid almanac string at line " + strconv.Itoa(i+1))

	}

	return nil
}