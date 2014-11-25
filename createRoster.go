package solver

import (
//"fmt"
)

func CreateRosters() []Player {
	c := make(chan []Player)       //c is the channel used to send rosters to later processing
	workComplete := make(chan int) //workComplete is the channel used to send an int to indicate
	//that all candidate rosters have been sent from a given goroutine
	var i = 0 //will be used to count the number of goroutines launched

	for _, player := range AllPlayers[0] {
		//go CreateRostersForRootNode(player, c, workComplete)
		go CreateFanDuelRosters(player, c, workComplete)
		i++
	}
	return FindWinningRoster(c, workComplete, i)
}

//This is the attempt at improving roster creation and evaluation. Big change is not building
//all rosters before sending to the channel that will keep the winner. We send the roster
//immediately after being built. This should remove the memory limit.
//Hard-coded with assumptions that no K or D is being picked and we start at the RB1 position.
func CreateFanDuelRosters(rootNode Player, c chan []Player, workComplete chan int) {
	//var winningRoster [7]Player
	winningRoster := make([]Player, 7)
	//var testRoster [7]Player
	testRoster := make([]Player, 7)
	salaryCheckRoster := make([]Player, 4)
	//var salaryCheckRoster [4]Player
	winningPoints := 0.0
	var salaryCap int64 = 50000

	for rb1Idx := range AllPlayers[1] {
		for rb2Idx := range AllPlayers[2] {
			for wr1Idx := range AllPlayers[3] {
				//check salary with QB, RB1, RB2, and WR1. If under $36k, move on to next WR2
				salaryCheckRoster[0] = rootNode
				salaryCheckRoster[1] = AllPlayers[1][rb1Idx]
				salaryCheckRoster[2] = AllPlayers[2][rb2Idx]
				salaryCheckRoster[3] = AllPlayers[3][wr1Idx]
				if UnderSalaryCap(salaryCheckRoster, 36000) {
					for wr2Idx := range AllPlayers[4] {
						for wr3Idx := range AllPlayers[5] {
							for teIdx := range AllPlayers[6] {
								testRoster[0] = rootNode
								testRoster[1] = AllPlayers[1][rb1Idx]
								testRoster[2] = AllPlayers[2][rb2Idx]
								testRoster[3] = AllPlayers[3][wr1Idx]
								testRoster[4] = AllPlayers[4][wr2Idx]
								testRoster[5] = AllPlayers[5][wr3Idx]
								testRoster[6] = AllPlayers[6][teIdx]

								if UnderSalaryCap(testRoster, salaryCap) {
									if !DuplicatePlayersFound(testRoster) {
										//Now test to see if this roster
										//has the most points yet
										if PointsForRoster(testRoster) > winningPoints {
											winningPoints = PointsForRoster(testRoster)
											winningRoster = testRoster
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	c <- winningRoster
	workComplete <- 1
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
		newRosters := make([][]Player, 0) //newRosters will be built,
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
