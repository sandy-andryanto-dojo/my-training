package main

import "fmt"

func f2() (int, int) {
	return 5, 6
}

func f(x int) {
	fmt.Println(x)
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {

	// xs := []float64{98, 93, 77, 82, 83}
	// fmt.Println(average(xs))

	// x := 5
	// f(x)

}
