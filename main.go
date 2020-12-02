package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	// Abrimos el archivo data.csv
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	// Generamos un reader
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err == io.EOF {
		fmt.Println("Error: ", err)
		return
	}

	X := make([]float64, len(lines)-1)    // Vector X -> rating madurez BIM
	Y := make([]float64, len(lines)-1)    // Vector Y -> desviación porcentual de costos
	Xind := make([]float64, len(lines)-1) // Vector Xind -> indicador de madurez BIM propuesto
	var weights []float64
	var origin bool = false

	for i, line := range lines[1:] {
		X[i], _ = strconv.ParseFloat(line[1], 64)
		Y[i], _ = strconv.ParseFloat(line[0], 64)
		Xind[i] = 4 / X[i]
	}

	alpha, beta := stat.LinearRegression(Xind, Y, weights, origin) // Vector Beta, con alpha = beta_0
	r2 := stat.RSquared(Xind, Y, weights, alpha, beta)             // Calculated r squared
	corrCoef := stat.Correlation(Xind, Y, weights)                 // calculated correlation coefficient
	numObservations := float64(len(Xind))                          // N observations
	pvalue := twoSidedPValue(corrCoef, numObservations)

	fmt.Printf("Vector Madurez BIM:           %.3f\n", X)
	fmt.Printf("Vector Indicador Madurez BIM: %.3f\n", Xind)
	fmt.Printf("Vector Desviación de costos:  %.3f\n", Y)
	fmt.Println("\n         Coeficiente    p-value       R^2")
	fmt.Printf("beta_1:       %.4f     %.4f    %.4f\n", beta, pvalue, r2)
	fmt.Printf("\nYi = %.4f + %.4f * Xi \n", alpha, beta)
}

func twoSidedPValue(r float64, n float64) float64 {
	// compute the test stat
	ts := r * math.Sqrt((n-2)/(1-r*r))

	// make a Student's t with (n-2) d.f. Asume que se trata de una distribución Normal(0,1)
	t := distuv.StudentsT{Mu: 0, Sigma: 1, Nu: (n - 2), Src: nil}

	// compute the p-value
	pval := 2 * t.CDF(-math.Abs(ts))

	return pval
}
