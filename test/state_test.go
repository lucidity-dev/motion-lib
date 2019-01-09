package test;

import (
    "testing"
    "github.com/lucidity-dev/motion-lib/trajectory"
)

func TestEqual(t *testing.T){
    s1 := trajectory.State{15.2,250.1,250,0}
    s2 := trajectory.State{15.2,250.1,250,0}
    assertEqual(t, s1, s2)
}

func TestNotEqual(t *testing.T){
    s1 := trajectory.State{}
    s2 := trajectory.State{15,250,250,0}
    assertNotEqual(t, s1, s2)
}

func TestString(t *testing.T){
    s := trajectory.State{1.12, 2.2432, 3.357, 4.4}
    //(Time, Pos, Vel, Accel) all rounded to 2 decimals
    expected := "(4.40, 1.12, 2.24, 3.36)"
    assertEqual(t, s.String(), expected)
}
