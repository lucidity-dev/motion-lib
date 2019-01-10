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

    start := path.Point{0,0,math.Pi/2.0}
    end := path.Point{1,1,math.Pi/2.0}
    spline := path.GenerateCubicSpline(start, end)

    fmt.Println(spline.String())

    spline.Plot()
}
