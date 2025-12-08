package main

import (
	"bufio"
	"fmt"
	"os"
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
		firstDigit := 0;
		secondDigit := 0;
		for i := 0; i < len(line); i++ {
			digit := int(line[i] - '0')
			if firstDigit < digit && i != len(line) - 1 {
				firstDigit = digit
				secondDigit = 0
			} else if secondDigit < digit {
				secondDigit = digit
			}
		}

		totalSum += firstDigit * 10 + secondDigit
	}

	fmt.Printf("The total sum of the largest possible joltage is : %d\n", totalSum)
}