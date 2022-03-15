package main

import "fmt"

// func zero(x int){
// 	x = 0
// }

// func main() {
// 	x:= 5
// 	zero(x)
// 	fmt.Println(x)
// }


// func zero(xPtr *int){
// 	*xPtr = 0
// }

// func main(){
// 	x :=  5
// 	zero(&x)
// 	fmt.Println(x)
// }

func one(xPtr *int){
	*xPtr = 1
}

func main(){
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}