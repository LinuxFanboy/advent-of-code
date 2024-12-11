package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var topographicMap [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var row []int
		for _, char := range line {
			num, _ := strconv.Atoi(string(char))
			row = append(row, num)
		}
		topographicMap = append(topographicMap, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
		return
	}

	totalScore := calculateTotalScore(topographicMap)
	fmt.Println("Part 1: Total Score of all trailheads:", totalScore)

	totalRating := calculateTotalRating(topographicMap)
	fmt.Println("Part 2: Total Rating of all trailheads:", totalRating)
}

func calculateTotalScore(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	totalScore := 0

	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == 0 {
				uniqueNines := findUniqueNines(grid, Point{x, y}, directions)
				totalScore += len(uniqueNines)
			}
		}
	}

	return totalScore
}

func findUniqueNines(grid [][]int, start Point, directions []Point) map[Point]bool {
	rows := len(grid)
	cols := len(grid[0])
	visited := make(map[Point]bool)
	uniqueNines := make(map[Point]bool)

	queue := []Point{start}
	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			next := Point{current.x + dir.x, current.y + dir.y}
			if next.x >= 0 && next.x < rows && next.y >= 0 && next.y < cols {
				if !visited[next] && grid[next.x][next.y] == grid[current.x][current.y]+1 {
					visited[next] = true
					queue = append(queue, next)
					if grid[next.x][next.y] == 9 {
						uniqueNines[next] = true
					}
				}
			}
		}
	}

	return uniqueNines
}

func calculateTotalRating(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	totalRating := 0

	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if grid[x][y] == 0 {
				rating := countDistinctTrails(grid, Point{x, y}, directions)
				totalRating += rating
			}
		}
	}

	return totalRating
}

func countDistinctTrails(grid [][]int, start Point, directions []Point) int {
	rows := len(grid)
	cols := len(grid[0])

	cache := make(map[Point]int)

	var dfs func(Point, int) int
	dfs = func(current Point, currentHeight int) int {
		if currentHeight == 9 {
			return 1
		}
		if val, exists := cache[current]; exists {
			return val
		}

		trailCount := 0
		for _, dir := range directions {
			next := Point{current.x + dir.x, current.y + dir.y}
			if next.x >= 0 && next.x < rows && next.y >= 0 && next.y < cols {
				if grid[next.x][next.y] == currentHeight+1 {
					trailCount += dfs(next, currentHeight+1)
				}
			}
		}

		cache[current] = trailCount
		return trailCount
	}

	return dfs(start, 0)
}

