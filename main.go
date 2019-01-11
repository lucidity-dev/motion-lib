package main;

import (
    "fmt"
    "math"

    //"github.com/lucidity-dev/motion-lib/trajectory"
    "github.com/lucidity-dev/motion-lib/path"
)

func main(){
    fmt.Println("Go motion library.....")

    /*
    config := trajectory.Config{0.01, 12, 12}

    goal := trajectory.ProfileGoal{100, 0}

    initialState := trajectory.State{10,12,0,0}

    profile := trajectory.GenerateProfile(config, initialState, goal)

    //fmt.Println(profile.String())

    profile.Plot()
    */

    a := path.Pose{20,-95,0}
    b := path.Pose{225,-95,15.0*math.Pi/180.0}
    c := path.Pose{225,50,90.0}
    spline1 := path.GenerateQuinticSpline(a, b)
    spline2 := path.GenerateQuinticSpline(b, c)

    splines := []path.Spline{spline1, spline2}

    path.PlotSplines(splines, "", "Splines.jpg")
}
