package Lesson12

import (
	"fmt"
	"time"
)

type Square struct {
	Side int
}

type OurType string

func Methods() {
	//definitions()
	rules()
}

func rules() {
	// 1. Тип должен быть объявлен в текущем пакете
	// 5. Является объявленным типом

	now := time.Now()
	fmt.Printf("%T, %#v \n", now, now)

	//var OurString OurType = "hello"

}

// 1. Можем так сделать, т.к. тип в этом же пакете
func (t OurType) Hello() { 
	fmt.Println("Hello")
}

// Так сделать не можем, т.к. time находится в другом пакете
/* 
func (t time.Time) HelloTime() {
	fmt.Println("Hello time")
}*/

// 2. Так не можем, т.к. тип не должен быть указателем:
type pInt *int

/*
func (pInt) Test(){
	fmt.Println("Test")
}
*/

// 3. Интерфейсный тип
type Tester interface {
	Hello()
}

// так тоже делать нельзя
/*
func (t Tester) Test() {
	fmt.Println("Test")
}*/

// 4. Builtin тип
// так нельзя
/*
func (i int64) Test() {
	fmt.Println("Test")
}*/

func definitions() {
	square := Square {Side: 4}
	pSquare := &square

	square2 := Square {Side : 2}

	square.Perimeter() // Square {Side : 4}
	square2.Perimeter()

	pSquare.Scale(2) // &Square{Side : 4}
	pSquare.Perimeter() // (*pSquare).Perimeter() Square{Side : 8}

	square.Scale(2)  // (&square).Scale
	pSquare.Perimeter() // (*pSquare).Perimeter

	square.WrongScale(2)
	square.Perimeter()
}

func (s Square) Perimeter() {
	fmt.Printf("%T, %#v \n", s, s)
	fmt.Printf("Периметр фигуры: %d \n\n", s.Side * 4)
}

func (s *Square) Scale(multiplier int) {
	fmt.Printf("%T, %#v \n\n", s, s)
	s.Side *= multiplier
}

func (s Square) WrongScale(multiplier int) {
	fmt.Printf("%T, %#v \n", s, s)
	s.Side *= multiplier
	fmt.Printf("%T, %#v \n\n", s, s)
}