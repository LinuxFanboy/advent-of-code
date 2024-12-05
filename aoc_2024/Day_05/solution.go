package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct{ X, Y int }

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := parseRules(scanner)
	updates := parseUpdates(scanner)

	part1, part2 := processUpdates(updates, rules)

	fmt.Println("Sum of middle page numbers (Part 1):", part1)
	fmt.Println("Sum of middle page numbers (Part 2):", part2)
}

func parseRules(scanner *bufio.Scanner) []Rule {
	rules := []Rule{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "|")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		rules = append(rules, Rule{X: x, Y: y})
	}
	return rules
}

func parseUpdates(scanner *bufio.Scanner) [][]int {
	updates := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		update := make([]int, len(parts))
		for i, part := range parts {
			update[i], _ = strconv.Atoi(part)
		}
		updates = append(updates, update)
	}
	return updates
}

func processUpdates(updates [][]int, rules []Rule) (int, int) {
	part1, part2 := 0, 0
	incorrect := [][]int{}

	for _, update := range updates {
		if isOrdered(update, rules) {
			part1 += middle(update)
		} else {
			incorrect = append(incorrect, update)
		}
	}

	for _, update := range incorrect {
		part2 += middle(correctOrder(update, rules))
	}
	return part1, part2
}

func isOrdered(update []int, rules []Rule) bool {
	pageIndex := make(map[int]int)
	for i, page := range update {
		pageIndex[page] = i
	}
	for _, rule := range rules {
		if x, okX := pageIndex[rule.X]; okX {
			if y, okY := pageIndex[rule.Y]; okY && x > y {
				return false
			}
		}
	}
	return true
}

func correctOrder(update []int, rules []Rule) []int {
	graph, indegree := buildGraph(update, rules)
	queue, sorted := []int{}, []int{}

	for page, degree := range indegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		sorted = append(sorted, cur)
		for _, neighbor := range graph[cur] {
			if indegree[neighbor]--; indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	return sorted
}

func buildGraph(update []int, rules []Rule) (map[int][]int, map[int]int) {
	graph := make(map[int][]int)
	indegree := make(map[int]int)
	for _, page := range update {
		graph[page] = []int{}
		indegree[page] = 0
	}
	for _, rule := range rules {
		if contains(update, rule.X) && contains(update, rule.Y) {
			graph[rule.X] = append(graph[rule.X], rule.Y)
			indegree[rule.Y]++
		}
	}
	return graph, indegree
}

func middle(slice []int) int {
	return slice[len(slice)/2]
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

