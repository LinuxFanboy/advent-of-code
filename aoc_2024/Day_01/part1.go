package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		diff := leftList[i] - rightList[i]
		if diff < 0 {
			diff = -diff
		}
		totalDistance += diff
	}

	fmt.Printf("Total Distance: %d\n", totalDistance)
}
