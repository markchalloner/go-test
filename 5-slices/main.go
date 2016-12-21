package main

import "fmt"

func main()  {
	slice1 := []int { 1, 2, 3, 4, 5 }
	slice2 := append(slice1, 6, 7)
	slice3 := slice2[1:len(slice2)-1]
	zeroSliceValues(&slice3)
	for i, value := range slice2 {
		fmt.Printf("index %d: value %d\n", i, value)
	}
}

func zeroSliceValues(slicePtr *[]int) {
	slice := *slicePtr
	for i := range slice {
		slice[i] = 0
	}
	*slicePtr = slice
}
