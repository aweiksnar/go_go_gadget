// write a function that takes an integer and halves it and returns true if it was even or false if it was odd
// for example half(1) should return (0, false) and half(2) should return (1, true)

package main

import "fmt"

func half(num int) (int, bool) {
	return (num / 2), (num%2 == 0)
}

func main() {
	fmt.Println(half(1)) // 0 false
	fmt.Println(half(2)) // 1 true
}
