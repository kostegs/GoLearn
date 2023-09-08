package Lesson11

import (
	"fmt"
	"time"
)

type OurString string
type OurInt int

type Person struct {
	Name string
	Age  int
}

func CustomTypes_Struct() {
	var customString OurString
	fmt.Printf("%T %#v \n", customString, customString)

	customString = "Hello, dudes"
	fmt.Printf("%T %#v \n", customString, customString)

	customInt := OurInt(5)
	fmt.Printf("%T %#v \n", customInt, customInt)
	fmt.Println(int(customInt))

	// create default
	var John Person
	fmt.Printf("%T, %#v \n", John, John)

	John = Person{}
	fmt.Printf("%T, %#v \n", John, John)

	// fields accessing
	John.Name = "John"
	John.Age = 30
	fmt.Println(John)

	// create with named field
	Brad := Person{
		Name: "Brad",
		Age:  25,
	}
	fmt.Println(Brad)

	// create without fields names
	Vladimir := Person{"Vladimir", 40}
	fmt.Println(Vladimir)

	// fields accessing through the pointer
	pVladimir := &Vladimir
	fmt.Println(*pVladimir)
	fmt.Println(pVladimir.Age)

	// create pointer to struct
	pIvan := &Person{"Ivan", 90}
	fmt.Println(*pIvan)

	// unnamed struct
	unnamedStruct := struct {
		Name, LastName, BirthDate string
	}{
		Name:      "NoName",
		LastName:  "NoLastName",
		BirthDate: fmt.Sprintf("%s", time.Now()),
		// BirthDate : time.Now().String(),
	}
	fmt.Println(unnamedStruct)
}
