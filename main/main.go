package main

import "fmt"

func doFunction(fn func(int) int, number int) int {
	out := make(chan int)

	go func(output chan int) {
		output <- fn(number)
		close(out)
	}(out)

	return <-out
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	res := make([]int, n)
	closeChannel := make(chan interface{})

	go func(ch chan interface{}) <-chan interface{} {
		for i := 0; i < n; i++ {
			go func() {
				index := i
				item1 := <-in1
				item2 := <-in2
				res[index] = doFunction(fn, item1) + doFunction(fn, item2)
			}()
		}
		close(closeChannel)
		return closeChannel
	}(closeChannel)

	<-closeChannel

	for i := 0; i < n; i++ {
		out <- res[i]
	}
}

func main() {
	do(20)
}

func do(n int) {
	fn := func(i int) int {
		return i * 10
	}

	chan1 := make(chan int, n)
	chan2 := make(chan int, n)
	out := make(chan int, n)

	for i := 0; i < n; i++ {
		chan1 <- i
		chan2 <- i + 1
	}

	go merge2Channels(fn, chan1, chan2, out, n)

	for i := 0; i < n; i++ {
		fmt.Println(<-out)
	}

	close(chan1)
	close(chan2)
	close(out)
}
