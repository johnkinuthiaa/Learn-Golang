package main

import (
	"fmt"
)

func main() {
	//words := flag.String("ls", "list", "getting all lists")
	//
	//flag.Parse()
	//fmt.Println(*words)
	x, _, _ := getNotification()
	fmt.Println(x)
	_, y, _ := getNotification()
	fmt.Println(y)

	car := Car{
		Make:   "20122",
		Model:  "Dodge challenger",
		Height: 120.9,
		width:  200.40}

	fmt.Println(car.getCarInfo())
}
func getNotification() (int, int, int) {
	return 10, 20, 30
}

type Car struct {
	Make   string
	Model  string
	Height float64
	width  float64
}

func (carInfo Car) getCarInfo() string {
	return fmt.Sprintf("Car name is %s and model %s", carInfo.Make, carInfo.Model)
}
