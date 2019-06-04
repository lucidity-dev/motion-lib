package path

import (
    "fmt"
    "math"

)

type CubicSpline struct {
    //Spline coefficients
    //x(t) = Ax*t^3 + Bx*t^2 + Cx*t + Dx
    //y(t) = Ay*t^3 + By*t^2 + Cy*t + Dy
    Ax, Bx, Cx, Dx float64
    Ay, By, Cy, Dy float64
}

func GenerateCubicSpline(start Pose, end Pose) *CubicSpline {
    x0 := start.X
    x1 := end.X
    y0 := start.Y
    y1 := end.Y
    scale := 1.5*Dist(start, end)
    vx0 := scale*Cos(start)
    vx1 := scale*Cos(end)
    vy0 := scale*Sin(start)
    vy1 := scale*Sin(end)

    //spline coeffs
    ax := CubicA_Pos1*x1 + CubicA_Pos0*x0 + vx0 + vx1
    ay := CubicA_Pos1*y1 + CubicA_Pos0*y0 + vy0 + vy1
    bx := CubicB_Pos1*x1 + CubicB_Pos0*x0 + CubicB_Vel0*vx0 - vx1
    by := CubicB_Pos1*y1 + CubicB_Pos0*y0 + CubicB_Vel0*vy0 - vy1
    cx := vx0
    cy := vy0
    dx := x0
    dy := y0

    rv := CubicSpline{ax, bx, cx, dx, ay, by, cy, dy}

    return &rv
}

//t ranges from 0 to 1
func (this* CubicSpline) GetPose(t float64) Pose {
    x := this.Ax*math.Pow(t, 3) + this.Bx*math.Pow(t, 2) + this.Cx*t + this.Dx
    y := this.Ay*math.Pow(t, 3) + this.By*math.Pow(t, 2) + this.Cy*t + this.Dy
    dx := 3.0*this.Ax*math.Pow(t, 2) + 2.0*this.Bx*t + this.Cx;
    dy := 3.0*this.Ay*math.Pow(t, 2) + 2.0*this.By*t + this.Cy;
    theta := math.Atan2(dy, dx)
    return Pose{x,y,theta}
}

func (this* CubicSpline) GetCurvature(t float64) float64 {
    return 0.0;
}

func (this *CubicSpline) String() string {
    return fmt.Sprintf("x(t) = %.2ft^3 + %.2ft^2 + %.2ft + %.2f\ny(t) = %.2ft^3 + %.2ft^2 + %.2ft + %.2f", this.Ax, this.Bx, this.Cx, this.Dx, this.Ay, this.By, this.Cy, this.Dy)
}

