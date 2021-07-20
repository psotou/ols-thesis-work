package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ReadFile(dataFile string) [][]string {
	// Abrimos el archivo data.csv
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	defer file.Close()

	// Generamos un reader
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err == io.EOF {
		fmt.Println("Error: ", err)
		panic(err)
	}
	return lines
}
