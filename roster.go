package solver

import (
//	"fmt"
	"time"
)

func CreateRosters() []Player {
	c := make(chan []Player) //c is the channel used to send rosters to later processing
	workComplete := make(chan int) //workComplete is the channel used to send an int to indicate
		//that all candidate rosters have been sent from a given goroutine
	var i = 0 //will be used to count the number of goroutines launched

	for _, player := range AllPlayers[0] {
//		fmt.Printf("About to call CreateRostersForRootNode for player %v\n", player)
		go CreateRostersForRootNode(player, c, workComplete)
		i++
	}
	return FindWinningRoster(c, workComplete, i)
}


func CreateRostersForRootNode(rootNode Player, c chan []Player, workComplete chan int) {
//	fmt.Printf("Creating rosters for root %v\n", rootNode)

	rootRoster := make([]Player, 0)
	rootRoster = append(rootRoster, rootNode)
	//fmt.Printf("Root roster = %v\n", rootRoster)

	var previousRosters [][]Player
	previousRosters = append(previousRosters, rootRoster)
	//fmt.Printf("previousRosters = %v\n", previousRosters)

	//Start at row[1] of AllPlayers since we are given the root node in this function
	for level := 1; level < len(AllPlayers); level++ {
	//	fmt.Printf("level = %v\n", level)
		var newRosters [][]Player
		for _, player := range AllPlayers[level] {
	//		fmt.Printf("PLAYER %v\n", player)
			//Now grab previously created rosters and add their players to newRoster
			for previousRoster := range previousRosters {
				newRoster := append(previousRosters[previousRoster], player)
		//		fmt.Printf("NEW ROSTER AFTER ADDING PREVIOUS GUYS %v\n", newRoster)
				newRosters = append(newRosters, newRoster)
			}
		}
		//fmt.Printf("newRosters = %v\n", newRosters)
		previousRosters = newRosters
	}
	//Validate roster composition and salary cap info
	var validRosters [][]Player
	for _, roster := range previousRosters {
		isValidRoster := ValidateRoster(roster)
		if isValidRoster != nil {
			validRosters = append(validRosters, isValidRoster)
		}
	}
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
			//fmt.Printf("Checking for winner for roster %v\n", roster)
			rosterPoints := PointsForRoster(roster)
			//fmt.Printf("Roster has points = %v\n", rosterPoints)
			if rosterPoints > highestPoints {
			//	fmt.Printf("HIGHEST SO FAR\n")
				highestPoints = rosterPoints
				winningRoster = roster
			}
			//Now compare and keep winning roster to send back
		case done := <-workComplete:
			completedWorkers = completedWorkers + done
			//fmt.Printf("completedWorkers = %v\n", completedWorkers)
			if waitForWorkerCount == completedWorkers {
			//	fmt.Printf("WINNING ROSTER = %v\n", winningRoster)
				return winningRoster
			}
		case <-time.After(time.Second * 3600):
			//fmt.Printf("Timed out\n")
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
