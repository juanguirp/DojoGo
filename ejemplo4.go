package main

import "fmt"

func main() {

	x := []float64{
  	48,96,86,68,
  	57,82,63,70,
  	37,34,83,27,
  	19,97, 9,17,
	}

	var promedio float64
	promedio = average(x)
	// Otra forma: promedio := average(x)
	fmt.Println("Promedio: ", promedio)

}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64 (len(xs))
}
