package day5_test

import (
	"strings"
	"testing"

	"github.com/MiguelTornero/advent-of-code-2023-golang/day5"
	"github.com/stretchr/testify/assert"
)

func TestMapper(t *testing.T) {
	mapper := day5.NewMapper()

	assert.Equal(t, mapper.GetItem("foo"), 0)

	assert.Equal(t, mapper.GetItem("foo"), 0)

	assert.Equal(t, mapper.GetItem("bar"), 1)

	assert.Equal(t, mapper.GetItem("hello"), 2)

	assert.Equal(t, mapper.GetItem("world"), 3)

	assert.Equal(t, mapper.GetItem("bar"), 1)

}

func TestGraphInt(t *testing.T) {
	g := day5.NewGraph[int](10)

	assert.Equal(t, g.GetSize(), 10)

	v, err := g.GetEdge(0, 1)

	assert.Nil(t, err)

	assert.Equal(t, v, 0)

	_, err = g.SetEdge(0, 1, 32)

	assert.Nil(t, err)

	v, err = g.GetEdge(0, 1)

	assert.Nil(t, err)

	assert.Equal(t, v, 32)
}

func TestGraphErr(t *testing.T) {

	assert.Panics(t, func() { day5.NewGraph[bool](-1) })

	g := day5.NewGraph[bool](5)

	assert.NotNil(t, g)

	_, err := g.SetEdge(0, 100, false)

	assert.NotNil(t, err)

	_, err = g.SetEdge(100, 0, false)

	assert.NotNil(t, err)

	_, err = g.SetEdge(100, 100, false)

	assert.NotNil(t, err)
}

func TestRange(t *testing.T) {
	assert.Panics(t, func() { day5.NewRange(5, 2, -1) })

	r := day5.NewRange(5, 10, -1)

	assert.Equal(t, r.InRange(5), true)

	assert.Equal(t, r.InRange(10), true)

	assert.Equal(t, r.InRange(0), false)

	assert.Equal(t, r.Transform(8), 7)

	assert.Equal(t, r.Transform(1), 1)
}

func TestRangeCollection(t *testing.T) {
	rc := day5.NewRangeCollection()

	assert.Equal(t, rc.GetSize(), 0)

	rc.AddRage(0, 9, -2)

	rc.AddRage(20, 29, 2)

	assert.Equal(t, rc.GetSize(), 2)

	assert.Equal(t, rc.Transform(5), 3)

	assert.Equal(t, rc.Transform(22), 24)

	assert.Equal(t, rc.Transform(13), 13)

	assert.Nil(t, rc.GetFirstMatchingRange(34))

	assert.Equal(t, rc.GetFirstMatchingRange(7).Transform(7), rc.Transform(7))

	assert.True(t, rc.InRange(8))

	assert.True(t, rc.InRange(23))

	assert.False(t, rc.InRange(12))

	assert.Equal(t, rc.GetFirstMatchingRange(9).InRange(9), rc.InRange(9))

	assert.Panics(t, func() { rc.AddRage(0, -1, 5) })
}

const almanacStr string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestParseMapNums(t *testing.T) {
	input := []string{"1 2 3", "4 5 6", ""}
	mapper := day5.NewMapper()

	a, b := day5.ParseMapNums(input, mapper)

	assert.NotNil(t, a)
	assert.Nil(t, b)

	r := a.GetFirstMatchingRange(2)
	assert.NotNil(t, r)

	r = a.GetFirstMatchingRange(5)
	assert.NotNil(t, r)

	assert.Equal(t, a.Transform(2), 3)
	assert.Equal(t, a.Transform(4), 5)

	assert.Equal(t, a.Transform(0), 0)

	assert.Equal(t, a.Transform(5), 9)
	assert.Equal(t, a.Transform(10), 14)
}

func TestAlmanacParser(t *testing.T) {
	lines := strings.Split(almanacStr, "\n")
	mapper := day5.NewMapper()

	_, _, err := day5.ParseAlmanac([]string{"a", "b"}, 10, day5.NewMapper())
	assert.NotNil(t, err)

	seeds, g, err := day5.ParseAlmanac(lines, 10, mapper)
	assert.Nil(t, err)

	assert.Equal(t, []int{79, 14, 55, 13}, seeds)

	assert.NotNil(t, g)

	seedNum := mapper.GetItem("seed")
	soilNum := mapper.GetItem("soil")

	rc, err := g.GetEdge(seedNum, soilNum)
	assert.NotNil(t, rc)
	assert.Nil(t, err)
}
