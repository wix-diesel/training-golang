package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.After(2 * time.Second)
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}

	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("fopher ", gopherID, " はスリープを終えました")
		case <-timeout:
			fmt.Println("忍耐の限度に達した！")
			return
		}
	}
}

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
}
