package main

import (
	"bufio"
	"fmt"
	"log"
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

	var table [][]byte;

	for scanner.Scan() {
		table = append(table, []byte(scanner.Text()))
	}

	totalRemovedRolls := 0
	table, removedRolls := removeRollsIteration(table)
	totalRemovedRolls += removedRolls

	for removedRolls > 0 {
		removedRolls = 0
		table, removedRolls = removeRollsIteration(table)
		totalRemovedRolls += removedRolls
	}

	file, err = os.Create("output.txt")

	if err != nil {
		log.Printf("failed to create file: %v\n", err)
		return
	}

	defer file.Close()

	for i := range table {
		file.Write(table[i])
		file.WriteString("\n")
	}

	fmt.Printf("There are %d accessible rolls.\n", totalRemovedRolls)
}

func removeRollsIteration(table [][]byte) ([][]byte, int) {
	removedRolls := 0
	for i := range table {
		for j := range table[i] {
			if table[i][j] == '@' && getSurroundingRolls(table, i, j) < 4 {
				table[i][j] = '.'
				removedRolls += 1
			}
		}
	}
	return table, removedRolls
}

func getSurroundingRolls(table [][]byte, x, y int) int {
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