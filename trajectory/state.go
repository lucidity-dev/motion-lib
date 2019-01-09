package trajectory;

import (
    "fmt"
)

type State struct{
    Pos float64
    Vel float64
    Acc float64
    Time float64
}

func (this *State) String() string{
    //(time, position, velocity, acceleration)
    return fmt.Sprintf("(%.2f, %.2f, %.2f, %.2f)", this.Time, this.Pos, this.Vel, this.Acc)
}
