package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	maxWidth := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := range grid {
		for len(grid[i]) < maxWidth {
			grid[i] = append(grid[i], ' ')
		}
	}

	height := len(grid)
	width := maxWidth

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var totalPrice int

	var floodFill func(x, y int, plant rune) (int, int)
	floodFill = func(x, y int, plant rune) (int, int) {
		stack := []Point{{x, y}}
		area := 0
		perimeter := 0

		for len(stack) > 0 {
			point := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if visited[point.x][point.y] {
				continue
			}
			visited[point.x][point.y] = true
			area++

			for _, d := range directions {
				nx, ny := point.x+d.x, point.y+d.y
				if nx >= 0 && nx < height && ny >= 0 && ny < width {
					if grid[nx][ny] == plant {
						if !visited[nx][ny] {
							stack = append(stack, Point{nx, ny})
						}
					} else {
						perimeter++
					}
				} else {
					perimeter++
				}
			}
		}

		return area, perimeter
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if !visited[i][j] && grid[i][j] != ' ' {
				plant := grid[i][j]
				area, perimeter := floodFill(i, j, plant)
				totalPrice += area * perimeter
			}
		}
	}

	fmt.Println(totalPrice)
}

