package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan bool)
	c := boring("Joe", quit)
	rand := rand.Intn(30)
	fmt.Println("Random number:", rand)
	for i := rand; i >= 0; i-- {
		fmt.Println(<-c)
	}
	// quit 只是帮助退出 goroutine
	quit <- true
}

func boring(msg string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing
			case <-quit:
				return
			}
		}
	}()
	return c
}
