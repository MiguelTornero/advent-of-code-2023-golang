package common

import (
	"bufio"
	"os"
)

func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func IsLowecaseLetter(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func FromFile(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	output := []string{}
	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanLines)
	for reader.Scan() {
		output = append(output, reader.Text())
	}

	return output, nil
}

func Sum(nums []int) int {
	result := 0

	for _, num := range nums {
		result += num
	}

	return result
}

func Hello() string {
	return "world"
}

type Queue[T any] struct {
	elements []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		elements: []T{},
	}
}

func (q *Queue[T]) Pop() T {
	top, new := q.elements[0], q.elements[1:]
	q.elements = new
	return top
}

func (q *Queue[T]) Push(elem T) {
	q.elements = append(q.elements, elem)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) <= 0
}
