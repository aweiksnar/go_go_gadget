// challenge: write a program that can swap two integers (x := 1; y := 2; swap(&x, &y)) should give you x=2 and y=1

package main

import "fmt"

func swap(xPtr *int, yPtr *int) {
	newX := *yPtr
	newY := *xPtr
	*xPtr = newX
	*yPtr = newY
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("x=", x, "y=", y)
}
