
### README for Advent od Code DAY 4 Ruby Card Matching Game Repository

#### Overview
This repository contains Ruby scripts designed to process and analyze card game data. The game involves a series of cards, each with a unique set of numbers, and a corresponding set of winning numbers. The scripts read input data, identify matches between card numbers and winning numbers, and calculate points based on these matches.

#### Contents
1. `input`: A text file containing the card data. Each line represents a card and its numbers, followed by winning numbers.
2. `part01.rb`: A Ruby script that reads the input file, calculates matches between card numbers and winning numbers, and computes total points.
3. `part02.rb`: Complementary Ruby script for processing card data, including functions for parsing input, counting matches, and additional data processing.

#### How to Run
- Ensure you have Ruby installed on your system.
- Place your card data in the `input` file.
- To run the scripts, use the following commands in your terminal:
  - `ruby part01.rb < input`
  - `ruby part02.rb < input`

#### Functionality
- **part01.rb**: Calculates total points by identifying matches between the card numbers and the winning numbers. The points are calculated based on the number of matches.
- **part02.rb**: Provides additional functions for handling card data, including parsing input and counting matches. It's structured to allow further expansion of the game's logic and data processing.

#### Requirements
- Ruby

Feel free to modify or extend the scripts as needed for your specific requirements. For any questions or contributions to the project, please open an issue or a pull request. 
