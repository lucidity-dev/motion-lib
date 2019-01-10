package trajectory;

import (
    "fmt"
    "bytes"
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/plotutil"
    "gonum.org/v1/plot/vg"
)

type MotionProfile struct{
    states []State
}

func (this *MotionProfile) Length() int {
    return len(this.states)
}

func (this *MotionProfile) AddState(state State) {
    this.states = append(this.states, state)
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

func (this *MotionProfile) Plot() {
    p, err := plot.New()
    if err != nil {
        panic(err)
    }
    p.Title.Text = "Motion Profile"
    p.X.Label.Text = "Time"
    p.Y.Label.Text = "Position, Velocity, Acceleration"

    posPoints := make(plotter.XYs, this.Length())
    velPoints := make(plotter.XYs, this.Length())
    accPoints := make(plotter.XYs, this.Length())

    for i := 0; i < this.Length(); i += 1 {
        posPoints[i].X = this.states[i].Time
        velPoints[i].X = this.states[i].Time
        accPoints[i].X = this.states[i].Time

        posPoints[i].Y = this.states[i].Pos
        velPoints[i].Y = this.states[i].Vel
        accPoints[i].Y = this.states[i].Acc
    }

    err = plotutil.AddLinePoints(p, "Position", posPoints, "Velocity", velPoints, "Acceleration", accPoints)
    if err != nil {
        panic(err)
    }

    err = p.Save(12*vg.Inch, 12*vg.Inch, "MotionProfile.jpg")
    if err != nil {
        panic(err)
    }
}
