//To execute 'go test Interfaces_test.go Interfaces.go'
//This program  will ""FAIL"" as the last pair of values is invalid.
package main

import "testing"

type testpair struct {
	input shape
	Peri  float64
}

// Input ::rect{width,height} and circle{radius} parameters
//Output :: Perimeter of Circle and Rectangle
var tests = []testpair{
	{rect{1, 2}, 6},
	{rect{1, 1}, 4},
	{rect{7, 3}, 20},
	{rect{4, 5}, 18},
	{rect{3, 6}, 18},
	{circle{4}, 25.13},
	{circle{3}, 18.84},
	{circle{10}, 62.83},
	{circle{12}, 62.83}, // This will FAIL as the output here is incorrect.
}

//TestCalculatePerimeter function to test the CalculatePerimeter function.
func TestCalculatePerimeter(tr *testing.T) {
	for _, pair := range tests {
		v := CalculatePerimeter(pair.input)
		if v != pair.Peri {
			tr.Error(
				"For", pair.input,
				"got", pair.Peri,
				"expected", v,
			)
		}
	}
}
