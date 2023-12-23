package main

import (
	"os"

	"github.com/MiguelTornero/advent-of-code-2023-golang/day5"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "a":
			day5.PartA("input.txt")
		case "b":
			day5.PartB("input.txt")
		}
		os.Exit(0)
	}
	os.Exit(1)
}
