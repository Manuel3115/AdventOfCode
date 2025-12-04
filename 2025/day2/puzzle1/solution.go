package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("../puzzle_input.txt")

	if err != nil {
		fmt.Printf("failed to open file: %v\n", err)
		return
	}

	intervals := strings.Split(string(fileContent), ",")

	invalidIdSum := 0

	for _, interval := range intervals {
		start, end, found := strings.Cut(interval, "-")

		if !found {
			fmt.Printf("Interval doesn't match correct format: %s\n", interval)
			return
		}

		endNum, _ := strconv.Atoi(end)

		halfSize := int(math.Ceil(float64(len(start)) / 2.0))
		halfSizeFloor := int(math.Floor(float64(len(start)) / 2.0))

		firstHalfStart := start[:halfSizeFloor]
		secondHalfStart := start[halfSizeFloor:]


		firstHalfMagnitude := int(math.Pow(10, float64(halfSize)))

		
		var firstHalfStartNum int

		if len(firstHalfStart) < len(secondHalfStart) {
			firstHalfStartNum = firstHalfMagnitude / 10	
		} else {
			firstHalfStartNum, _ = strconv.Atoi(firstHalfStart)
			secondHalfStartNum, _ := strconv.Atoi(secondHalfStart)

			if firstHalfStartNum < secondHalfStartNum {
				firstHalfStartNum += 1
			}
		}

		invalidId := firstHalfStartNum * firstHalfMagnitude + firstHalfStartNum

		for invalidId <= endNum {
			invalidIdSum += invalidId
			firstHalfStartNum += 1

			if firstHalfStartNum == firstHalfMagnitude {
				firstHalfMagnitude *= 10
			}

			invalidId = firstHalfStartNum * firstHalfMagnitude + firstHalfStartNum
			
		}
	}

	fmt.Printf("The total sum of all invalid ids are: %d\n", invalidIdSum)
}