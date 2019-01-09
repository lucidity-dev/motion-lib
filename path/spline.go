package path

type CubicSpline struct{
    //Spline coefficients
    //x(t) = Ax*t^3 + Bx*t^2 + Cx*t + Dx
    //y(t) = Ay*t^3 + By*t^2 + Cy*t + Dy
    Ax,Bx,Cx,Dx,Ay,By,Cy,Dy float64
}
