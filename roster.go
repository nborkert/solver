package solver

import (
	"time"
)

// A Roster is a list of Player structs and their combined salary
type Roster struct {
	Players []Player
	Salary  int64
}

//This method is the centralized "pulling" function that evaluates all possible rosters
//and finds the highest projected winner subject to the expressed constraints
func FindWinningRoster(c chan []Player, workComplete chan int, waitForWorkerCount int) []Player {
	completedWorkers := 0
	var highestPoints float64 = 0.0
	var winningRoster []Player
	for {
		select {
		case roster := <-c:
			rosterPoints := PointsForRoster(roster)
			//Now compare and keep winning roster to send back
			if rosterPoints > highestPoints {
				highestPoints = rosterPoints
				winningRoster = roster
			}
		case done := <-workComplete:
			completedWorkers = completedWorkers + done
			if waitForWorkerCount == completedWorkers {
				return winningRoster
			}
		case <-time.After(time.Second * 36000): //Ten hours timeout value
			return nil
		}
	}
}

func PointsForRoster(roster []Player) float64 {
	var points float64 = 0.0
	for _, player := range roster {
		points += player.ProjectedPoints
	}
	return points
}
