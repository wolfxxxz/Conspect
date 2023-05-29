package main

import (
	"fmt"
	"runtime"
	"time"
)

func writeChanal(ch chan<- int, b int) {
	ch <- b
	close(ch)
}

func read(ch, quit <-chan int) {
	for {
		select {
		case x := <-ch:
			fmt.Println("ch = ", x)
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("default")
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan int)

	go read(ch, quit)

	go writeChanal(quit, 10)
	go writeChanal(ch, 5)
	runtime.Gosched()
	time.Sleep(time.Second * 1)

}
