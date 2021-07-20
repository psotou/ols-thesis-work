package utils

func MakeRange(min, max int) []float64 {
	a := make([]float64, (max-min+1)*5)
	a[0] = 1.0
	for i := 1; i < len(a); i++ {

		a[i] = 0.2 + a[i-1]
	}
	return a
}
