package main

import (
	"fmt"
	"strconv"
	"time"
)

func philosopher(id int, f_left, f_right chan struct{}, query <-chan struct{}, result chan<- string) {
	var eating bool
	var count int

	go func() { // go routines for recieving query and sending results back to main.
		for {
			<-query
			result <- p_print(id, eating, count)
		}
	}()

	for {
		<-f_left // Recieve forks
		<-f_right
		eating = true
		count++
		fmt.Println("Philosopher", id, "is eating")
		time.Sleep(1 * time.Second) // eating for a second
		f_left <- struct{}{}        // Put forks back
		f_right <- struct{}{}
		eating = false
	}
}

func p_print(id int, eating bool, count int) string {
	if eating {
		return "Philosopher " + strconv.Itoa(id) + " is eating and has eaten " + strconv.Itoa(count) + " times"
	} else {
		return "Philosopher " + strconv.Itoa(id) + " is thinking and has eaten " + strconv.Itoa(count) + " times"
	}

}
