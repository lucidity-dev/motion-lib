package path

import (
    "fmt"
    "math"
)

type Point struct {
    X, Y, Theta float64
}

func (this *Point) String() string {
    return fmt.Sprintf("(%.2f, %.2f, %.2f)", this.X, this.Y, this.Theta)
}

func Dist(start Point, end Point) float64 {
    dx := end.X - start.X
    dy := end.Y - start.Y
    rv := math.Sqrt(dx*dx + dy*dy)
    return rv
}

func Cos(p Point) float64 {
    return math.Cos(p.Theta)
}

func Sin(p Point) float64 {
    return math.Sin(p.Theta)
}
