package main

import "fmt"

// func main() {

// 	xs := []float64{98, 93, 100, 77, 82, 83}
// 	fmt.Println(average(xs))

// }

// func average(xs []float64) float64{
// 	total := 0.0
// 	for _, v := range xs{
// 		total += v
// 	}
// 	return total / float64(len(xs))
// }


// var x int  = 5

// func f(){
// 	fmt.Println(x)
// }

// func main(){
// 	f()
// }


// func f() (int, int){
// 	return 5,6
// }

// func main(){
// 	x, y := f()
// }

// func add (args ... int) int {
// 	total := 0
// 	for _, v:= range args{
// 		total += v
// 	}
// 	return total
// }

// func main(){
// 	fmt.Println(add(1,2,3,4,5,6))
// }

// func main(){
// 	add := func (x, y int) int{
// 		return x + y
// 	}
// 	fmt.Println(add(12, 30))
// }

// func main(){
// 	x := 0
// 	increment := func() int {
// 		x++
// 		return x
// 	}
// 	fmt.Println(increment())
// 	fmt.Println(increment())
// 	fmt.Println(increment())
// }


func makeEventGenerator() func() uint{
	i := uint(0)
	return func() (ret uint){
		ret = i
		i += 2
		return
	}
}

func main(){
	nextEvent := makeEventGenerator()

	i:= 1
	max := 1000
	for i <= max {
		fmt.Println(nextEvent())
		i++
	}
	
}



























