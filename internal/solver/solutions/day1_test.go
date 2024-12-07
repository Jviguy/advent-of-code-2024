package solutions

import (
	"bufio"
	"strings"
	"testing"
)

func TestDay1_Part1(t *testing.T) {
	day1 := Day1{}

	input := "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"
	actual, err := day1.Part1(bufio.NewScanner(strings.NewReader(input)))
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	expected := 11
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestDay1_Part2(t *testing.T) {
	day1 := Day1{}

	input := "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"
	actual, err := day1.Part2(bufio.NewScanner(strings.NewReader(input)))
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	expected := 31
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}
