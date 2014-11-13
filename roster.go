package solver

import (
	"time"
	"fmt"
)

func CreateRosters() []Player {
	c := make(chan []Player) //c is the channel used to send rosters to later processing
	workComplete := make(chan int) //workComplete is the channel used to send an int to indicate
		//that all candidate rosters have been sent from a given goroutine
	var i = 0 //will be used to count the number of goroutines launched

	for _, player := range AllPlayers[0] {
		go CreateRostersForRootNode(player, c, workComplete)
		i++
	}
	return FindWinningRoster(c, workComplete, i)
}

//"previousRosters" and "newRosters" variables in this method should be considered as a one-dimensional
//list of rosters. Each roster is an array of Player structs, thus the use of []Player.
func CreateRostersForRootNode(rootNode Player, c chan []Player, workComplete chan int) {
	rootRoster := make([]Player, 0)
	rootRoster = append(rootRoster, rootNode)

fmt.Printf("AllPlayers = %v\n", AllPlayers)

	var previousRosters [][]Player
	previousRosters = append(previousRosters, rootRoster)
	//Start at row[1] of AllPlayers since we are given the root node in this function
	for level := 1; level < len(AllPlayers); level++ {
		//var newRosters [][]Player
		newRosters := make([][]Player, 0)
		for _, player := range AllPlayers[level] {
			fmt.Printf("Len of newRosters should be zero every time: %v\n", len(newRosters))
			fmt.Printf("Len of previousRosters should never be zero: %v\n", len(previousRosters))
	//		fmt.Printf("Adding player %v to previous rosters\n", player)
			//Now grab previously created rosters and add their players to newRoster
			fmt.Printf("Now adding this player to previous rosters: %v\n", player)
			for i := range previousRosters {
				fmt.Printf("previousRoster = %v\n", previousRosters[i])
				newRoster := append(previousRosters[i], player)
	//			fmt.Printf("Appended player %v to new roster\n", player)
				newRosters = append(newRosters, newRoster)
				fmt.Printf("newRosters are now %v\n", newRosters)
			}
		}
		previousRosters = make([][]Player, 0)  //clean out previousRosters
		previousRosters = newRosters
		fmt.Printf("previousRosters are now %v\n", previousRosters)

	}
	//Validate roster composition and salary cap info
	var validRosters [][]Player
	for _, roster := range previousRosters {
		fmt.Printf("Roster to validate %v\n", roster)
		isValidRoster := ValidateRoster(roster)
		if isValidRoster != nil {
			validRosters = append(validRosters, isValidRoster)
		}
	}
	fmt.Printf("Found number of valid rosters for %v - %v\n", rootNode, len(validRosters))
	//Send each valid roster to channel to find winner
	for _, validRosterMightWin := range validRosters {
		c <- validRosterMightWin
	}

	//Send "completed work" indicator after all possible rosters have been sent
	workComplete <- 1
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
		case <-time.After(time.Second * 3600): //One hour timeout value
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
