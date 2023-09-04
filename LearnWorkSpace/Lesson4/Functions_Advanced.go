package Lesson4

import (
	"fmt"
)

func Functions_Advanced() {

	first, second := 5, 3

	// Работа с функциями, как со значениями
	var multiplier func(x, y int) int                         // объявление переменной с типом функции
	multiplier = func(x, y int) int { return x * y }          // присваивание функции данной переменной
	fmt.Printf("multiplier: %v\n", multiplier(first, second)) // непосредственный вызов функции

	divider := func(x, y int) int { return x / y }
	fmt.Printf("divider: %v\n", divider(first, second))

	// Передача функции в аргументы другой функции
	sumFunc := func(x, y int) int {
		return x + y
	}

	subtractFunc := func(x, y int) (result int) {
		result = x - y
		return
	}

	fmt.Printf("sumFunc: %v\n", calculate(first, second, sumFunc))
	fmt.Printf("subtractFunc: %v\n", calculate(first, second, subtractFunc))

	// Возврат функции из функции
	divideBy2 := createDivider(2)
	divideBy10 := createDivider(10)

	fmt.Println("Divide by 2:", divideBy2(100))
	fmt.Println("Divide by 10:", divideBy10(100))

	// Замыкания
	dollar := 30

	getDollarValue := func() int {
		return dollar
	}

	fmt.Println("Dollar value1: ", getDollarValue())

	dollar = 70
	fmt.Println("Dollar value2: ", getDollarValue())

	// Closure example
	digitForTest := 10
	closureFunc1 := TestingClosure(digitForTest)

	fmt.Print("ClosureFunc1 (10): ")
	closureFunc1()

	digitForTest = 20

	fmt.Print("ClosureFunc1 (20): ")
	closureFunc1()

}

func calculate(x, y int, action func(x, y int) int) int {
	return action(x, y)
}

func createDivider(divider int /* 2 */) func(x int) int {
	dividerFunc := func(x int) int {
		return x / divider /* 2 */
	}

	return dividerFunc
}

func TestingClosure(digit int) func() {

	ClosureFunc := func() {
		fmt.Println("Digit = ", digit)
	}

	return ClosureFunc
}
