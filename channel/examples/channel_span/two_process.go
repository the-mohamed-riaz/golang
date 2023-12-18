package main

import (
	"fmt"
	"time"

	"github.com/the-mohamed-riaz/golang/common"
)

func process2(n int, channel chan int) {
	time.Sleep(time.Second * 1)
	channel <- (n * 100)
}
func main() {
	defer common.Exec()()
	println("i am process 1")
	ints := []int{1, 2, 3, 4}
	l := len(ints)
	channel := make(chan int, l)
	// write in channel
	for _, n := range ints {
		go process2(n, channel) // with this total exec time is ~ 1 sec
		// process2(n, channel) // with this total exec time is ~ 4 sec
	}
	for i := 0; i < l; i++ {
		val := <-channel
		fmt.Printf("%v\n", val)
	}
}
