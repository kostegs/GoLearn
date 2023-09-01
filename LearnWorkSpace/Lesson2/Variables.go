package main

import "fmt"

func main() {

	name := "vasya"
	hello := fmt.Sprintf("Hello %s", name);
	
	fmt.Println(hello);

	fmt.Printf("Type: %T Value: %v\n", hello, hello)

	ConvertTypes();
}
