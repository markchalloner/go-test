package main

import "fmt"
import "time"

func main() {
	var timeChannel chan string = make(chan string)

	emitTime := createTimeEmitter()

	go emitTime(timeChannel)
	go speak(timeChannel)

	var input string
	fmt.Scanln(&input)
}

func speak(timeChannel chan string) {
	for {
		timeString := <- timeChannel
		fmt.Printf("The time is %s\n...", timeString)
	}
}

func createTimeEmitter() (func(chan string)) {
	ticker := time.NewTicker(1 * time.Second)
	prev := time.Now().Add(-time.Minute);
	return func(timeChannel chan string) {
		for {
			now := <- ticker.C
			if now.Minute() != prev.Minute() {
				timeChannel <- now.Format("15:04")
				prev = now
			}
		}
	}
}
