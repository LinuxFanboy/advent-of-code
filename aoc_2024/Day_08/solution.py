from collections import defaultdict
from itertools import combinations


def parse_input(file_path):
    with open(file_path, 'r') as file:
        grid = [line.strip() for line in file if line.strip()]
    
    locations = defaultdict(list)
    for r, row in enumerate(grid):
        for c, char in enumerate(row):
            if char != '.':
                locations[char].append((r, c))
    return locations, len(grid), len(grid[0])

def calculate_antinodes_part1(locations, rows, cols):
    antinodes = set()
    for freq, positions in locations.items():
        for (x1, y1), (x2, y2) in combinations(positions, 2):
            dx, dy = x2 - x1, y2 - y1
            potential_antinodes = [
                (x1 - dx, y1 - dy),
                (x2 + dx, y2 + dy),
            ]
            for r, c in potential_antinodes:
                if 0 <= r < rows and 0 <= c < cols:
                    antinodes.add((r, c))
    return antinodes


def calculate_antinodes_part2(locations, rows, cols):
    antinodes = set()
    for freq, positions in locations.items():
        for (x1, y1), (x2, y2) in combinations(positions, 2):
            dx, dy = x2 - x1, y2 - y1

            row, col = x1, y1
            while 0 <= row < rows and 0 <= col < cols:
                antinodes.add((row, col))
                row -= dx
                col -= dy

            row, col = x2, y2
            while 0 <= row < rows and 0 <= col < cols:
                antinodes.add((row, col))
                row += dx
                col += dy
    return antinodes


def solve_puzzle(file_path):
    locations, rows, cols = parse_input(file_path)

    antinodes_part1 = calculate_antinodes_part1(locations, rows, cols)
    part1_result = len(antinodes_part1)

    antinodes_part2 = calculate_antinodes_part2(locations, rows, cols)
    part2_result = len(antinodes_part2)

    return part1_result, part2_result

if __name__ == "__main__":
    file_path = "input"
    part1_result, part2_result = solve_puzzle(file_path)
    print(f"Part 1 Solution: {part1_result}")
    print(f"Part 2 Solution: {part2_result}")
