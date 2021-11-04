package main

import (
	"fmt"
	"sync"
)

var x int64
var wg sync.WaitGroup
var ch = make(chan int64,1)

func add() {
	for i := 0; i < 50000; i++ {
		x :=<-ch
		x++
		ch <- x
	}
	wg.Done()
}
func main() {
	ch <- 0
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(<-ch)
}
