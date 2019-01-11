package path

import (
    "fmt"
    "math"
)

type Pose struct {
    X, Y, Theta float64
}

func (this *Pose) String() string {
    return fmt.Sprintf("(%.2f, %.2f, %.2f)", this.X, this.Y, this.Theta)
}

func Dist(start Pose, end Pose) float64 {
    dx := end.X - start.X
    dy := end.Y - start.Y
    rv := math.Sqrt(dx*dx + dy*dy)
    return rv
}

func Cos(p Pose) float64 {
    return math.Cos(p.Theta)
}

func Sin(p Pose) float64 {
    return math.Sin(p.Theta)
}
