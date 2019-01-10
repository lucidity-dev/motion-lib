package path

import (
    "fmt"
    "math"

    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/plotutil"
    "gonum.org/v1/plot/vg"
)

type CubicSpline struct{
    //Spline coefficients
    //x(t) = Ax*t^3 + Bx*t^2 + Cx*t + Dx
    //y(t) = Ay*t^3 + By*t^2 + Cy*t + Dy
    Ax,Bx,Cx,Dx,Ay,By,Cy,Dy float64
}

func GenerateCubicSpline(start Point, end Point) CubicSpline {
    x0 := start.X
    x1 := end.X
    y0 := start.Y
    y1 := end.Y
    scale := 1.25*Dist(start, end)
    vx0 := scale*Cos(start)
    vx1 := scale*Cos(end)
    vy0 := scale*Sin(start)
    vy1 := scale*Sin(end)

    //spline coeffs
    ax := -2.0*x1 + 2.0*x0 + vx0 + vx1
    ay := -2.0*y1 + 2.0*y0 + vy0 + vy1
    bx := 3.0*x1 - 3.0*x0 - 2.0*vx0 - vx1
    by := 3.0*y1 - 3.0*y0 - 2.0*vy0 - vy1
    cx := vx0
    cy := vy0
    dx := x0
    dy := y0

    rv := CubicSpline{ax, bx, cx, dx, ay, by, cy, dy}

    return rv
}

//time ranges from 0 to 1
func (this* CubicSpline) GetPoint(time float64) Point {
    x := this.Ax*math.Pow(time, 3) + this.Bx*math.Pow(time, 2) + this.Cx*time + this.Dx
    y := this.Ay*math.Pow(time, 3) + this.By*math.Pow(time, 2) + this.Cy*time + this.Dy
    dx := 3.0*this.Ax*math.Pow(time, 2) + 2.0*this.Bx*time + this.Cx;
    dy := 3.0*this.Ay*math.Pow(time, 2) + 2.0*this.By*time + this.Cy;
    theta := math.Atan2(dy, dx)
    return Point{x,y,theta}
}

func (this *CubicSpline) String() string {
    return fmt.Sprintf("x(t) = %.2ft^3 + %.2ft^2 + %.2ft + %.2f\ny(t) = %.2ft^3 + %.2ft^2 + %.2ft + %.2f", this.Ax, this.Bx, this.Cx, this.Dx, this.Ay, this.By, this.Cy, this.Dy)
}

func (this *CubicSpline) Plot() {
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = "Cubic Spline"
    p.X.Label.Text = "X (position)"
    p.Y.Label.Text = "Y (position)"

    iters := 1000
    step := 1.0/(float64)(iters)
    points := make(plotter.XYs, iters)

    accum := 0.0
    for i := 0; i < iters; i += 1 {
        p := this.GetPoint(accum)
        points[i].X = p.X
        points[i].Y = p.Y
        accum += step
    }

    err = plotutil.AddLinePoints(p, "Cubic Spline", points)
    if err != nil {
        panic(err)
    }

    err = p.Save(12*vg.Inch, 12*vg.Inch, "CubicSpline.jpg")
    if err != nil {
        panic(err)
    }
}
