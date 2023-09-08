package Lesson9

import "fmt"

func Pointers() {

	// default value
	var intPointer *int
	fmt.Printf("%T %#v \n", intPointer, intPointer)

	// nil pointer panic
	// fmt.Printf("%T %#v %#v \n", intPointer, intPointer, *intPointer);

	// Получение not-nil указателей
	var a int64 = 7
	fmt.Printf("%T %#v \n", a, a)

	// get variable pointer
	var pointerA *int64 = &a
	fmt.Printf("%T %#v %#v \n", pointerA, pointerA, *pointerA)

	// get pointer via new keyword
	var newPointer = new(float32)
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
	*newPointer = 3
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)

	// Pointers usage

	// side effects
	num := 3
	var numPointer *int = &num
	fmt.Printf("Значение num до вызова square = %#v \n", numPointer)
	square(num)
	fmt.Println(num)

	squarePointer(&num)
	fmt.Println(num)

	// empty value flag
	var wallet1 *int
	fmt.Println("has wallet 1 ", hasWallet(wallet1))

	wallet2 := 0
	fmt.Println("has wallet 2 ", hasWallet(&wallet2))

	wallet3 := 100
	fmt.Println("has wallet 3 ", hasWallet(&wallet3))
}

func square(num int) {
	var numPointer *int = &num
	fmt.Printf("Значение num в момент вызова square = %#v \n", numPointer)
	num *= num
}

func squarePointer(num *int) {
	fmt.Printf("Значение num в момент вызова squarePointer = %#v \n", num)
	*num *= *num
}

func hasWallet(money *int) bool {
	return money != nil
}
