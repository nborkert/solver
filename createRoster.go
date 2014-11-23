package solver

import (
	//"fmt"
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

	var previousRosters [][]Player
	previousRosters = append(previousRosters, rootRoster)
	//Start at row[1] of AllPlayers since we are given the root node in this function
	for level := 1; level < len(AllPlayers); level++ {
		newRosters := make([][]Player, 0)  //newRosters will be built,
			//then will replace previousRosters on each iteration.

		for _, player := range AllPlayers[level] {
			for i := range previousRosters {
				//Below line was actually overwriting the previousRosters array while creating 
				//newRoster
				//newRoster := append(previousRosters[i], player)
				newRoster := make([]Player, 1)
				newRoster[0] = player
				newRoster = append(newRoster, previousRosters[i]...)
				newRosters = append(newRosters, newRoster)
			}
		}
		previousRosters = newRosters
	}
	//Validate roster composition and salary cap info
	var validRosters [][]Player
	for _, roster := range previousRosters {
		validRoster := ValidateRoster(roster)
		if validRoster != nil {
			validRosters = append(validRosters, validRoster)
		}
	}
	//Send each valid roster to channel to find winner
	for _, validRosterMightWin := range validRosters {
		c <- validRosterMightWin
	}

	//Send "completed work" indicator after all possible rosters have been sent
	workComplete <- 1
}


