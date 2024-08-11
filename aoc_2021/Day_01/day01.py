with open('data') as f:
	data = [int(x) for x in f]

# Step1
print(sum(x < y for x, y in zip(data, data[1:])))

# Step2
print(sum(x < y for x, y in zip(data, data[3:])))
