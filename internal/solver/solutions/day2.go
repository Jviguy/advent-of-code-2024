package solutions

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

type Day2 struct{}

func (d Day2) Part1(scanner *bufio.Scanner) (int, error) {
	result := 0
	for scanner.Scan() {
		list := scanner.Text()
		listScanner := bufio.NewScanner(strings.NewReader(list))
		listScanner.Split(bufio.ScanWords)
		safe := true
		previous := -1
		decreasing := false
		foundDirection := false
		for listScanner.Scan() {
			currentStr := listScanner.Text()
			current, err := strconv.Atoi(currentStr)
			if err != nil {
				return 0, err
			}
			if previous != -1 {
				distance := current - previous
				if !foundDirection {
					decreasing = distance < 0
					foundDirection = true
				}
				safe = !(distance == 0 || math.Abs(float64(distance)) > 3 ||
					(decreasing && distance > 0) || (!decreasing && distance < 0))
				if !safe {
					break
				}

			}
			previous = current
		}
		if safe {
			result++
		}
	}
	return result, nil
}

func (d Day2) Part2(scanner *bufio.Scanner) (int, error) {
	result := 0
	for scanner.Scan() {
		array := make([]int, 0)
		list := scanner.Text()
		listScanner := bufio.NewScanner(strings.NewReader(list))
		listScanner.Split(bufio.ScanWords)
		for listScanner.Scan() {
			currentStr := listScanner.Text()
			current, err := strconv.Atoi(currentStr)
			if err != nil {
				return 0, err
			}
			array = append(array, current)
		}
		if isSafe(array) {
			result++
			continue
		}
		for i := 0; i < len(array); i++ {
			var tempArr []int
			for j := 0; j < i; j++ {
				tempArr = append(tempArr, array[j])
			}
			for j := i + 1; j < len(array); j++ {
				tempArr = append(tempArr, array[j])
			}
			if isSafe(tempArr) {
				result++
				break
			}
		}
	}
	return result, nil
}

func isSafe(array []int) bool {
	previous := -1
	decreasing := false
	foundDirection := false
	for _, current := range array {
		if previous != -1 {
			distance := current - previous
			if !foundDirection {
				decreasing = distance < 0
				foundDirection = true
			}
			if distance == 0 || math.Abs(float64(distance)) > 3 ||
				(decreasing && distance > 0) || (!decreasing && distance < 0) {
				return false
			}
		}
		previous = current
	}
	return true
}
