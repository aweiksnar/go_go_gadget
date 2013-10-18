// countdown timer using channels

package main

import (
	"fmt"
	"time"
)

// decrements through time and sends to channel
func timer(c chan string, min int) {
	t := time.Minute * time.Duration(min)
	for {
		c <- time.Duration.String(t)
		t = t - time.Second*1
		if t < time.Second*0 {
			break
		}
	}
}

// prints channel stream
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

// opens timing channel
func startTimer(mins int) {
	var c chan string = make(chan string)

	go timer(c, mins)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

//counts down from this many minutes until it hits zero
func main() {
	startTimer(2)
}
