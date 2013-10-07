/*
challenge: write a program that finds the smallest mumber in this list:
x := []int{
  48, 96, 86, 68,
  57, 82, 63, 70,
  37, 34, 83, 27,
  19, 97, 9, 17,
}
*/

package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(smallestNum(x)) // 9
}

func smallestNum(nums []int) int {
	smallNum := nums[0]
	for _, value := range nums {
		if value < smallNum {
			smallNum = value
		}
	}
	return smallNum
}
