/*
->DEFER is used to schedule a function call to be run after a function completes.It is usually used to free resources as per the run time requirements.
->RECOVER stops the panic and returns the value that was passed to the call to panic.
->PANIC generally indicates a programmer error like division by zero,array index out of bounds,etc.
A call to panic immediately stops execution of the function.

In order for our system to recover back from the errorneous situation we can club the panic function with the defer as shown belows.

Example : Suppose in real time, we need to find the addition of two number even after system has thrown an exception(Panic).
We can achieve this using "Defer,Recover and Panic" all at once.
Below is an implementation for the same.
*/

package main

import "fmt"

func main() {
	value1, value2 := 5, 10
	add(value1, value2)
}
func add(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
		fmt.Println("-----Recovered from Panic----")
		fmt.Println("Performing addition ::", num1+num2)
	}()
	panic("FATALERROR::PANIC!")
}
