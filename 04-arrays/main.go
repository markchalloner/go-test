package main

import "fmt"

func main() {
	x := [5]float32 { 1, 2, 3, 4, 5 }
	// Set
	for i, value := range x {
		x[i] += value + float32(i) / 2
	}
	// Get
	for _, value := range x {
		fmt.Println(value)
	}
}
