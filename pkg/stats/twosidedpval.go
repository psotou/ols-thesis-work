package stats

import (
	"math"

	"gonum.org/v1/gonum/stat/distuv"
)

func TwoSidedPValue(r float64, n float64) float64 {
	// compute the test stat
	ts := r * math.Sqrt((n-2)/(1-r*r))

	// make a Student's t with (n-2) d.f. Asume que se trata de una distribución Normal(0,1)
	t := distuv.StudentsT{Mu: 0, Sigma: 1, Nu: (n - 2), Src: nil}

	// compute the p-value
	pval := 2 * t.CDF(-math.Abs(ts))

	return pval
}
