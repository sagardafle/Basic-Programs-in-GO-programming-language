package main

import (
	"fmt"
	"time"
)

func main() {
	var duration int
	fmt.Println("Enter the value for sleep time")
	fmt.Scan(&duration)
	if duration > 0 {
		fmt.Println("--------------------")
		z := Sleep(duration)
		fmt.Println("Diffrence in intervals in seconds is:", z)
	} else {
		//Prompt for positive time interval
		fmt.Println("Enter the positive value for sleep time")
		fmt.Scan(&duration)
		fmt.Println("--------------------")
		interval := Sleep(duration)
		fmt.Println("Diffrence of time intervals in seconds is:", interval)
	}
}

/* Sleep function
 * Receives the 'sleepout' time interval as integer.
 * The Sleep function in turns then calls the After method.
 *
 * @param (sleepout int) The interval for which the After function is called.
 * @return (seconds) Integer The diffrence between time intervals before and after the sleep function is called.
 */
func Sleep(sleepout int) int {
	fmt.Println("Calling time.After()::Will return in", sleepout, "seconds")
	tinit := time.Now()
	<-time.After(time.Second * time.Duration(sleepout))
	z := time.Since(tinit)
	return int(z.Seconds())
}
