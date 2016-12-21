package main

import "fmt"

func main() {
	myMap := map[string]int { "a": 1, "b": 2 }
	var key string

	fmt.Print("Enter a key name: ")
	fmt.Scanf("%s", &key)
	fmt.Println()

	if _, exists := myMap[key]; exists {
		fmt.Printf("The key %s exists\n\n", key)
	} else {
		fmt.Printf("The key %s does not exist\n\n", key)
	}

	fmt.Println("All keys/values:\nabc")
	for i, value := range myMap {
		fmt.Printf("index %s: value %d\n", i, value)
	}
}
