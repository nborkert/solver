package solver

import (
	"fmt"
)

func HowManyAtPosition(position string) int {
	return len(playersByPosition[position])
}

func GetPositions() []string {
	positions := make([]string, 1)
	for position := range playersByPosition {
		positions = append(positions, position)
	}
	return positions
}

func CreateRosters() {
	c := make(chan []Player) //c is the channel used to send rosters to later processing
	for position := range playersByPosition {
		for _, player := range playersByPosition[position] {
			fmt.Printf("About to call CreateRostersForRootNode for player %v\n", player)
			go CreateRostersForRootNode(player, c)
		}
		break  //only needed that first entry from playersByPosition
	}
	ValidateRoster(c)
	ValidateRoster(c)
}

func CreateRostersForRootNode(rootNode Player, c chan []Player) {
	fmt.Printf("Creating rosters for root %v\n", rootNode)

	roster := make([]Player, 0)
	roster = append(roster, rootNode)

	fmt.Printf("About to call ValidateRoster for roster %v\n", roster)
	//Send to channel for each possible roster
	c <- roster
}

//This method is the centralized "pulling" function that evaluates all possible rosters
//and finds the highest projected winner subject to the expressed constraints
func ValidateRoster(c chan []Player) {
	fmt.Printf("In ValidateRoster\n")
	roster := <-c

	fmt.Printf("Validating roster %v\n", roster)
}

