package trajectory;

import (
    "fmt"
    "bytes"
)

type MotionProfile struct{
    points []Point
}

func (this *MotionProfile) Length() int {
    return len(this.points)
}

func AddPoint(mp *MotionProfile, point *Point) {
    mp.points = append(mp.points, *point)
}

func (this *MotionProfile) String() string{
    var buffer bytes.Buffer

    buffer.WriteString(fmt.Sprintf("Total length: %d", this.Length()))
    buffer.WriteString("(Time, Pos, Vel, Accel)\n")

    for _, point  := range this.points{
        buffer.WriteString(point.String())
        buffer.WriteString("\n")
    }

    return buffer.String()
}
