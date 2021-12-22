package main

import "fmt"

// const x string = "Hello World"

func main() {

	// fmt.Println(x)

	fmt.Print("Enter a number : ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)

}
