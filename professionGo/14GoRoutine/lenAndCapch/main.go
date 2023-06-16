package main

import (
	"fmt"
	"time"
)

func calc(num int, ch chan int) {
	fmt.Println("len empty chan ", len(ch))
	f := num * 2
	ch <- f
	fmt.Println("len ch = ", len(ch))
}
func main() {
	ch := make(chan int, 2)
	fmt.Println("len(ch):", len(ch), " cap(ch):", cap(ch))
	for i := 1; i < 4; i++ {
		go calc(i, ch)
	}
	go func() {

		for v := range ch {
			fmt.Println(v)
		}

	}()

	time.Sleep(time.Second * 3)
}
