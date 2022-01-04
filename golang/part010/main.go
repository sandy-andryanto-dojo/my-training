package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {

	// var i uint
	// i = 4

	// bilangan := factorial(i)
	// fmt.Println(bilangan)

	// defer second()
	// first()

	// panic("PANIC")
	// str := recover()
	// fmt.Println(str)

	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()

	// panic("PANIC")

}
