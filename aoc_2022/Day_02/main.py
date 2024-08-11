with open('input', 'r') as f:
    games = f.read().split('\n')

WIN = 6
DRAW = 3
LOSS = 0

ROCK = 1
PAPER = 2
SCISSORS = 3

game_scores = {
    'A X': DRAW + ROCK,
    'A Y': WIN + PAPER,
    'A Z': LOSS + SCISSORS,
    'B X': LOSS + ROCK,
    'B Y': DRAW + PAPER,
    'B Z': WIN + SCISSORS,
    'C X': WIN + ROCK,
    'C Y': LOSS + PAPER,
    'C Z': DRAW + SCISSORS,
}

print(f'Part 1: {sum([game_scores.get(game, 0) for game in games])}')

game_scores = {
    'A X': LOSS + SCISSORS,
    'A Y': DRAW + ROCK,
    'A Z': WIN + PAPER,
    'B X': LOSS + ROCK,
    'B Y': DRAW + PAPER,
    'B Z': WIN + SCISSORS,
    'C X': LOSS + PAPER,
    'C Y': DRAW + SCISSORS,
    'C Z': WIN + ROCK,
}

print(f'Part 2: {sum([game_scores.get(game, 0) for game in games])}')
