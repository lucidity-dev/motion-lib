package trajectory;

import (
    "bytes"
)

type MotionProfile struct{
    points []Point
    length int
}

func (this *MotionProfile) Length() int {
    return this.length;
}

func AddPoint(mp *MotionProfile, point *Point) {
    mp.points = append(mp.points, *point)
    mp.length += 1
}

func (this *MotionProfile) String() string{
    var buffer bytes.Buffer

    buffer.WriteString("(Time, Pos, Vel, Accel)\n")

    for _, point  := range this.points{
        buffer.WriteString(point.String())
        buffer.WriteString("\n")
    }

    return buffer.String()
}
