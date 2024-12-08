package solutions

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type Day3 struct {
}

func (d Day3) Part1(scanner *bufio.Scanner) (int, error) {
	scanner.Split(bufio.ScanRunes)
	text := ""
	for scanner.Scan() {
		text += scanner.Text()
	}
	startIdx := -1
	result := 0
	for i := 0; i < len(text); i++ {
		if text[i] == '(' {
			startIdx = i
		} else if text[i] == ')' {
			if startIdx-3 >= 0 {
				multiRes, err := d.DoMulInstruction(text[startIdx-3 : i+1])
				if errors.Is(err, invalidInstruction) {
					continue
				} else if err != nil {
					return 0, nil
				}
				result += multiRes
			}
			startIdx = -1
		}
	}
	return result, nil
}

func (d Day3) Part2(scanner *bufio.Scanner) (int, error) {
	scanner.Split(bufio.ScanRunes)
	text := ""
	for scanner.Scan() {
		text += scanner.Text()
	}
	startIdx := -1
	result := 0
	enabled := true
	for i := 0; i < len(text); i++ {
		if i > 4 {
			if text[i-4:i] == "do()" {
				enabled = true
			}
		}
		if i > 7 {
			if text[i-7:i] == "don't()" {
				enabled = false
			}
		}
		if text[i] == '(' {
			startIdx = i
		} else if text[i] == ')' {
			if startIdx-3 >= 0 && enabled {
				multiRes, err := d.DoMulInstruction(text[startIdx-3 : i+1])
				if errors.Is(err, invalidInstruction) {
					continue
				} else if err != nil {
					return 0, nil
				}
				result += multiRes
			}
			startIdx = -1
		}
	}
	return result, nil
}

func (d Day3) DoMulInstruction(instruction string) (int, error) {
	if !strings.HasPrefix(instruction, "mul(") || !strings.HasSuffix(instruction, ")") {
		return 0, invalidInstruction
	}
	args := instruction[4 : len(instruction)-1]
	// args = "2,4"
	argSplit := strings.Split(args, ",")
	if len(argSplit) != 2 {
		return 0, invalidInstruction
	}
	// argSplit = ["2", "4"]
	arg1, err := strconv.Atoi(argSplit[0])
	if err != nil {
		return 0, invalidInstruction
	}
	arg2, err := strconv.Atoi(argSplit[1])
	if err != nil {
		return 0, invalidInstruction
	}
	return arg1 * arg2, nil
}

var invalidInstruction = errors.New("invalid instruction")
