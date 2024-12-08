package main

import "fmt"

func main() {
	fmt.Println("welcome to the calc")
	type square struct {
		length float64
		width  float64
	}
	type trapezium struct {
		a float64
		b float64
	}
	trapezium1 := trapezium{a: 12, b: 20}
	calculateAreaOfTrapezium(trapezium1.a, trapezium1.b)

	name := "john"
	fmt.Println("first letter:", name[0:1])
	for i := 0; i < 100; i++ {
		fmt.Println("i", i)
	}
}
func calculateAreaOfTrapezium(a float64, b float64) {
	area := 0.5 * a * b
	fmt.Println("area of trapezium is:", area)
}
