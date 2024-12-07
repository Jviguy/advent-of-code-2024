package solver

import "bufio"

// Solution is a struct that contains the solution to an AoC problem
type Solution interface {
	Part1(scanner *bufio.Scanner) (int, error)
	Part2(scanner *bufio.Scanner) (int, error)
}
