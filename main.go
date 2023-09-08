package main

import "fmt"

func main() {
	xi := []float64{3, 5, 9}
	fmt.Println("The average is ", Average(xi))
}

func Average(n []float64) float64 {
	var sum float64 = 0

	for _, v := range n {
		sum += v
	}
	a := len(n)
	var l float64 = float64(a)
	Ave := sum / l
	return Ave
}
