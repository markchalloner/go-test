package main

import "fmt"

func main() {
	numbers := []float64 {6, 2, 1, 5, 5, 5}
	mean, mode, median := average(numbers)
	fmt.Printf("Mean: %.1f\n", mean)
	fmt.Printf("Mode: %.1f\n", mode)
	fmt.Printf("Median: %.1f\n", median)
	fmt.Printf("Total: %.1f\n", add(numbers...))
}

func add(args ...float64) float64 {
	total := 0.0
	for _, value := range args {
		total += value
	}
	return total
}

func average(numbers []float64) (float64, float64, float64) {
	total := 0.0
	numbersLen := len(numbers)

	// Calculate mean
	for i := 0; i < numbersLen; i++ {
		total += numbers[i]
	}
	mean := total / float64(numbersLen)

	// Calculate mode
	mode := 0.0
	instances := make(map[float64]int)
	for i := 0; i < numbersLen; i++ {
		value := numbers[i]
		if _, exists := instances[value]; exists {
			instances[value]++
		} else {
			instances[value] = 1
		}
		if instances[value] > instances[mode] {
			mode = value
		}
	}

	// Calculate median
	var median float64
	if numbersLen % 2 == 0 {
		median = (numbers[numbersLen / 2 - 1] + numbers[numbersLen / 2 + 1]) / 2
	} else {
		median = numbers[numbersLen / 2]
	}

	return mean, mode, median
}
