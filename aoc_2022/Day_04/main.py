with open("input") as file:
    lines = [i for i in file.read().splitlines()]

full_total = 0
partial_total = 0

for line in lines:
    section1, section2 = line.split(",")
    section1_begin, section1_end = section1.split("-")
    section2_begin, section2_end = section2.split("-")
    range_1 = [num for num in range(int(section1_begin), int(section1_end) + 1)]
    range_2 = [num for num in range(int(section2_begin), int(section2_end) + 1)]
    if all(x in range_1 for x in range_2) or all(x in range_2 for x in range_1):
        full_total += 1
    if any(x in range_1 for x in range_2):
        partial_total += 1

print(f"Complete Overlap {full_total}")
print(f"Partial Overlap {partial_total}")