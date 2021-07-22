package main

import "fmt"

/*
	- Ngay: 22.07.2021
	- By: Nguyen Dang Duc
*/

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

func main() {
	// 01. Khai báo biến Struct
	var (
		myCar = Car{Name: "Honda",
			Model:      "Airblade",
			Color:      "Black",
			WeightInKg: 100}
	)

	//02 . In màn hình
	fmt.Println("Information : ==> ", myCar)
	fmt.Println("Name MyCar: ", myCar.Name)
	fmt.Println("Model MyCar: ", myCar.Model)
	fmt.Println("Color MyCar: ", myCar.Color)
	fmt.Println("Weight MyCar: ", myCar.WeightInKg)
}
