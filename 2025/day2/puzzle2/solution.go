package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Divisor struct {
	a int
	b int
}


func main() {
	primeDivisorMap := map[int][]Divisor{
		1: {},
		2: {{2, 1}},
		3: {{3, 1}},
		4: {{4, 1}, {2, 2}},
		5: {{5, 1}},
		6: {{6, 1}, {3, 2}, {2, 3}},
		7: {{7, 1}},
		8: {{8, 1}, {2, 4}},
		9: {{9, 1}, {3, 3}},
		10: {{10, 1}, {2, 5}, {5, 2}},
		11: {{11, 1}},
		12: {{12, 1}, {2, 6}, {3, 4}},
		13: {{13, 1}},
		14: {{14, 1}, {7, 2}},
		15: {{15, 1}, {5, 3}, {3, 5}},
	}


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

		startNum, _ := strconv.Atoi(start)
		endNum, _ := strconv.Atoi(end)

		magnitude := len(start)

		for magnitude <= len(end) {
			divisors := primeDivisorMap[magnitude]
			var firstSum int
			for i, divisor := range divisors {
				geometricSeriesResult := calculateGeometricSeries(divisor.a, divisor.b)
				minimumInvalidId := math.Ceil(math.Max(float64(startNum), math.Pow(10, float64(magnitude - 1))) / geometricSeriesResult)
				maximumInvalidId := math.Floor(math.Min(float64(endNum), math.Pow(10, float64(magnitude)) - 1) / geometricSeriesResult)

				if minimumInvalidId > maximumInvalidId {
					continue
				}

				additionalSum := int(((minimumInvalidId + maximumInvalidId) * (maximumInvalidId - minimumInvalidId + 1) / 2) * geometricSeriesResult)
				if (i == 0) {
					firstSum = additionalSum
				} else {
					additionalSum -= firstSum
				}
				
				invalidIdSum += additionalSum
			}
			magnitude += 1
		}
	}

	fmt.Printf("The total sum of all invalid ids are: %d\n", invalidIdSum)
}

func calculateGeometricSeries(a, b int) float64 {
	return ((math.Pow(10, float64(a * b)) - 1) / (math.Pow(10, float64(b)) - 1))
}
