package path

import (
    "fmt"
)

type Point struct {
    X, Y, Theta float64
}

func (this *Point) String() string {
    return fmt.Sprintf("(%.2f, %.2f, %.2f)", this.X, this.Y, this.Theta)
}
