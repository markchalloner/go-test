package main

import "fmt"
import "time"

func main() {
	var timeChannel chan time.Time = make(chan time.Time)

	emitTime := createTimeEmitter()

	go emitTime(timeChannel)
	go speak(timeChannel)

	var input string
	fmt.Scanln(&input)
}

func speak(timeChannel chan time.Time) {
	for {
		timeStruct := <- timeChannel
		fmt.Printf("The time is %s...\n", timeStruct.Format("15:04"))
	}
}

func createTimeEmitter() (func(chan time.Time)) {
	ticker := time.NewTicker(1 * time.Second)
	prev := time.Now().Add(-time.Minute);
	return func(timeChannel chan time.Time) {
		for {
			now := <- ticker.C
			if now.Minute() != prev.Minute() {
				timeChannel <- now
				prev = now
			}
		}
	}
}
