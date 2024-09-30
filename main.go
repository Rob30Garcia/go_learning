package main

import (
	"fmt"

	"rsc.io/quote"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	fmt.Println(quote.Go())
	fmt.Println()

	// change elements of a slice
	fmt.Println("change elements of a slice")
	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])
	fmt.Println()

	// Append elements to a slice
	fmt.Println("Append elements to a slice")
	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
	fmt.Println()

	// Append one slice to another slice
	fmt.Println("Append one slice to another slice")
	myslice1 = []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))
	fmt.Println()

	// Change the length of a slice
	fmt.Println("Change the length of a slice")
	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 = arr1[1:5]                  // slice array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	fmt.Println("--change length by re-slicing the array--")
	myslice1 = arr1[1:3] // change length by re-slicing the array
	fmt.Printf("myslice1 =  %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	fmt.Println("--change length by appending items--")
	myslice1 = append(myslice1, 20, 21, 22, 23) // change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
	fmt.Println()

	// Memory Efficiency
	fmt.Println("Memory efficiency")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	fmt.Println("Create copy with only needed numbers")
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	fmt.Println(numbersCopy)
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

	fmt.Println()
	fmt.Println()
	// Single-Case switch example
	fmt.Println("Single-Case switch example")
	day := 7

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a weekday")
	}

	day = 7
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	fmt.Println()
	fmt.Println()

	// Access Struct Members
	fmt.Println("Access Struct Members")
	var person_one Person
	var person_two Person

	// person_one specification
	person_one.name = "Hege"
	person_one.age = 45
	person_one.job = "Teacher"
	person_one.salary = 6000

	// person_two specification
	person_two.name = "Cecilie"
	person_two.age = 24
	person_two.job = "Marketing"
	person_two.salary = 4500

	// Access and print person_one info
	fmt.Println("Name: ", person_one.name)
	fmt.Println("Age: ", person_one.age)
	fmt.Println("Job: ", person_one.job)
	fmt.Println("Salary: ", person_one.salary)
	fmt.Println()

	// Access and print person_two info
	fmt.Println("Name: ", person_two.name)
	fmt.Println("Age: ", person_two.age)
	fmt.Println("Job: ", person_two.job)
	fmt.Println("Salary: ", person_two.salary)
	fmt.Println()

	// Pass struct as Function Arguments
	fmt.Println("Pass struct as Function Arguments")
	print_person(person_one)
	print_person(person_two)

	fmt.Println()

	// go maps
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
	fmt.Println()

	// Create Maps Using make() function:
	var a_make = make(map[string]string) // The map is empty now
	a_make["brand"] = "Ford"
	a_make["model"] = "Mustang"
	a_make["year"] = "1964"

	b_make := make(map[string]int)
	b_make["Oslo"] = 1
	b_make["Bergen"] = 2
	b_make["Trondheim"] = 3
	b_make["Stavanger"] = 4

	fmt.Printf("a_make\t%v\n", a_make)
	fmt.Printf("b_make\t%v\n", b_make)
}

func print_person(person Person) {
	fmt.Println("Name: ", person.name)
	fmt.Println("Age: ", person.age)
	fmt.Println("Job: ", person.job)
	fmt.Println("Salary: ", person.salary)
	fmt.Println()
}
