package main

import "fmt"

/*
	- Ngay: 22.07.2021
	- By: Nguyen Dang Duc
*/

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// 01. Khai báo biến có kiểu dữ liệu là Person
	var (
		person1 = Person{
			FirstName: "Duc",
			LastName:  "Nguyen Dang",
			Age:       30,
		}
		person2 Person
		person3 = Person{FirstName: "Duy"}
	)

	// 02. In Màn hình
	fmt.Println("01. Person 1:", person1)
	fmt.Println("02. Person 2:", person2)
	fmt.Println("03. Person 3:", person3)
}
