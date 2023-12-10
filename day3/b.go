package day3

import (
	"fmt"
	"log"
	"strconv"

	common "github.com/MiguelTornero/advent-of-code-2023-golang"
)

func GetAllAdjecentValues[T any](mat [][]T, row int, column int, height int, width int, defualtValue T) (T, T, T, T, T, T, T, T) {
	ul, uc, ur, cl, cr, ll, lc, lr := defualtValue, defualtValue, defualtValue, defualtValue, defualtValue, defualtValue, defualtValue, defualtValue

	if row > 0 {
		if column > 0 {
			// upper left
			ul = mat[row-1][column-1]
		}

		// upper center
		uc = mat[row-1][column]

		if column < width-1 {
			// upper right
			ur = mat[row-1][column+1]
		}
	}

	if column > 0 {
		// left
		cl = mat[row][column-1]
	}

	if column < width-1 {
		// right
		cr = mat[row][column+1]
	}

	if row < height-1 {
		if column > 0 {
			// lower left
			ll = mat[row+1][column-1]
		}

		// lower center
		lc = mat[row+1][column]

		if column < width-1 {
			// lower right
			lr = mat[row+1][column+1]
		}
	}

	return ul, uc, ur, cl, cr, ll, lc, lr
}

func GetWholeNumber(s []rune, index int) ([]rune, int, int) {
	if !common.IsDigit(s[index]) {
		return nil, -1, -1
	}

	lower := index - 1
	upper := index + 1
	length := len(s)

	for lower >= 0 && common.IsDigit(s[lower]) {
		lower--
	}
	lower++

	for upper < length && common.IsDigit(s[upper]) {
		upper++
	}

	return s[lower:upper], lower, upper
}

func AppendStrToIntSlice(arr []int, s string) []int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	arr = append(arr, num)

	return arr
}

func GetGearNumbers(mat [][]rune, row int, col int, height int, width int) []int {
	output := []int{}
	ul, uc, ur, cl, cr, ll, lc, lr := GetAllAdjecentValues[rune](mat, row, col, height, width, '.')

	//upper part
	if common.IsDigit(ul) {
		numStr, _, end := GetWholeNumber(mat[row-1], col-1)

		output = AppendStrToIntSlice(output, string(numStr))

		if end <= col && common.IsDigit(ur) {
			numStr, _, _ := GetWholeNumber(mat[row-1], col+1)

			output = AppendStrToIntSlice(output, string(numStr))
		}
	} else if common.IsDigit(uc) {
		numStr, _, _ := GetWholeNumber(mat[row-1], col)

		output = AppendStrToIntSlice(output, string(numStr))
	} else if common.IsDigit(ur) {
		numStr, _, _ := GetWholeNumber(mat[row-1], col+1)

		output = AppendStrToIntSlice(output, string(numStr))
	}

	// lower part
	if common.IsDigit(ll) {
		numStr, _, end := GetWholeNumber(mat[row+1], col-1)

		output = AppendStrToIntSlice(output, string(numStr))

		if end <= col && common.IsDigit(lr) {
			numStr, _, _ := GetWholeNumber(mat[row+1], col+1)

			output = AppendStrToIntSlice(output, string(numStr))
		}
	} else if common.IsDigit(lc) {
		numStr, _, _ := GetWholeNumber(mat[row+1], col)

		output = AppendStrToIntSlice(output, string(numStr))
	} else if common.IsDigit(lr) {
		numStr, _, _ := GetWholeNumber(mat[row+1], col+1)

		output = AppendStrToIntSlice(output, string(numStr))
	}

	// left side
	if common.IsDigit(cl) {
		numStr, _, _ := GetWholeNumber(mat[row], col-1)
		output = AppendStrToIntSlice(output, string(numStr))
	}

	//right side
	if common.IsDigit(cr) {
		numStr, _, _ := GetWholeNumber(mat[row], col+1)
		output = AppendStrToIntSlice(output, string(numStr))
	}

	return output
}

func Product(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	result := 1

	for _, n := range arr {
		result *= n
	}

	return result
}

func PartB(inputName string) {
	lines, err := common.FromFile(inputName)
	if err != nil {
		log.Fatal(err)
	}

	mat := GetRuneMatrix(lines)
	rows := len(mat)

	//PrintMatrix(mat)

	result := 0
	for i := 0; i < rows; i++ {
		cols := len(mat[i])
		for j := 0; j < cols; j++ {
			c := mat[i][j]
			if c == '*' {
				fmt.Println("Found asterisk at", i+1, j+1)

				nums := GetGearNumbers(mat, i, j, rows, cols)
				fmt.Println(nums)
				if len(nums) == 2 {
					fmt.Println("Asterisk is gear")

					ratio := Product(nums)
					fmt.Println("Gear ratio:", ratio)

					result += ratio
				}
			}
		}
	}

	fmt.Println("RESULT:", result)
}
