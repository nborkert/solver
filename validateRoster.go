package solver

import (
//"fmt"
)

const salaryCap int = 60000
const minWinningRosterSalary int = 50000 //This is the minimum salary expected for winning rosters
const minPlayerSalary int = 4500         //4500 on real data
const maxPlayerSalary int = 9000         //9000 on real data

//Insure a player is not on the roster twice and that the cost of the roster is under or equal to the salary cap
//Returns roster if valid, nil if not valid
func ValidateRoster(roster []Player) []Player {
	if DuplicatePlayersFound(roster) {
		return nil
	}
	if !UnderSalaryCap(roster, salaryCap) {
		return nil
	}
	return roster
}

// Players are deemed duplicate if they have the same name and team and appear in different positions in the roster array
func DuplicatePlayersFound(roster []Player) bool {
	for basePos, basePlayer := range roster {
		for movingPos, movingPlayer := range roster {
			if (movingPlayer.PlayerName == basePlayer.PlayerName) && (movingPlayer.Team == basePlayer.Team) && (basePos != movingPos) {
				//								fmt.Printf("FOUNDDUPONROSTER %v and %v\n", basePlayer, movingPlayer)
				return true
			}
		}
	}
	//		fmt.Printf("NODUPSFOUND on roster %v\n", roster)
	return false
}

func UnderSalaryCap(roster []Player, max int) bool {
	total := RosterSalary(roster)
	if total > max {
		//			fmt.Printf("Overcap of %v for roster %v\n", max, roster)
		return false
	}
	//fmt.Printf("Undercap of %v for roster %v\n", max, roster)
	return true
}

func RosterSalary(roster []Player) int {
	var total int
	total = 0
	for _, player := range roster {
		total += player.Salary
	}
	return total
}

func PointsForRoster(roster []Player) float64 {
	var points float64 = 0.0
	for _, player := range roster {
		points += player.ProjectedPoints
	}
	return points
}
