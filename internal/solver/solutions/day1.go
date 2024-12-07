package solutions

import (
	"bufio"
	"errors"
	"github.com/jviguy/advent-of-code-2024/internal/solver/dsa"
	"math"
	"strconv"
)

type Day1 struct{}

func (d Day1) Part1(scanner *bufio.Scanner) (int, error) {
	left := dsa.NewMinHeap()
	right := dsa.NewMinHeap()
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		first, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		if !scanner.Scan() {
			return 0, errors.New("expected two values on each line")
		}

		second, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}
		left.Insert(first)
		right.Insert(second)
	}
	// The idea is to have two heaps, one for the left side and one for the right side.
	// the answer is the distance between the k-th smallest value of the two different lists, left and right.
	result := 0

	// Iterate over all values in the left heap. Left and Right heap are the same size
	size := left.Size()
	for k := 0; k < size; k++ {
		// Remove the smallest value from the left heap
		leftValue := left.ExtractMin()
		// Remove the smallest value from the right heap
		rightValue := right.ExtractMin()
		// Add the difference between the two values to the result
		result += int(math.Abs(float64(rightValue - leftValue)))
	}
	return result, nil
}

func (d Day1) Part2(scanner *bufio.Scanner) (int, error) {
	left := make([]int, 0)
	right := make(map[int]int)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		first, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		if !scanner.Scan() {
			return 0, errors.New("expected two values on each line")
		}

		second, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		left = append(left, first)

		if _, ok := right[second]; ok {
			right[second]++
		} else {
			right[second] = 1
		}
	}
	// The idea is to have a list of all the values on the left.
	// And a count of all the values on the right stored in a map.
	result := 0
	for _, value := range left {
		count := 0
		if c, ok := right[value]; ok {
			count = c
		}
		result += value * count
	}
	return result, nil
}
