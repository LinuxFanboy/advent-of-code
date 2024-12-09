package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	filePath := "input"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("failed to read input file: %v", err))
	}
	input := strings.TrimSpace(string(data))

	fmt.Println("Part 1 Checksum:", part1(input))
	fmt.Println("Part 2 Checksum:", part2(input))
}

func part1(input string) int {
	diskMap := parseInput(input)

	freeSpacePtr, fileBlockPtr := 0, len(diskMap)-1

	for freeSpacePtr < fileBlockPtr {
		for fileBlockPtr >= 0 && diskMap[fileBlockPtr] == -1 {
			fileBlockPtr--
		}
		for freeSpacePtr < len(diskMap) && diskMap[freeSpacePtr] != -1 {
			freeSpacePtr++
		}
		if freeSpacePtr < fileBlockPtr {
			diskMap[freeSpacePtr], diskMap[fileBlockPtr] = diskMap[fileBlockPtr], -1
		}
	}

	return computeChecksum(diskMap)
}

func part2(input string) int {
	diskMap := parseInput(input)
	upperLimit := len(diskMap) - 1

	for {
		fileBlockStart, fileBlockEnd, fileID, found := findPreviousFileBlock(diskMap, upperLimit)
		if !found {
			break
		}

		fileBlockLen := fileBlockEnd - fileBlockStart

		freeSpaceStart, found := findContiguousFreeSpace(diskMap, fileBlockLen, fileBlockStart)
		if found {
			for i := 0; i < fileBlockLen; i++ {
				diskMap[freeSpaceStart+i] = fileID
				diskMap[fileBlockStart+i] = -1
			}
		}

		upperLimit = fileBlockStart - 1
	}

	return computeChecksum(diskMap)
}

func parseInput(input string) []int {
	var diskMap []int
	fileID := 0

	for i, ch := range input {
		length := int(ch - '0')
		if i%2 == 0 {
			for j := 0; j < length; j++ {
				diskMap = append(diskMap, fileID)
			}
			fileID++
		} else {
			for j := 0; j < length; j++ {
				diskMap = append(diskMap, -1)
			}
		}
	}
	return diskMap
}

func computeChecksum(diskMap []int) int {
	checksum := 0
	for pos, fileID := range diskMap {
		if fileID != -1 {
			checksum += pos * fileID
		}
	}
	return checksum
}

func findPreviousFileBlock(diskMap []int, limit int) (int, int, int, bool) {
	for limit >= 0 && diskMap[limit] == -1 {
		limit--
	}
	if limit < 0 {
		return -1, -1, -1, false
	}
	fileID := diskMap[limit]
	start := limit
	for start >= 0 && diskMap[start] == fileID {
		start--
	}
	return start + 1, limit + 1, fileID, true
}

func findContiguousFreeSpace(diskMap []int, size, limit int) (int, bool) {
	start, end := 0, 0
	for end-start < size && start < limit {
		for start < limit && diskMap[start] != -1 {
			start++
		}
		if start >= limit {
			break
		}
		end = start
		for end < limit && diskMap[end] == -1 {
			end++
		}
		if end-start >= size {
			return start, true
		}
		start = end
	}
	return -1, false
}
