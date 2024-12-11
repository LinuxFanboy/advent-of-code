package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func blinkNTimes(iterations int) int {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}
	input := strings.TrimSpace(string(data))

	stoneStrs := strings.Fields(input)
	stones := make(map[int]int)
	for _, s := range stoneStrs {
		num, _ := strconv.Atoi(s)
		stones[num]++
	}

	for i := 0; i < iterations; i++ {
		newStones := make(map[int]int)
		for rock, count := range stones {
			blinkResults := blink(rock)
			for _, blinkResult := range blinkResults {
				newStones[blinkResult] += count
			}
		}
		stones = newStones
	}

	total := 0
	for _, count := range stones {
		total += count
	}

	return total
}

func blink(stone int) []int {
	if stone == 0 {
		return []int{1}
	} else if len(strconv.Itoa(stone))%2 == 0 {
		left, right := splitNumber(stone)
		return []int{left, right}
	} else {
		return []int{stone * 2024}
	}
}

func splitNumber(num int) (int, int) {
	s := strconv.Itoa(num)
	mid := len(s) / 2
	left, _ := strconv.Atoi(s[:mid])
	right, _ := strconv.Atoi(s[mid:])
	return left, right
}

func main() {
	resultPart1 := blinkNTimes(25)
	fmt.Println("Number of stones after 25 blinks:", resultPart1)

	resultPart2 := blinkNTimes(75)
	fmt.Println("Number of stones after 75 blinks:", resultPart2)
}

