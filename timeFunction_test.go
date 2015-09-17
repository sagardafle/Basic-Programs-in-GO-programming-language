//To execute 'go test timeFunction_test.go timeFunction.go'
package main

import "testing"

type testpair struct {
	time  int
	sleep int
}

// The input parameter is the user input for time interval and the output parameter is the the output from Sleep function.
var tests = []testpair{
	{0, 0},
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 4},
	{5, 5},
	{6, 6},
	{7, 7},
	{8, 8},
	{9, 9},
	{10, 10},
}

func TestSleep(t *testing.T) {
	for _, pair := range tests {
		v := Sleep(pair.time)
		if v != pair.sleep {
			t.Error(
				"For", pair.time,
				"got", pair.sleep,
				"expected", v,
			)
		}
	}
}
