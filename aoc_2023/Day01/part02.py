import re

replacements_dictionary = {
    'one': 'on1e', 'two': 'tw2o', 'three': 'thr3e', 'four': 'fo4ur', 
    'five': 'fi5ve', 'six': 'si6x', 'seven': 'sev7en', 'eight': 'ei8ght', 
    'nine': 'ni9ne'
}

with open("input") as file:
    values = [line.rstrip() for line in file]

total = 0
for line in values:
    for key, value in replacements_dictionary.items():
        line = line.replace(key, value)
    
    digits_only = re.sub('\\D', '', line)
    if digits_only:
        total += int(digits_only[0] + digits_only[-1])

print("Output:", total)
