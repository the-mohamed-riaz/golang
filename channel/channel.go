package main

import "fmt"

type Que struct {
	items chan int
}

func (q *Que) Enque(item int) {
	q.items <- item
}

func (q *Que) Deque() int {
	return <-q.items
}

func main() {
	fmt.Println("start")
	person := []int{1, 2, 3, 4, 5}
	q := Que{
		items: make(chan int, len(person)),
	}

	for _, n := range person {
		q.Enque(n)
		fmt.Print("in Que: ", n, " ")
		fmt.Println("Deque: ", q.Deque())
	}
	// fmt.Println("Deque: ", q.Deque())
}

// NOTE
// ERROR: fatal error: all goroutines are asleep - deadlock!
// You will get this error, when you are trying to over-read from the channel
// over-read? expecting an input from the channel, when it has nothing.
// uncomment the extra deque print statement to observe the error.
