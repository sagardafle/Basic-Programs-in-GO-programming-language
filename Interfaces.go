package main

// implementing the fmt and math packages. math is used as we need the PI value to calculate the perimeter of circle.
import (
	"fmt"
	"math"
)

// An interface for shape type that contains the perimeter method
type shape interface {
	perimeter() float64
}

// A rectangle struct containing it's variable width and height.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func main() {
	var width, height, radius float64
	//Accepting the rectangle and circle parameters
	fmt.Println("Enter the height and width of Rectangle")
	fmt.Scan(&height)
	fmt.Scan(&width)
	if height <= 0 || width <= 0 {
		fmt.Println("Enter the positive value for height and width of Rectangle")
		fmt.Scan(&height)
		fmt.Scan(&width)
	}
	fmt.Println("Enter the radius of Circle")
	fmt.Scan(&radius)
	if radius <= 0 {
		fmt.Println("Enter the positive value for radius of circle")
		fmt.Scan(&radius)
	}

	rectangle := rect{width, height}
	circle := circle{radius}

	CalculatePerimeter(rectangle)
	CalculatePerimeter(circle)
}

/*
 * Function perimeter
 * Receives the 'Shape' as parameters in the form of circle and rectangle
 * It computes the parameters of the geometric figures.
 * @param (r rect,c circle) Shape Parameter containing the dimensions of their attributes.
 * @return (perimter) Integer type
 */
func (r rect) perimeter() float64 {
	Perirect := 2*r.height + 2*r.width
	fmt.Println("The perimeter of rectangle is :", Perirect)
	return Perirect
}
func (c circle) perimeter() float64 {
	Pericircle := 2 * math.Pi * c.radius
	//truncating the float64 to precision of 2 digits.
	truncated := float64(int(Pericircle*100)) / 100
	fmt.Println("The perimeter of circle is :", truncated)
	return truncated
}

/*
 * Function CalculatePerimeter
 * Generic method that uses parameters in the form of 'shape' interface.
 * @param (s shape) r:rect and c:circle as input
 * @return (float64)call to perimeter function
 */

// CalculatePerimeter :: using interface type as argument to function
func CalculatePerimeter(s shape) float64 {
	return s.perimeter()
}
