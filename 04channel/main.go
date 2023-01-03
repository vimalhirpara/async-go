package main

import "fmt"

func sum(a int, b int, c chan int) {
	c <- a + b // send sum to c
}

func main() {
	c := make(chan int)
	go sum(20, 20, c)
	x := <-c // receive from c

	fmt.Println(x)
}
