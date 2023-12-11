package day4_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/MiguelTornero/advent-of-code-2023-golang/day4"
)

func TestB(t *testing.T) {
	lines := strings.Split(input, "\n")

	result, err := day4.GetCardsTotal(lines)

	if err != nil {
		_ = fmt.Errorf("%v", err)
	}

	if result != 30 {
		_ = fmt.Errorf("not passed. (expected %d, got %d)", 30, result)
	}
}
