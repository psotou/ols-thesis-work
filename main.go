package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

var (
	flagFile  string
	flagScale float64
)

func main() {
	flag.StringVar(&flagFile, "file", "data.csv", "file is the name of the csv file")
	flag.Float64Var(&flagScale, "rmax", 4.0, "mrate is the maximum rating according to the sale of BIM measurement used")
	flag.Parse()

	// Abrimos el archivo data.csv
	file, err := os.Open(flagFile)
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
	var (
		weights []float64
		origin  bool = false
	)

	for i, line := range lines[1:] {
		X[i], _ = strconv.ParseFloat(line[1], 64)
		Y[i], _ = strconv.ParseFloat(line[0], 64)
		Xind[i] = flagScale / X[i]
	}

	alpha, beta := stat.LinearRegression(Xind, Y, weights, origin) // Vector Beta, con alpha = beta_0
	r2 := stat.RSquared(Xind, Y, weights, alpha, beta)             // Calculated r squared
	corrCoef := stat.Correlation(Xind, Y, weights)                 // calculated correlation coefficient
	numObservations := float64(len(Xind))                          // N observations
	pvalue := twoSidedPValue(corrCoef, numObservations)

	fmt.Println("\n=================================")
	fmt.Printf("%10s %10s %7v/Xi\n", "Xi", "Yi", flagScale)
	fmt.Println("---------------------------------")
	for i := range X {
		fmt.Printf("%10.2f %10.2f %10.2f\n", X[i], Y[i], Xind[i])
	}
	fmt.Println("---------------------------------")

	fmt.Println("\nXi  : Madurez BIM")
	fmt.Printf("%v/Xi: Indicador de inmadurez BIM\n", flagScale)
	fmt.Println("Yi  : Crecimiento de costos de construcción")

	fmt.Println("\n============================================")
	fmt.Println("         Coeficiente    p-value    R-squared")
	fmt.Println("--------------------------------------------")
	fmt.Printf("    B1:       %.4f     %.4f       %.4f\n", beta, pvalue, r2)
	fmt.Println("--------------------------------------------")
	fmt.Printf("\n\n")
	fmt.Println("============================================")
	fmt.Println("       Ecuación del modelo propuesto")
	fmt.Println("--------------------------------------------")
	fmt.Printf("     Yi = %.4f + %.4f * (%v / Xi) \n", alpha, beta, flagScale)
	fmt.Println("--------------------------------------------")

	// PLOT
	modelPlot(flagScale, alpha, beta)
}

func makeRange(min, max int) []float64 {
	a := make([]float64, (max-min+1)*5)
	a[0] = 1.0
	for i := 1; i < len(a); i++ {

		a[i] = 0.2 + a[i-1]
	}
	return a
}

func modelPlot(scale, alpha, beta float64) error {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Modelo desviación de costos versus  madurez BIM"
	p.X.Label.Text = "Madurez BIM"
	p.Y.Label.Text = "Desviación costos"
	p.Add(plotter.NewGrid())

	// we generate the point for our estimated function
	xvalues := makeRange(1, int(flagScale))
	pts := make(plotter.XYs, len(xvalues))
	for i := range pts {
		pts[i].X = scale / xvalues[i]
		pts[i].Y = alpha + beta*xvalues[i]
	}

	s, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}

	p.Add(s)
	p.Legend.Add(fmt.Sprintf("y = a + b * %v/x", flagScale), s)

	p.X.Min = 1
	p.X.Max = flagScale + 0.5
	p.Y.Min = 0
	p.Y.Max = 0.35

	// we save to a png file
	if err := p.Save(7.5*vg.Inch, 5.5*vg.Inch, "model_estimate_plot.png"); err != nil {
		panic(err)
	}
	return err
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
