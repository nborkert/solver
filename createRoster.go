package solver

import (
//	"fmt"
)

func CreateRosters() []Player {
	c := make(chan []Player)       //c is the channel used to send rosters to later processing
	workComplete := make(chan int) //workComplete is the channel used to send an int to indicate
	//that all candidate rosters have been sent from a given goroutine
	var i = 0 //will be used to count the number of goroutines launched

	for _, player := range AllPlayers[0] {
		go CreateFootballRosters(player, c, workComplete)
		i++
	}
	return FindWinningRoster(c, workComplete, i)
}

//This is the attempt at improving roster creation and evaluation. Big change is not building
//all rosters before sending to the channel that will keep the winner. We send the roster
//immediately after being built. This should remove the memory limit.
//Hard-coded with assumptions that no K or D is being picked and we start at the RB1 position.
//The salaryCap var is set in ValidateRoster.go
func CreateFootballRosters(rootNode Player, c chan []Player, workComplete chan int) {
	testRoster := make([]Player, 7)
	salaryCheckRoster := make([]Player, 4)
	winningPoints := 0.0
	testRosterPoints := 0.0
	winningRoster := make([]Player, 7)

	testRoster[0] = rootNode
	salaryCheckRoster[0] = rootNode

	for rb1Idx := range AllPlayers[1] {
		testRoster[1] = AllPlayers[1][rb1Idx]
		salaryCheckRoster[1] = AllPlayers[1][rb1Idx]

		for rb2Idx := range AllPlayers[2] {
			testRoster[2] = AllPlayers[2][rb2Idx]
			salaryCheckRoster[2] = AllPlayers[2][rb2Idx]

			for wr1Idx := range AllPlayers[3] {
				testRoster[3] = AllPlayers[3][wr1Idx]
				salaryCheckRoster[3] = AllPlayers[3][wr1Idx]

				//check salary with QB, RB1, RB2, and WR1. If under $36k, move on to next position 
				if UnderSalaryCap(salaryCheckRoster, 36000) {
					for wr2Idx := range AllPlayers[4] {
						testRoster[4] = AllPlayers[4][wr2Idx]

						for wr3Idx := range AllPlayers[5] {
							testRoster[5] = AllPlayers[5][wr3Idx]

							for teIdx := range AllPlayers[6] {
								testRoster[6] = AllPlayers[6][teIdx]

								if UnderSalaryCap(testRoster, salaryCap) {
									if !DuplicatePlayersFound(testRoster) {
										//Now test to see if this roster
										//has the most points yet
								//		fmt.Printf("Found no dups for this roster above%v\n", testRoster)
										testRosterPoints = PointsForRoster(testRoster)
										if testRosterPoints > winningPoints {
											winningPoints = testRosterPoints
											//winningRoster = testRoster THIS doesn't make a safe copy, seems to retain the pointer
											//winningRoster = append(winningRoster, testRoster...)
											copy(winningRoster, testRoster)
										//	fmt.Printf("testRoster = %v\n", testRoster)
										//	fmt.Printf("winningRoster = %v\n", winningRoster)
										//	fmt.Printf("WinningPoints so far = %v\n", winningPoints)
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
	//fmt.Printf("Winning points = %v\n", winningPoints)
	//fmt.Printf("Winning roster = %v\n", winningRoster)
	c <- winningRoster
	workComplete <- 1
}

