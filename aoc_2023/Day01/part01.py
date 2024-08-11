import re

with open("input") as file:
    values = [line.rstrip() for line in file]

total = sum(int(re.sub('\\D', '', line)[0] + re.sub('\\D', '', line)[-1]) for line in values)

print("Output:", total)

