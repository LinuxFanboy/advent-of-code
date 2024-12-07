package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	part1Sum := 0
	part2Sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		target, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("Error parsing target:", err)
			continue
		}

		numStrings := strings.Fields(parts[1])
		numbers := make([]int, len(numStrings))
		for i, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				continue
			}
			numbers[i] = num
		}

		if canProduceTargetPart1(numbers, target) {
			part1Sum += target
		}

		if canProduceTargetPart2(numbers, target) {
			part2Sum += target
		}
	}

	fmt.Println("Total calibration result (Part 1):", part1Sum)
	fmt.Println("Total calibration result (Part 2):", part2Sum)
}

func canProduceTargetPart1(numbers []int, target int) bool {
	if len(numbers) == 0 {
		return false
	}

	var helper func(idx int, current int) bool
	helper = func(idx int, current int) bool {
		if idx == len(numbers) {
			return current == target
		}

		if helper(idx+1, current+numbers[idx]) {
			return true
		}

		if helper(idx+1, current*numbers[idx]) {
			return true
		}

		return false
	}

	return helper(1, numbers[0])
}

func canProduceTargetPart2(numbers []int, target int) bool {
	if len(numbers) == 0 {
		return false
	}

	var helper func(idx int, current int) bool
	helper = func(idx int, current int) bool {
		if idx == len(numbers) {
			return current == target
		}

		if helper(idx+1, current+numbers[idx]) {
			return true
		}

		if helper(idx+1, current*numbers[idx]) {
			return true
		}

		concatenated, err := concatNumbers(current, numbers[idx])
		if err == nil && helper(idx+1, concatenated) {
			return true
		}

		return false
	}

  return helper(1, numbers[0])
}

func concatNumbers(a, b int) (int, error) {
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)
	resultStr := aStr + bStr
	result, err := strconv.Atoi(resultStr)
	if err != nil {
		return 0, err
	}
	return result, nil
}

