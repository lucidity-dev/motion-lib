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

    ax := QuinticA_Pos0*x0 + QuinticA_Vel0*vx0 + QuinticA_Acc0*ax0 + QuinticA_Acc1*ax1 + QuinticA_Vel1*vx1 + QuinticA_Pos1*x1
    ay := QuinticA_Pos0*y0 + QuinticA_Vel0*vy0 + QuinticA_Acc0*ay0 + QuinticA_Acc1*ay1 + QuinticA_Vel1*vy1 + QuinticA_Pos1*y1
    bx := QuinticB_Pos0*x0 + QuinticB_Vel0*vx0 + QuinticB_Acc0*ax0 + QuinticB_Acc1*ax1 + QuinticB_Vel1*vx1 + QuinticB_Pos1*x1
    by := QuinticB_Pos0*y0 + QuinticB_Vel0*vy0 + QuinticB_Acc0*ay0 + QuinticB_Acc1*ay1 + QuinticB_Vel1*vy1 + QuinticB_Pos1*y1
    cx := QuinticC_Pos0*x0 + QuinticC_Vel0*vx0 + QuinticC_Acc0*ax0 + QuinticC_Acc1*ax1 + QuinticC_Vel1*vx1 + QuinticC_Pos1*x1
    cy := QuinticC_Pos0*y0 + QuinticC_Vel0*vy0 + QuinticC_Acc0*ay0 + QuinticC_Acc1*ay1 + QuinticC_Vel1*vy1 + QuinticC_Pos1*y1
    dx := QuinticD_Acc0*ax0
    dy := QuinticD_Acc0*ay0
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
