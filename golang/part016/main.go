package main

import "fmt"

// func main() {

// 	i:= factorial(3)
// 	fmt.Println(i)
// }

// func factorial(x uint) uint{
// 	if x == 0{
// 		return 1
// 	}
// 	return x * factorial(x-1)
// }

// func first(){
// 	fmt.Println("1st")
// }

// func second(){
// 	fmt.Println("2nd")
// }

// func main(){
// 	// defer second()
// 	// first()
// 	first()
// 	second()
// }

func main(){
	defer func(){
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}