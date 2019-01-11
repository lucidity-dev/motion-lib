package path

import (
    "fmt"
    "math"
)

type QuinticSpline struct {
    //x(t) = Ax*t^5 + Bx*t^4 + Cx*t^3 + Dx*t^2 + Ex*t + F
    //y(t) = Ay*t^5 + By*t^4 + Cy*t^3 + Dy*t^2 + Ey*t + F
    Ax, Bx, Cx, Dx, Ex, Fx float64
    Ay, By, Cy, Dy, Ey, Fy float64
}

func GenerateQuinticSpline(start Pose, end Pose) *QuinticSpline {
    x0 := start.X
    x1 := end.X
    y0 := start.Y
    y1 := end.Y
    scale := 1.2*Dist(start, end)
    vx0 := scale*Cos(start)
    vx1 := scale*Cos(end)
    vy0 := scale*Sin(start)
    vy1 := scale*Sin(end)
    ax0 := 0.0
    ay0 := 0.0
    ax1 := 0.0
    ay1 := 0.0

    ax := -6.0*x0 + -3.0*vx0 - 0.5*ax0 + 0.5*ax1 - 3.0*vx1 + 6.0*x1
    ay := -6.0*y0 + -3.0*vy0 - 0.5*ay0 + 0.5*ay1 - 3.0*vy1 + 6.0*y1
    bx := 15.0*x0 + 8.0*vx0 + 1.5*ax0 - 1.0*ax1 + 7.0*vx1 - 15.0*x1
    by := 15.0*y0 + 8.0*vy0 + 1.5*ay0 - 1.0*ay1 + 7.0*vy1 - 15.0*y1
    cx := -10.0*x0 - 6.0*vx0 - 1.5*ax0 + 0.5*ax1 - 4.0*vx1 + 10.0*x1
    cy := -10.0*y0 - 6.0*vy0 - 1.5*ay0 + 0.5*ay1 - 4.0*vy1 + 10.0*y1
    dx := 0.5*ax0
    dy := 0.5*ay0
    ex := vx0
    ey := vy0
    fx := x0
    fy := y0

    rv := QuinticSpline{ax, bx, cx, dx, ex, fx, ay, by, cy, dy, ey, fy}

    return &rv
}

func (this* QuinticSpline) GetPose(t float64) Pose {
    x := this.Ax*math.Pow(t, 5) + this.Bx*math.Pow(t, 4) + this.Cx*math.Pow(t, 3) + this.Dx*math.Pow(t, 2) + this.Ex*t + this.Fx
    y := this.Ay*math.Pow(t, 5) + this.By*math.Pow(t, 4) + this.Cy*math.Pow(t, 3) + this.Dy*math.Pow(t, 2) + this.Ey*t + this.Fy
    dx := 5*this.Ax*math.Pow(t, 4) + 4*this.Bx*math.Pow(t, 3) + 3*this.Cx*math.Pow(t, 2) + 2*this.Dx*t + this.Ex
    dy := 5*this.Ay*math.Pow(t, 4) + 4*this.By*math.Pow(t, 3) + 3*this.Cy*math.Pow(t, 2) + 2*this.Dy*t + this.Ey
    theta := math.Atan2(dy, dx)
    return Pose{x,y,theta}
}

func (this* QuinticSpline) GetCurvature(t float64) float64 {
    return 0.0
}

func (this *QuinticSpline) String() string {
    return fmt.Sprintf("x(t) = %.2ft^5 + %.2ft^4 + %.2ft^3 + %.2ft^2 + %.2ft + %.2f\ny(t) = %.2ft^5 + %.2ft^4 + %.2ft^3 + %.2ft^2 + %.2ft + %.2f", 
    this.Ax, this.Bx, this.Cx, this.Dx, this.Ex, this.Fx, this.Ay, this.By, this.Cy, this.Dy, this.Ey, this.Fy)
}
