package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEventGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	// fmt.Println(add(1, 2, 3))

	// xs := []int{1, 2, 3, 4}
	// fmt.Println(add(xs...))

	// add := func(x, y int) int {
	// 	return x + y
	// }

	// fmt.Println(add(1, 1))

	// x := 0

	// increment := func() int {
	// 	x++
	// 	return x
	// }

	// fmt.Println(increment())
	// fmt.Println(increment())

	nextEvent := makeEventGenerator()
	max := 100

	for i := 1; i < max; i++ {
		fmt.Println(nextEvent())
	}

}
