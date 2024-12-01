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
		panic(err)
	}
	defer file.Close()

	var leftList, rightList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		if len(tokens) == 2 {
			left, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic(err)
			}
			leftList = append(leftList, left)
			rightList = append(rightList, right)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	rightCount := make(map[int]int)
	for _, num := range rightList {
		rightCount[num]++
	}

	similarityScore := 0
	for _, num := range leftList {
		count := rightCount[num]
		similarityScore += num * count
	}

	fmt.Printf("Similarity Score: %d\n", similarityScore)
}
