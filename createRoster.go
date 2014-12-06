package solver

import (
	"fmt"
)

func CreateRosters(minPoints float64) []Player {
	c := make(chan []Player)       //c is the channel used to send rosters to later processing
	workComplete := make(chan int) //workComplete is the channel used to send an int to indicate
	//that all candidate rosters have been sent from a given goroutine
	var i = 0 //will be used to count the number of goroutines launched

	for _, player := range AllPlayers[0] {
		go CreateFootballRosters(player, c, workComplete, minPoints)
		i++
	}
	return FindWinningRoster(c, workComplete, i)
}

//This is the attempt at improving roster creation and evaluation. Big change is not building
//all rosters before sending to the channel that will keep the winner. We send the roster
//immediately after being built. This should remove the memory limit.
//Hard-coded with assumptions that a roster will have these positions:
//QB, RB1, RB2, WR1, WR2, WR3, TE, K, D.
//The salaryCap var is set in ValidateRoster.go
func CreateFootballRosters(rootNode Player, c chan []Player, workComplete chan int, minPoints float64) {
	testRoster := make([]Player, 9)
	winningPoints := 0.0
	testRosterPoints := 0.0
	winningRoster := make([]Player, 9)

	testRoster[0] = rootNode

	for rb1Idx := range AllPlayers[1] {
		testRoster[1] = AllPlayers[1][rb1Idx]

		for rb2Idx := range AllPlayers[2] {
			if rb2Idx <= rb1Idx {
				continue
			}
			testRoster[2] = AllPlayers[2][rb2Idx]

			for wr1Idx := range AllPlayers[3] {
				testRoster[3] = AllPlayers[3][wr1Idx]
				//Experiments have shown that this optimization is overcome by the overhead
				//of performing the test after we reach 15 players
				if wr1Idx < 17 {
					EraseRosterAfterLevel(testRoster, 3)
					//check salary with QB, RB1, RB2, and WR1. If already too expensive,
					//try next WR1.
					if !UnderSalaryCap(testRoster, SalaryCapAtLevel(3)) {
						continue
					}
					//check salary again for minimum level
					if UnderSalaryCap(testRoster, SalaryMinAtLevel(3)) {
						continue
					}
				}
				for wr2Idx := range AllPlayers[4] {
					if wr2Idx <= wr1Idx {
						continue
					}

					testRoster[4] = AllPlayers[4][wr2Idx]
					if wr2Idx < 17 {
						EraseRosterAfterLevel(testRoster, 4)
						if !UnderSalaryCap(testRoster, SalaryCapAtLevel(4)) {
							continue
						}
						if UnderSalaryCap(testRoster, SalaryMinAtLevel(4)) {
							continue
						}
					}

					for wr3Idx := range AllPlayers[5] {
						if wr3Idx <= wr2Idx {
							continue
						}
						testRoster[5] = AllPlayers[5][wr3Idx]
						if wr3Idx < 17 {
							EraseRosterAfterLevel(testRoster, 5)
							if !UnderSalaryCap(testRoster, SalaryCapAtLevel(5)) {
								continue
							}
							if UnderSalaryCap(testRoster, SalaryMinAtLevel(5)) {
								continue
							}
						}
						for teIdx := range AllPlayers[6] {
							testRoster[6] = AllPlayers[6][teIdx]

							for kIdx := range AllPlayers[7] {
								testRoster[7] = AllPlayers[7][kIdx]

								for dIdx := range AllPlayers[8] {
									testRoster[8] = AllPlayers[8][dIdx]

									if UnderSalaryCap(testRoster, salaryCap) {
										if !DuplicatePlayersFound(testRoster) {
											//Now test to see if this roster
											//has the most points yet
											testRosterPoints = PointsForRoster(testRoster)
											if testRosterPoints > minPoints && RosterSalary(testRoster) > 59500 {
												fmt.Printf("%v,%v,%v\n", testRosterPoints, RosterSalary(testRoster), testRoster)
											}
											if testRosterPoints > winningPoints {
												winningPoints = testRosterPoints
												//winningRoster = testRoster THIS doesn't make a safe copy, seems to retain the pointer
												//winningRoster = append(winningRoster, testRoster...)
												copy(winningRoster, testRoster)
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
	}
	c <- winningRoster
	workComplete <- 1
}

//This function indicates what the salary cap is after adding the player at the indicated level.
//Level 0 is the QB, level 8 is the D.
func SalaryCapAtLevel(level int) int {
	return (salaryCap - ((8 - level) * minPlayerSalary))
}

//This function indicates the minimum salary a roster must have after adding the player at
//the given level. If salary is not the minimum, it doesn't have a chance to beat the
//expected minimum salary for a winning roster.
func SalaryMinAtLevel(level int) int {
	return (minWinningRosterSalary - ((8 - level) * maxPlayerSalary))
}
