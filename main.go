package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var age int = 21
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
	//i is only available in for loop

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 15 {
		fmt.Println(i)
		i++
	}

	j := 20
	for {
		if j > 25 {
			break
		}
		fmt.Println("Infinite loop")
		j++
	}

	nums := [5]int{100, 200, 300, 400, 500}
	for index, value := range nums {
		fmt.Println(index, value)
	}

	var x int64 = 10
	var y int64 = 20
	fmt.Println("Sum of the numbers is", add(x, y))
	fmt.Println("The value of x is", x)

	//Pointers
	var xp *int64
	xp = &x
	fmt.Println("The value of xp is", xp)
	xp = &y
	fmt.Println("The value of xp is", xp)

	//pass by reference
	fmt.Println("The current value of x is", x)
	fmt.Println("The value of x is", addByReference(&x, y))
	fmt.Println("The value of x after add by reference is", x)
}

func add(x int64, y int64) int64 {
	x = x + 10
	return x + y
}

func addByReference(x *int64, y int64) int64 {
	*x = *x + 10
	return *x + y
}
