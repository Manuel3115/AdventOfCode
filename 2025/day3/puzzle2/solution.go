package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	DigitCount = 12
)

func main() {
	file, err := os.Open("../puzzle_input.txt")

	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalSum := 0;

	for scanner.Scan() {
		line := scanner.Text()

		digits := make([]int, DigitCount)

		for i := 0; i < len(line); i++ {
			digit := int(line[i] - '0')
			for j := range DigitCount {
				if digits[j] < digit && i < len(line) - (DigitCount - j - 1) {
					digits[j] = digit
					for k := j + 1; k < DigitCount; k++ {
						digits[k] = 0
					}
					break
				}
			}
		}

		for i := range DigitCount {
			totalSum += digits[DigitCount - i - 1] * int(math.Pow(10, float64(i)))
		}
	}

	fmt.Printf("The total sum of the largest possible joltage is : %d\n", totalSum)
}

