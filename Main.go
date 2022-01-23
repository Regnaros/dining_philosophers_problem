package main

import (
	"fmt"
	"time"
)

func main() {

	f1 := make(chan struct{})
	f2 := make(chan struct{})
	f3 := make(chan struct{})
	f4 := make(chan struct{})
	f5 := make(chan struct{})

	p_query := make(chan struct{}, 5) // Send out a query to 5 different philosophers.
	p_result := make(chan string, 5)  // Let the philosopher send the results back to the main go routine.

	f_query := make(chan struct{}, 5) // Send out a query to 5 different forks.
	f_result := make(chan string, 5)  // Let the fork send the results back to the main go routine.

	go philosopher(1, f1, f2, p_query, p_result)
	go philosopher(2, f2, f3, p_query, p_result)
	go philosopher(3, f4, f3, p_query, p_result) // Picks up the right fork first. While everyone picks up the left fork first.
	go philosopher(4, f4, f5, p_query, p_result)
	go philosopher(5, f5, f1, p_query, p_result)

	go fork(1, f1, f_query, f_result)
	go fork(2, f2, f_query, f_result)
	go fork(3, f3, f_query, f_result)
	go fork(4, f4, f_query, f_result)
	go fork(5, f5, f_query, f_result)

	for {
		time.Sleep(time.Second * 5) // Query every 5 second
		fmt.Println()
		for i := 1; i <= 5; i++ {
			p_query <- struct{}{}
		}
		for i := 1; i <= 5; i++ {
			result := <-p_result
			fmt.Println(result)
		}
		fmt.Println()
		for i := 1; i <= 5; i++ {
			f_query <- struct{}{}
		}
		for i := 1; i <= 5; i++ {
			result := <-f_result
			fmt.Println(result)
		}
		fmt.Println()
	}
}
