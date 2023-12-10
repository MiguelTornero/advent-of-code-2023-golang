package main

import (
	"os"

	"github.com/MiguelTornero/advent-of-code-2023-golang/day3"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "a":
			day3.PartA("input.txt")
			os.Exit(0)
		case "b":
			day3.PartB("input.txt")
			os.Exit(0)
		}
	}
	os.Exit(1)
}
