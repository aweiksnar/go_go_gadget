// write a function with one variadic parameter that finds the greatest number in a list of numbers

package main

import "fmt"

func max(nums ...int) int {
  var maxNum int
	for _, i := range nums {
		if i > maxNum {
			maxNum = i
		}
	}
  return maxNum
}

func main() {
  fmt.Println(max(1, 2, 13, 3, 8, 4)) // 13
}
