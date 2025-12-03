package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../puzzle_input.txt")

	if err != nil {
		log.Fatalf("failed to open file: %v", err)
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

		switch direction {
		case "R":
			total = (total + num) % 100
		case "L":
			total = (total - num) % 100
		default:
			fmt.Printf("The line direction format in the input file is wrong (%s).\n", direction)
			return
		}

		if total == 0 {
			zeroCounts += 1
		}
	}

    fmt.Printf("The dial reached zero %d times.\n", zeroCounts)
}
