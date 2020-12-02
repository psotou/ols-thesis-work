package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
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

	X := make([]float64, len(lines)-1)    // Vectior X -> rating madurez BIM
	Y := make([]float64, len(lines)-1)    // Vectir Y -> desviación porcentual de costos
	Xind := make([]float64, len(lines)-1) // Vectir Xind -> indicador de madurez BIM propuesto
	var weights []float64
	var origin bool = false

	for i, line := range lines[1:] {
		X[i], _ = strconv.ParseFloat(line[1], 64)
		Y[i], _ = strconv.ParseFloat(line[0], 64)
		Xind[i] = 4 / X[i]
	}

	alpha, beta := stat.LinearRegression(Xind, Y, weights, origin)
	r2 := stat.RSquared(Xind, Y, weights, alpha, beta)

	fmt.Println("DATOS GENERALES")
	fmt.Printf("Vector Madurez BIM:           %.3f\n", X)
	fmt.Printf("Vector Indicador Madurez BIM: %.3f\n", Xind)
	fmt.Printf("Vector Desviación de costos:  %.3f\n", Y)
	fmt.Println("ESTADÍSTICOS")
	fmt.Printf("Constante:  %.4f\n", alpha)
	fmt.Printf("Estimador:   %.4f\n", beta)
	fmt.Printf("R^2:         %.4f\n", r2)
}
