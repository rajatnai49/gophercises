package main

import (
	"fmt"
	"time"
)


func main() {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	resultChan := make(chan string)
	go func() {
		// resultChan <- "Operation Completed"
	}()

	select {
	case result := <-resultChan:
		fmt.Println(result)
	case <-timer.C:
		fmt.Println("TimedOut")
	}
}
