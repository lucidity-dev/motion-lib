package main;

import (
    "fmt"

    "github.com/lucidity-dev/motion-lib/trajectory"
)

func main(){
    fmt.Println("Go motion library.....")

    config := trajectory.Config{0.01, 12, 12}

    goal := trajectory.ProfileGoal{100, 0}

    initialState := trajectory.State{10,12,0,0}

    profile := trajectory.GenerateProfile(config, initialState, goal)

    //fmt.Println(profile.String())

    profile.Plot()
}
