package solver

import (
	"fmt"
	"time"
)

func CreateRosters() {
	c := make(chan []Player) //c is the channel used to send rosters to later processing
	for _, player := range AllPlayers[0] {
		fmt.Printf("About to call CreateRostersForRootNode for player %v\n", player)
		go CreateRostersForRootNode(player, c)
	}
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
	for {
		select {
		case roster := <-c:
			fmt.Printf("Validating roster %v\n", roster)
		case <-time.After(time.Second * 3):
			fmt.Printf("Timed out\n")
			return
		}
	}
}

