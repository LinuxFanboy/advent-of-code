
# Python Script Overview (Part 01)

## Description
This script processes data from an "input" file. It reads each line, extracts numeric values using regular expressions, and sums these values to calculate a total.

## Key Functionalities

1. **Reading Data from File**
   ```python
   with open("input") as file:
       values = [line.rstrip() for line in file]
   ```
   - Reads lines from the 'input' file, removing trailing characters.

2. **Extracting Numeric Values and Summing**
   ```python
   total = sum(int(re.sub('\D', '', line)[0] + re.sub('\D', '', line)[-1]) for line in values)
   ```
   - Uses regular expressions to extract numeric values from each line and calculates their sum.

3. **Outputting Total**
   ```python
   print("Output:", total)
   ```
   - Prints the calculated total.

## Usage Instructions
- Place the data file named 'input' in the same directory as the script.
- Run the script using a Python interpreter.


---


# Python Script Overview (Part 02)

## Description
This script reads from an "input" file, performs replacements based on a dictionary, and then calculates a total. It transforms specific words into a special format and sums up numeric values from the transformed data.

## Key Functionalities

1. **Dictionary for Replacements**
   ```python
   replacements_dictionary = { ... }
   ```
   - Defines a dictionary for word-to-word replacements.

2. **Reading and Transforming Data**
   ```python
   for line in values:
       for key, value in replacements_dictionary.items():
           line = line.replace(key, value)
   ```
   - Reads data from file and applies replacements based on the dictionary.

3. **Calculating Total**
   ```python
   total = 0
   for line in values:
       digits_only = re.sub('\D', '', line)
       if digits_only:
           total += int(digits_only)
   ```
   - Extracts numeric values from transformed data and calculates their sum.

## Usage Instructions
- Ensure a data file named 'input' is in the same directory as the script.
- Execute the script with a Python interpreter.
