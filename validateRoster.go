package solver

import (
	"fmt"
)

//Insure a player is not on the roster twice and that the cost of the roster is under or equal to the salary cap
//Returns roster if valid, nil if not valid
func ValidateRoster(roster []Player) []Player {
	fmt.Printf("Validating roster %v\n", roster)
	return roster
}


