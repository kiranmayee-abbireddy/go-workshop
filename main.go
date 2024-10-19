package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var age int = 5
	var name string = "Kiranmayee"
	var height float32 = 5.03
	fmt.Printf("My name is %s and my age is %d\n", name, age)
	fmt.Printf("My height is %.2f\n", height)

	place := "Hyderabad"
	runs := 100
	fmt.Println("I am from", place)

	if runs > 100 {
		fmt.Println("You have scored a century")
	} else {
		fmt.Println("Yoy have scored less than a century")
	}
}
