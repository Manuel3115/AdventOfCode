package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../puzzle_input.txt")

	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 50
	zeroCounts := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0:1]
		num, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Printf("The line number in the input file is wrong (%s).\n", line[1:])
			return
		}

		negativeBonus := 1
		if total == 0 {
			negativeBonus = 0
		}

		switch direction {
		case "R":
			total += num
		case "L":
			total -= num
		default:
			fmt.Printf("The line direction format in the input file is wrong (%s).\n", direction)
			return
		}

		if total <= 0 {
			zeroCounts += -total / 100 + negativeBonus
		} else {
			zeroCounts += total / 100
		}
		
		total = mod(total, 100)
	}

    fmt.Printf("The dial reached zero %d times.\n", zeroCounts)
}


func mod(dividend, modulus int) int {
	return (dividend % modulus + modulus) % modulus
}