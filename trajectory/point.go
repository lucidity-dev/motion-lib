package trajectory;

import (
    "fmt"
)

type Point struct{
    Pos float32
    Vel float32
    Acc float32
    Time float32
}

func (this *Point) String() string{
    //(time, position, velocity, acceleration)
    return fmt.Sprintf("(%.2f, %.2f, %.2f, %.2f)", this.Time, this.Acc, this.Vel, this.Pos)
}
