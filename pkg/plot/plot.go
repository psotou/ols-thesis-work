package plot

import (
	"fmt"
	"ols-mem/pkg/utils"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func ModelPlot(scale, alpha, beta float64) error {
	p := plot.New()
	p.Title.Text = "Crecimiento costos de construcción versus madurez BIM"
	p.X.Label.Text = "Nivel de madurez BIM"
	p.Y.Label.Text = "Crecimiento costos de construcción"
	p.Add(plotter.NewGrid())

	// we generate the point for our estimated function
	xvalues := utils.MakeRange(1, int(scale))
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
	p.Legend.Add(fmt.Sprintf("y = a + b * %v/x", scale), s)

	p.X.Min = 1
	p.X.Max = scale + 0.5
	p.Y.Min = 0
	p.Y.Max = 0.35

	// we save to a png file
	if err := p.Save(7.5*vg.Inch, 5.5*vg.Inch, "model_estimate_plot.png"); err != nil {
		panic(err)
	}
	return err
}
