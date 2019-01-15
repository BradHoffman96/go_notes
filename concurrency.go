package main

//A goroutine is a lightweight thread managed by the Go runtime
//Goroutines run in the same address space, so access to shared memory must be synchronized.

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

//select statements let a goroutine wait on multiple communication operations
//use the default case to try a send or receive without blocking
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("      .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	//go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}

	//Channels must be created before use with the make() function
	//By default, sends and receives block until the other side is ready
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	//Data flows in the direction of the arrow. In this case, from the channel to the variable
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	//channels can be buffered, in that there is a buffer length that is taken as an argument to make
	//Send will block when the buffer is full, receives will block when the buffer is empty
	//ch := make(chan int, 100)

	//Channels can be closed using close(), but this should only be done from the sending side, not the receiving side
	//Essentially it is done when the sender knows no more values are going to be sent through the channel

	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fibonacci(ch, quit)
}
