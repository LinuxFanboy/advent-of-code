package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Game struct {
	grid          [][]rune
	startingPoint Coordinate
}

type Coordinate struct {
	x, y int
}

func main() {
	PartOneAndTwo()
}

func PartOneAndTwo() {
	game := loadData()

	// Part 1: Number of distinct positions visited
	distinctVisited := simulateMovement(game)
	fmt.Println("The answer to part 1 is:", distinctVisited)

	// Part 2: Number of positions where adding an obstruction causes a loop
	loopPositions := findLoopObstructionPositions(game)
	fmt.Println("The answer to part 2 is:", loopPositions)
}

func loadData() Game {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	game := Game{}
	var grid [][]rune
	var expectedRowLength int

	for rowIndex := 0; scanner.Scan(); rowIndex++ {
		line := scanner.Text()
		line = trimWhitespace(line)
		if len(line) == 0 {
			continue // Skip empty lines
		}

		row := []rune(line)
		if expectedRowLength == 0 {
			expectedRowLength = len(row) // Set expected row length from the first valid row
		}
		if len(row) != expectedRowLength {
			fmt.Printf("Row %d has inconsistent length: %d (expected %d)\n", rowIndex, len(row), expectedRowLength)
			panic("Inconsistent row length in the grid!")
		}

		for colIndex, cell := range row {
			if cell == '^' {
				game.startingPoint = Coordinate{x: rowIndex, y: colIndex}
				row[colIndex] = '.' // Clear guard's starting marker
			}
		}
		grid = append(grid, row)
	}
	game.grid = grid
	return game
}

func trimWhitespace(input string) string {
	return strings.TrimSpace(input)
}

func simulateMovement(game Game) int {
	directions := []Coordinate{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // UP, RIGHT, DOWN, LEFT
	currentDir := 0 // Start facing UP
	currentPos := game.startingPoint
	visited := map[Coordinate]bool{currentPos: true}

	for {
		// Determine the next position
		nextPos := Coordinate{
			x: currentPos.x + directions[currentDir].x,
			y: currentPos.y + directions[currentDir].y,
		}

		// Check boundaries and obstacles
		if !inBounds(game.grid, nextPos.x, nextPos.y) {
			break
		}
		if game.grid[nextPos.x][nextPos.y] == '#' {
			// Turn right if an obstacle is ahead
			currentDir = (currentDir + 1) % 4
		} else {
			// Move forward
			currentPos = nextPos
			visited[currentPos] = true
		}
	}

	return len(visited)
}

func findLoopObstructionPositions(game Game) int {
	directions := []Coordinate{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // UP, RIGHT, DOWN, LEFT
	startingPosition := game.startingPoint
	grid := game.grid
	visited := getVisited(grid, startingPosition)

	loopPositions := 0

	// Test each visited position as a potential obstruction
	for position := range visited {
		// Skip the starting position
		if position == startingPosition {
			continue
		}

		// Temporarily place an obstruction
		grid[position.x][position.y] = '#'

		// Check if this causes a loop
		if causesLoop(grid, startingPosition, directions) {
			loopPositions++
		}

		// Remove the obstruction
		grid[position.x][position.y] = '.'
	}

	return loopPositions
}

func getVisited(grid [][]rune, start Coordinate) map[Coordinate]bool {
	directions := []Coordinate{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // UP, RIGHT, DOWN, LEFT
	currentDir := 0 // Start facing UP
	currentPos := start
	visited := map[Coordinate]bool{currentPos: true}

	for {
		// Determine the next position
		nextPos := Coordinate{
			x: currentPos.x + directions[currentDir].x,
			y: currentPos.y + directions[currentDir].y,
		}

		// Check boundaries and obstacles
		if !inBounds(grid, nextPos.x, nextPos.y) {
			break
		}
		if grid[nextPos.x][nextPos.y] == '#' {
			// Turn right if an obstacle is ahead
			currentDir = (currentDir + 1) % 4
		} else {
			// Move forward
			currentPos = nextPos
			visited[currentPos] = true
		}
	}

	return visited
}

func causesLoop(grid [][]rune, start Coordinate, directions []Coordinate) bool {
	currentDir := 0 // Start facing UP
	currentPos := start
	visited := map[Coordinate]int{currentPos: 1}

	for {
		// Determine the next position
		nextPos := Coordinate{
			x: currentPos.x + directions[currentDir].x,
			y: currentPos.y + directions[currentDir].y,
		}

		// Check boundaries and obstacles
		if !inBounds(grid, nextPos.x, nextPos.y) {
			return false // Exit if out of bounds
		}
		if grid[nextPos.x][nextPos.y] == '#' {
			// Turn right if an obstacle is ahead
			currentDir = (currentDir + 1) % 4
		} else {
			// Move forward
			currentPos = nextPos
			visited[currentPos]++

			// If a position is visited 5 times, a loop is detected
			if visited[currentPos] >= 5 {
				return true
			}
		}
	}
}

func inBounds(grid [][]rune, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0])
}

