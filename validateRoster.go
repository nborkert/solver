package solver

import (
//"fmt"
)

// Assumes not choosing D and K positions
//var salaryCap int64 = 40000
var salaryCap int64 = 50000

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

func UnderSalaryCap(roster []Player, max int64) bool {
	var total int64
	total = 0
	for _, player := range roster {
		total += player.Salary
		if total > max {
			//			fmt.Printf("Overcap of %v for roster %v\n", max, roster)
			return false
		}
	}
	//fmt.Printf("Undercap of %v for roster %v\n", max, roster)
	return true
}
