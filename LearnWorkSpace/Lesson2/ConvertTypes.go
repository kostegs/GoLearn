package Lesson2

import (
	"fmt";
	"unsafe"
)

func ConvertTypes(){

	fmt.Println(unsafe.Sizeof(int8(1))); // 1
	fmt.Println(unsafe.Sizeof(int16(1))); // 2
	fmt.Println(unsafe.Sizeof(int32(1))); // 4
	fmt.Println(unsafe.Sizeof(int64(1))); // 8
	fmt.Println(unsafe.Sizeof(string("Hello!"))); // 16
	fmt.Println(unsafe.Sizeof(string("Hello World! This is the program created on Go-programming language"))); // 30

	var num1 int64 = 15;
	num2 := 15;

	fmt.Println(num1 + int64(num2));


}