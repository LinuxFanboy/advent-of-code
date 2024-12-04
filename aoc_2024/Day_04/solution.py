def find_word(grid, y, x, word):
    found = 0
    for dy, dx in ((-1, 0), (-1, 1), (0, 1), (1, 1), (1, 0), (1, -1), (0, -1), (-1, -1)):
        cur = ""
        for i in range(len(word)):
            ny = y + i * dy
            nx = x + i * dx
            if ny not in range(len(grid)) or nx not in range(len(grid[ny])):
                break
            cur += grid[ny][nx]
        if cur == word:
            found += 1
    return found

def count_x_mas_in_grid(grid):
    rows, cols = len(grid), len(grid[0])
    count = 0
    for y in range(1, rows - 1):
        for x in range(1, cols - 1):
            if grid[y][x] == "A":
                letter_pairs = [
                    grid[y - 1][x - 1] + grid[y + 1][x + 1],
                    grid[y + 1][x - 1] + grid[y - 1][x + 1]
                ]
                if all(pair in ("MS", "SM") for pair in letter_pairs):
                    count += 1
    return count

if __name__ == "__main__":
    import sys
    if len(sys.argv) != 2:
        sys.exit(1)
    input_file = sys.argv[1]
    with open(input_file, 'r') as file:
        grid = file.read().splitlines()
    found1 = sum(find_word(grid, y, x, "XMAS") for y, row in enumerate(grid) for x, letter in enumerate(row))
    found2 = count_x_mas_in_grid(grid)
    print(found1)
    print(found2)

