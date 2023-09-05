package Lesson8

import (
	"fmt"
	"math/rand"
	"time"
)

const(
	min = 1
	max = 5
)

func SwitchCase(){

	rand.NewSource(time.Now().UnixNano())

	value := rand.Intn(max - min) + min;

	// if example
	fmt.Print("If example: ")
	if value == 1{
		fmt.Println("Number is one")
	} else if value == 2 || value == 3{
		fmt.Println("Number is two or three")
	} else if value == getFour(){
		fmt.Println("Number is four")
	} else {
		fmt.Println("Default case is shown")
	}

	// base switch
	fmt.Print("Base switch: ")
	switch value{
	case 1:
		fmt.Println("Number is one")
	case 2,3:
		fmt.Println("Number is two or three")
	case getFour():
		fmt.Println("Number is four")
	default:
		fmt.Println("Default case is shown")
	}

	// switch with local variable definition
	fmt.Print("Switch with local variable definition: ")
	switch num := rand.Intn(max - min) + min; num {
	case 1:
		fmt.Println("Number is one")
	case 2,3:
		fmt.Println("Number is two or three")
	case getFour():
		fmt.Println("Number is four")
		fallthrough // выполнить строчку из следующего условия
	case 10:
		fmt.Println("Strange things happens")
	default:
		fmt.Println("Default case is shown")	
	}

	// switch without condition
	fmt.Print("Switch without condition: ")
	switch {
	case value > 2:
		fmt.Printf("Value %d is greater than 2\n", value)
	case value < 2:
		fmt.Printf("Value %d is less than 2\n", value)		
	default:
		fmt.Println("Value is equals 2")
	}
}

func getFour() int {	
	return 4;
}
