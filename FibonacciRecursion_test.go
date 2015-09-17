//To execute 'go test FibonacciRecursion_test.go FibonacciRecursion.go'
package main

import "testing"

type testpair struct {
	values int
	fib    int
}

// The input consists of user input for which Fibonacci(n) is to be computed whereas the output is the value of Fibonacci(n).
var tests = []testpair{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 558},
}

func TestFibonacci(t *testing.T) {
	for _, pair := range tests {
		v := Fibonacci(pair.values)
		if v != pair.fib {
			t.Error(
				"For", pair.values,
				"got", pair.fib,
				"expected", v,
			)
		}
	}
}
