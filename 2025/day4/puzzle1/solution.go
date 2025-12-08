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

	var table []string;

	for scanner.Scan() {
		table = append(table, scanner.Text())
	}

	accessibleRolls := 0

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == '@' && getSurroundingRolls(table, i, j) < 4 {
				accessibleRolls += 1
			}
		}
	}

	fmt.Printf("There are %d accessible rolls.\n", accessibleRolls)
}

func getSurroundingRolls(table []string, x, y int) int {
	surroundingRolls := 0
	for i := -1; i <= 1; i++ {
		if x + i >= len(table) || x + i < 0 {
			continue
		}
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if y + j >= len(table[x + i]) || y + j < 0 {
				continue
			}
			if table[x + i][y + j] == '@' {
				surroundingRolls += 1
			}
		}
	}
	return surroundingRolls
}