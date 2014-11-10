package solver

import (
//	"fmt"
)

var salaryCap int64 = 60000

//Insure a player is not on the roster twice and that the cost of the roster is under or equal to the salary cap
//Returns roster if valid, nil if not valid
func ValidateRoster(roster []Player) []Player {
//	fmt.Printf("Validating roster %v\n", roster)
	dupTest := NoDuplicatePlayersFound(roster)
	if !dupTest {
		return nil
	}

	salaryCapTest := UnderSalaryCap(roster, salaryCap)
	if !salaryCapTest {
		return nil
	}

	return roster
}

// Players are deemed duplicate if they have the same name and team and appear in different positions in the roster array
func NoDuplicatePlayersFound(roster []Player) bool {
	for basePos, basePlayer := range roster {
		for movingPos, movingPlayer := range roster {
//			fmt.Printf("basePlayer = %v\n", basePlayer)
//			fmt.Printf("movingPlayer = %v\n", movingPlayer)
			if (movingPlayer.PlayerName == basePlayer.PlayerName) && (movingPlayer.Team == basePlayer.Team) && (basePos != movingPos) {
//				fmt.Printf("FOUND DUPLICATE PLAYER\n")
				return false
			}
		}
	}
	return true
}

func UnderSalaryCap(roster []Player, cap int64) bool {
//	fmt.Printf("Checking roster salary against cap of %v\n", cap)
	var total int64
	total = 0
	for _, player := range roster {
		total += player.Salary
		if (total > cap) {
			return false
		}
	}
	return true
}

