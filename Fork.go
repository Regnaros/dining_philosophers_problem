package main

import "strconv"

func fork(id int, f_ch chan struct{}, query <-chan struct{}, result chan<- string) {
	var using bool
	var count int

	go func() { // go routines for recieving query and sending results back to main routine.
		for {
			<-query
			result <- f_print(id, using, count)
		}
	}()

	for {
		f_ch <- struct{}{} // Sends signal: ready to be picked up
		using = true
		count++
		<-f_ch // Recieve signal: forks are returned
		using = false
	}

}

func f_print(id int, using bool, count int) string {
	if using {
		return "Fork " + strconv.Itoa(id) + " is in use and has been used " + strconv.Itoa(count) + " times"
	} else {
		return "Fork " + strconv.Itoa(id) + " is not in use and has been used " + strconv.Itoa(count) + " times"
	}
}
