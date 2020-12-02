package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func main() {
	var x []float64 = []float64{1.0, 2.35, 1.82, 2.22, 4.0}
	// var x []float64 = []float64{4.0, 1.7021, 2.1978, 1.8018, 1.0}
	var y []float64 = []float64{0.33, 0.25, 0.162, 0.118, 0.001}
	var weights []float64
	var origin bool = false

	var xs = make([]float64, len(x))
	for i := range xs {
		xs[i] = 4 / x[i]
	}

	alpha, beta := stat.LinearRegression(xs, y, weights, origin)
	r2 := stat.RSquared(xs, y, weights, alpha, beta)

	fmt.Printf("Constante:  %.4f\n", alpha)
	fmt.Printf("Estimador:   %.4f\n", beta)
	fmt.Printf("R^2: %.6f\n", r2)
}
