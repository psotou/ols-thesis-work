package stats

import (
	"strconv"

	"gonum.org/v1/gonum/stat"
)

type Stats struct {
	Beta
	RSquared        float64
	CorrCoef        float64
	NumObservations float64
	PValue          float64
}

type Beta struct {
	Beta0 float64
	Beta1 float64
}

func StatsValues(data [][]string, scale float64) (Stats, []float64, []float64, []float64) {
	var stVal Stats
	X := make([]float64, len(data)-1)    // Vector X -> rating madurez BIM
	Y := make([]float64, len(data)-1)    // Vector Y -> desviación porcentual de costos
	Xind := make([]float64, len(data)-1) // Vector Xind -> indicador de madurez BIM propuesto
	var (
		weights []float64
		origin  bool = false
	)

	for i, line := range data[1:] {
		X[i], _ = strconv.ParseFloat(line[1], 64)
		Y[i], _ = strconv.ParseFloat(line[0], 64)
		Xind[i] = scale / X[i]
	}

	stVal.Beta.Beta0, stVal.Beta.Beta1 = stat.LinearRegression(Xind, Y, weights, origin)
	stVal.RSquared = stat.RSquared(Xind, Y, weights, stVal.Beta.Beta0, stVal.Beta.Beta1)
	stVal.CorrCoef = stat.Correlation(Xind, Y, weights)
	stVal.NumObservations = float64(len(Xind))
	stVal.PValue = TwoSidedPValue(stVal.CorrCoef, stVal.NumObservations)
	return stVal, X, Y, Xind
}
