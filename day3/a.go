package day3

import (
	"fmt"
	"log"

	common "github.com/MiguelTornero/advent-of-code-2023-golang"
)

func Foo() string {
	return "bar"
}

func IsSymbol(b rune) bool {
	isDigit := common.IsDigit(b)
	isDot := b == '.'

	return !isDigit && !isDot
}

func PadWithDots(row []rune, size int) []rune {
	diff := size - len(row)

	for i := 0; i < diff; i++ {
		row = append(row, '.')
	}

	return row[:size]
}

func GetPaddedMatrix(lines []string) [][]rune {
	rows := len(lines) + 2
	output := make([][]rune, rows)

	length := 2
	if rows > 0 {
		length += len(lines[0])
	}

	dotRow := make([]rune, length)
	dotRow = PadWithDots(dotRow[:0], length)

	output[0] = dotRow
	output[rows-1] = dotRow

	for i, line := range lines {
		line := []rune(line)
		runes := append([]rune{'.'}, line...)
		runes = PadWithDots(runes, length)
		output[i+1] = runes
	}

	return output
}

func GetRuneMatrix(lines []string) [][]rune {
	output := make([][]rune, len(lines))

	for i, line := range lines {
		output[i] = []rune(line)
	}

	return output
}

func IsTouchingSymbol(paddedRunes [][]rune, row int, column int, height int, width int) bool {
	var ul, uc, ur, cr, cl, ll, lc, lr rune = '.', '.', '.', '.', '.', '.', '.', '.'

	if row > 0 {
		if column > 0 {
			// upper left
			ul = paddedRunes[row-1][column-1]
		}

		// upper center
		uc = paddedRunes[row-1][column]

		if column < width-1 {
			// upper right
			ur = paddedRunes[row-1][column+1]
		}
	}

	if column > 0 {
		// left
		cl = paddedRunes[row][column-1]
	}

	if column < width-1 {
		// right
		cr = paddedRunes[row][column+1]
	}

	if row < height-1 {
		if column > 0 {
			// lower left
			ll = paddedRunes[row+1][column-1]
		}

		// lower center
		lc = paddedRunes[row+1][column]

		if column < width-1 {
			// lower right
			lr = paddedRunes[row+1][column+1]
		}
	}
	return IsSymbol(ul) || IsSymbol(uc) || IsSymbol(ur) || IsSymbol(cr) || IsSymbol(cl) || IsSymbol(ll) || IsSymbol(lc) || IsSymbol(lr)
}

func PrintMatrix(mat [][]rune) {
	for _, row := range mat {
		for _, c := range row {
			fmt.Print(string([]rune{c}), " ")
		}
		fmt.Println("")
	}
}

func getValidNumbersSum(lines []string) int {
	output := 0

	if len(lines) == 0 {
		return output
	}

	mat := GetRuneMatrix(lines)
	PrintMatrix(mat)

	var inNumber bool
	var isValid bool
	var currentNumber int

	rows := len(mat)
	for i := 0; i < rows; i++ {

		cols := len(mat[i])

		for j := 0; j < cols; j++ {
			c := mat[i][j]
			if common.IsDigit(c) {
				inNumber = true
				currentNumber *= 10
				currentNumber += int(c - '0')
				if !isValid {
					isValid = IsTouchingSymbol(mat, i, j, rows, cols)
				}
			} else if inNumber {
				if isValid {
					output += currentNumber

					fmt.Println("valid number:", currentNumber)
				} else {
					fmt.Println("not valid number:", currentNumber)
				}
				currentNumber = 0
				inNumber = false
				isValid = false
			}
		}

		if inNumber {
			if isValid {
				output += currentNumber

				fmt.Println("valid number:", currentNumber)
			} else {
				fmt.Println("not valid number:", currentNumber)
			}
			currentNumber = 0
			inNumber = false
			isValid = false
		}
	}

	return output
}

func PartA(inputName string) {
	lines, err := common.FromFile(inputName)
	if err != nil {
		log.Fatal(err)
	}

	res := getValidNumbersSum(lines)
	fmt.Println(res)
}
