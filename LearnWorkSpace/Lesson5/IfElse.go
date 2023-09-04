package Lesson5

import "fmt"

func IfElse() {

	age := 15

	// Basic if
	if age < 18 {
		fmt.Println("You are too young (full)")
	}

	// Short syntax
	if isChild := isChildren(age); isChild {
		fmt.Println("You are very young (short)")
		fmt.Printf("You are child = %v\n", isChild)
	}

	// isChild - здесь недоступна!

	// If ... else
	age = 20

	if age < 18 {
		fmt.Println("You are too young (if ... else)")
	} else {
		fmt.Println("You are in school (if ... else)")
	}

	// &&
	if age >= 7 && age <= 18 {
		fmt.Println("You are in school (&&)")
	}

	// ||
	if age == 14 || age == 20 || age == 40 {
		fmt.Println("You have to get a new passport (||)")
	}

	// !
	if age != 14 {
		fmt.Println("You aren't 14 years old")
	}

}

func isChildren(age int) bool {
	return age < 18
}
