package solver

import (
	"fmt"
	"time"
)

func CreateRosters() {
	c := make(chan []Player) //c is the channel used to send rosters to later processing
	workComplete := make(chan int) //workComplete is the channel used to send an int to indicate
		//that all candidate rosters have been sent from a given goroutine
	var i = 0 //will be used to count the number of goroutines launched

	for _, player := range AllPlayers[0] {
		fmt.Printf("About to call CreateRostersForRootNode for player %v\n", player)
		go CreateRostersForRootNode(player, c, workComplete)
		i++
	}
	ValidateRosters(c, workComplete, i)
}

func CreateRostersForRootNode(rootNode Player, c chan []Player, workComplete chan int) {
	fmt.Printf("Creating rosters for root %v\n", rootNode)

	roster := make([]Player, 0)
	roster = append(roster, rootNode)

	fmt.Printf("About to call ValidateRoster for roster %v\n", roster)
	//Send each possible roster to channel
	c <- roster

	//Send "completed work" indicator after all possible rosters have been sent
	workComplete <- 1
}

//This method is the centralized "pulling" function that evaluates all possible rosters
//and finds the highest projected winner subject to the expressed constraints
func ValidateRosters(c chan []Player, workComplete chan int, waitForWorkerCount int) {
	fmt.Printf("In ValidateRosters\n")
	completedWorkers := 0

	for {
		select {
		case roster := <-c:
			fmt.Printf("Validating roster %v\n", roster)
		case done := <-workComplete:
			completedWorkers = completedWorkers + done
			fmt.Printf("completedWorkers = %v\n", completedWorkers)
			if waitForWorkerCount == completedWorkers {
				return
			}
		case <-time.After(time.Second * 3):
			fmt.Printf("Timed out\n")
			return
		}
	}
}

