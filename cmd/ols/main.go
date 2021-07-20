package main

import (
	"flag"
	"fmt"
	"ols-mem/pkg/plot"
	"ols-mem/pkg/stats"
	"ols-mem/pkg/utils"
)

var (
	flagFile  string
	flagScale float64
)

func main() {
	flag.StringVar(&flagFile, "file", "data.csv", "file is the name of the csv file")
	flag.Float64Var(&flagScale, "rmax", 4.0, "rmax is the maximum rating according to the sale of BIM measurement used")
	flag.Parse()

	lines := utils.ReadFile(flagFile)
	st, X, Y, Xind := stats.StatsValues(lines, flagScale)

	fmt.Println("\n============================================")
	fmt.Printf("%13s %10s %7v/Xi\n", "Xi", "Yi", flagScale)
	fmt.Println("--------------------------------------------")
	for i := range X {
		fmt.Printf("%13.2f %10.2f %10.2f\n", X[i], Y[i], Xind[i])
	}
	fmt.Println("--------------------------------------------")

	fmt.Println("\n============================================")
	fmt.Printf("%35s\n", "Vector coeficientes estimados")
	fmt.Println("--------------------------------------------")
	fmt.Printf("%18s %7.4f\n", "B0 =", st.Beta.Beta0)
	fmt.Printf("%18s %7.4f\n", "B1 =", st.Beta.Beta1)
	fmt.Println("--------------------------------------------")

	fmt.Println("\n============================================")
	fmt.Println("       Relación funcional estimada")
	fmt.Println("--------------------------------------------")
	fmt.Printf("     Yi = %.4f + %.4f * (%v / Xi) \n", st.Beta.Beta0, st.Beta.Beta1, flagScale)
	fmt.Printf("--------------------------------------------\n\n")

	fmt.Println("============================================")
	fmt.Println("         Coeficiente    p-value    R-squared")
	fmt.Println("--------------------------------------------")
	fmt.Printf("    B1:       %.4f     %.4f       %.4f\n", st.Beta.Beta1, st.PValue, st.RSquared)
	fmt.Println("--------------------------------------------")
	fmt.Printf("\n\n")

	// // PLOT
	plot.ModelPlot(flagScale, st.Beta.Beta0, st.Beta.Beta1)
}
