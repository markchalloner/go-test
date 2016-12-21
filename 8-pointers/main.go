package main

import "fmt"

func main() {
	number := 1
	numberPtr1 := &number
	numberPtr2 := new(int)
	*numberPtr2 = 1

	addOne(numberPtr1)
	addOne(numberPtr1)
	addOne(numberPtr1)

	addOne(numberPtr2)

	swap(&number, numberPtr2)

	fmt.Println(number)
	fmt.Println(*numberPtr2)
}

func addOne(numberPtr *int) {
	*numberPtr++
}

func swap(numberPtr1 *int, numberPtr2 *int) {
	// XOR swap for kicks
	*numberPtr1 = *numberPtr1 ^ *numberPtr2
	*numberPtr2 = *numberPtr1 ^ *numberPtr2
	*numberPtr1 = *numberPtr1 ^ *numberPtr2
}

