package Lesson3

import (
	"fmt"
)

func Functions() {
	Greet();
	PersonGreet("Vasya");
	FioGreet("John", "Smith");

	//
	first, second := 1,2
	Sum := Sum(first, second);
	fmt.Printf("Sum(): Sum of %v and %v = %v\n", first, second, Sum);

	//
	Summa, Multiply := SumAndMultiply(first, second);
	fmt.Printf("SumAndMultiply(): Sum = %v, Multiply = %v\n", Summa, Multiply);

	//
	_, Multiply64 := NamedSumAndMultiply(first, second);
	fmt.Printf("NamedSumAndMultiply(): Multiply = %v\n", Multiply64);

}

func Greet(){
	fmt.Println("Greet(): Hello guys!")
}

func PersonGreet(name string){
	fmt.Printf("PersonGreet(): Zdarova, %s\n", name);
}

func FioGreet(name, surname string){
	fmt.Printf("FioGreet(): Hi, %s%s!\n", name, surname);
}

func Sum(first, second int) int {
	sum := first + second;
	return sum;
}

func SumAndMultiply(first, second int)(int, int){
	return first + second, first * second;	
}

func NamedSumAndMultiply(first, second int) (sum int64, multiply int64){
	sum = int64(first + second);
	multiply = int64(first * second);
	return; // или return sum, multiply
}

