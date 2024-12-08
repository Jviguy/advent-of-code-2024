package solver

import (
	"bufio"
	"fmt"
	"github.com/jviguy/advent-of-code-2024/internal/solver/solutions"
	"os"
)

type AoCHandler struct {
	Solutions []Solution
}

func NewAoCHandler() *AoCHandler {
	s := []Solution{
		solutions.Day1{},
		solutions.Day2{},
		solutions.Day3{},
	}
	return &AoCHandler{
		Solutions: s,
	}
}

func (a *AoCHandler) SolveDay(day int, part int) (int, error) {
	if day < 1 || day > len(a.Solutions) {
		return 0, ErrDayNotFound
	}
	day-- // 0 based index
	solution := a.Solutions[day]

	// increase back to 1 based index for the file.
	file, err := os.Open(fmt.Sprintf("./data/day%d/input.txt", day+1))
	if err != nil {
		return 0, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	if part == 1 {
		return solution.Part1(scanner)
	} else {
		return solution.Part2(scanner)
	}
}
