package main

import "fmt"

func main() {
	var n int //n is the user input to find the fibonacci(n)
	fmt.Println("Enter the number for which fibonacci series is to be printed :")
	fmt.Scan(&n)
	if n >= 0 {
		fmt.Println("Fibonacci of (", n, ") is:", Fibonacci(n))
	} else {
		fmt.Println("Please enter positive number :")
		fmt.Scan(&n)
		fmt.Println("Fibonacci of (", n, ") is:", Fibonacci(n))
	}

}

/*Fibonacci Function to find fib(n)/*
 * Function Fibonacci
 * Receives the input in the form of an integer
 * It computes the fibonaci(n) as per the logic in the code.
 *
 * @param (number int) User input in the form of an integer.
 * @return (perimter) Integer with the value = fibonacci(n)
 */
func Fibonacci(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		//Recursive call to fibonaci function
		return Fibonacci(number-1) + Fibonacci(number-2)
	}
}
