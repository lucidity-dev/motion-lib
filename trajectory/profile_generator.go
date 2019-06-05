package trajectory

import (
	"math"
)

type Config struct {
	Dt     float64
	MaxVel float64
	Acc    float64
}

type ProfileGoal struct {
	Pos float64
	Vel float64
}

func GenerateProfile(config Config, initialState State, goal ProfileGoal) MotionProfile {
	var rv MotionProfile

	initialState = State{initialState.Pos, math.Min(initialState.Vel, config.MaxVel), math.Min(initialState.Acc, config.Acc), initialState.Time}
	dist := goal.Pos - initialState.Pos
	clampedMaxVel := math.Sqrt(initialState.Vel*initialState.Vel+goal.Vel*goal.Vel)/2.0 + dist*config.Acc
	cruiseVel := math.Min(clampedMaxVel, config.MaxVel)
	accTime := (cruiseVel - initialState.Vel) / config.Acc
	decTime := (cruiseVel - goal.Vel) / config.Acc
	accDist := (initialState.Vel + cruiseVel) / 2.0 * accTime
	decDist := (cruiseVel + goal.Vel) / 2.0 * decTime
	cruiseDist := dist - accDist - decDist
	cruiseTime := cruiseDist / cruiseVel
	totalTime := accTime + cruiseTime + decTime

	iters := (int)(totalTime / config.Dt)
	currTime := 0.0
	for i := 0; i < iters; i++ {
		var currPos, currVel, currAcc float64

		if currTime < accTime {
			currPos = initialState.Pos + initialState.Vel*currTime + 0.5*config.Acc*currTime*currTime
			currVel = initialState.Vel + config.Acc*currTime
			currAcc = config.Acc
		} else if currTime < (accTime + cruiseTime) {
			currPos = initialState.Pos + accDist + cruiseVel*(currTime-accTime)
			currVel = cruiseVel
			currAcc = 0.0
		} else {
			revCurrTime := totalTime - currTime
			relCurrTime := currTime - accTime - cruiseTime
			relCurrPos := goal.Vel*revCurrTime + 0.5*config.Acc*revCurrTime*revCurrTime

			currPos = goal.Pos - relCurrPos
			currVel = cruiseVel - config.Acc*relCurrTime
			currAcc = -config.Acc
		}

		s := State{currPos, currVel, currAcc, currTime}
		rv.AddState(s)
		currTime += config.Dt
	}

	return rv
}
