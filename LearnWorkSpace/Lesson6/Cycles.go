package Lesson6

import "fmt"

func Cycles() {

	// Increment
	i := 0
	i = i + 1
	i += 1
	i++

	fmt.Println("increment i = ", i)

	// Decrement
	i = 10

	i = i - 1
	i -= 1
	i--

	fmt.Println("decrement i = ", i)

	// usual cycle
	fmt.Print("cycle, i = ")
	for i := 0; i < 10; i++ {
		fmt.Print(" ", i)
	}

	// for as while
	fmt.Print("\nfor as while, sum = ")
	sum := 1
	for sum <= 20 {
		sum += sum
		fmt.Print(" ", sum)
	}

	fmt.Print("\nfor as while 2, sum = ")
	for sum <= 100 {
		sum++
		fmt.Print(" ", sum)
	}

	// continue
	fmt.Print("\ncontinue, i = ")
	for i := 0; i <= 20; i++ { // 5 % 2 = 1
		if i%2 == 1 {
			continue
		}

		fmt.Print(" ", i)
	}

	// continue with label
	fmt.Print("\ncontinue with label, i | j:\n")
Label1:
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i, " | ", j)

			if i >= 10 {
				continue Label1
			}
		}
	}

	// break
	fmt.Print("\ncycle with break, i = ")
	for i := 0; i <= 20; i++{
		if i >= 10{
			break;
		}
		fmt.Print(" ", i);
	}

	// break with label
	fmt.Print("\nbreak with label, i | j:\n")
Label2:
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i, " | ", j)

			if i >= 10 {
				break Label2
			}
		}
	}
}
