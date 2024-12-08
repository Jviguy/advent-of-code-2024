package solutions

import (
	"bufio"
	"strings"
	"testing"
)

func TestDay2_Part1(t *testing.T) {
	day2 := Day2{}

	input := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"

	actual, err := day2.Part1(bufio.NewScanner(strings.NewReader(input)))
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	expected := 2
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestDay2_Part2(t *testing.T) {
	day2 := Day2{}
	input := "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"

	actual, err := day2.Part2(bufio.NewScanner(strings.NewReader(input)))
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	expected := 4
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}
