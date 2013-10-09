// the fibonacci sequence is defined as fib(0) = 0 , fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
// write a recursive function that can find fib(n)

package main

import "fmt"

func fib(seqNum int) int {
	if seqNum == 0 {
		return 0
	} else if seqNum == 1 {
		return 1
	} else {
		return fib(seqNum-1) + fib(seqNum-2)
	}
}

func main() {
	fmt.Println(fib(0))  //  0
	fmt.Println(fib(1))  //  1
	fmt.Println(fib(9))  //  34
	fmt.Println(fib(14)) // 377
}
