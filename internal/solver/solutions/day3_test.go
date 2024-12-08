package solutions

import (
	"bufio"
	"strings"
	"testing"
)

func TestDay3_SolvePart1(t *testing.T) {
	day3 := Day3{}

	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	actual, err := day3.Part1(bufio.NewScanner(strings.NewReader(input)))
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	expected := 161
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestDay3_SolvePart2(t *testing.T) {
	day3 := Day3{}

	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	actual, err := day3.Part2(bufio.NewScanner(strings.NewReader(input)))
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	expected := 2
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}
