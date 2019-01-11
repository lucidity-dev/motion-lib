package path

import (
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/plotutil"
    "gonum.org/v1/plot/vg"
)

type Spline interface {
    GetPose(t float64) Pose
    GetCurvature(t float64) float64
    String() string
}

func PlotSpline(s Spline, name string, path string) {
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = name
    p.X.Label.Text = "X (position)"
    p.Y.Label.Text = "Y (position)"

    iters := 1000
    step := 1.0/(float64)(iters)
    points := make(plotter.XYs, iters)

    accum := 0.0
    for i := 0; i < iters; i += 1 {
        p := s.GetPose(accum)
        points[i].X = p.X
        points[i].Y = p.Y
        accum += step
    }

    err = plotutil.AddLinePoints(p, name, points)
    if err != nil {
        panic(err)
    }

    err = p.Save(12*vg.Inch, 12*vg.Inch, path)
    if err != nil {
        panic(err)
    }
}

func PlotSplines(splines []Spline, name string, path string) {
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = name
    p.X.Label.Text = "X (position)"
    p.Y.Label.Text = "X (position)"

    for _, s := range splines {
        iters := 1000
        step := 1.0/(float64)(iters)
        points := make(plotter.XYs, iters)

        accum := 0.0
        for i := 0; i < iters; i += 1 {
            p := s.GetPose(accum)
            points[i].X = p.X
            points[i].Y = p.Y
            accum += step
        }

        err = plotutil.AddLinePoints(p, name, points)
        if err != nil {
            panic(err)
        }
    }

    err = p.Save(12*vg.Inch, 12*vg.Inch, path)
    if err != nil {
        panic(err)
    }

}
