package trajectory;

import (
    "fmt"
    "bytes"
)

type MotionProfile struct{
    states []State
}

func (this *MotionProfile) Length() int {
    return len(this.states)
}

func AddState(mp *MotionProfile, state State) {
    mp.states = append(mp.states, state)
}

func (this *MotionProfile) String() string{
    var buffer bytes.Buffer

    buffer.WriteString(fmt.Sprintf("Total length: %d\n", this.Length()))
    buffer.WriteString("(Time, Pos, Vel, Accel)\n")

    for _, state := range this.states{
        buffer.WriteString(state.String())
        buffer.WriteString("\n")
    }

    return buffer.String()
}
