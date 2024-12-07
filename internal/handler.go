package internal

import (
	"bufio"
	"fmt"
	"os"
)

type AoCHandler struct {
	Solutions []Solution
}

func NewAoCHandler() *AoCHandler {
	solutions := []Solution{}
	return &AoCHandler{
		Solutions: solutions,
	}
}

func (a *AoCHandler) SolveDay(day int, part int) (string, error) {
	if day < 0 || day >= len(a.Solutions) {
		return "", ErrDayNotFound
	}
	data, err := a.GetData(day)
	if err != nil {
		return "", err
	}
	if part == 1 {
		return a.Solutions[day].Part1(data), nil
	} else {
		return a.Solutions[day].Part2(data), nil
	}
}

func (a *AoCHandler) GetData(day int) (*bufio.Scanner, error) {
	if day < 0 || day >= len(a.Solutions) {
		return nil, ErrDayNotFound
	}
	file, err := os.Open(fmt.Sprintf("data/day%d.txt", day))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return bufio.NewScanner(file), nil
}
