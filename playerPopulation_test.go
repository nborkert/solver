package solver

import (
	//	"fmt"
	"testing"
)

func TestAddPlayer(t *testing.T) {
	newPlayer := "QB,Peyton Manning,DEN,18,35000"
	playerAdded := CreatePlayer(newPlayer)
	if playerAdded.Position != "QB" {
		t.Errorf("Expected QB, received %v", playerAdded.Position)
	}
	if playerAdded.ProjectedPoints != 18 {
		t.Errorf("Expected 18, recieved %v", playerAdded.ProjectedPoints)
	}
	AddPlayerToPopulation(playerAdded)

	player2 := "RB1,Yo,SEA,32.33,40000"
	playerAdded2 := CreatePlayer(player2)
	if playerAdded2.PlayerName != "Yo" {
		t.Error("Expected name of Yo, received %v", playerAdded2.PlayerName)
	}
	if playerAdded2.Salary != 40000 {
		t.Error("Expected salary of 40000, received %v", playerAdded2.Salary)
	}
	if playerAdded2.Team != "SEA" {
		t.Error("Expected team of SEA, received %v", playerAdded2.Team)
	}

	testResult := AddPlayerToPopulation(playerAdded2)
	if testResult != 1 {
		t.Error("Expected 1, received %v", testResult)
	}

	player3 := "RB2,Yo2,SEA,32.33,40000"
	playerAdded3 := CreatePlayer(player3)
	AddPlayerToPopulation(playerAdded3)

	player4 := "WR1,x,x,1.1,1"
	player5 := "WR2,x,x,1.1,1"
	player6 := "WR3,x,x,1.1,1"
	player7 := "TE,y,x,1.1,1"
	player8 := "K,x,x,1.1,1"
	player9 := "D,x,x,1.1,1"

	p4 := CreatePlayer(player4)
	p5 := CreatePlayer(player5)
	p6 := CreatePlayer(player6)
	p7 := CreatePlayer(player7)
	p8 := CreatePlayer(player8)
	p9 := CreatePlayer(player9)
	AddPlayerToPopulation(p4)

	AddPlayerToPopulation(p5)
	AddPlayerToPopulation(p6)
	AddPlayerToPopulation(p7)
	AddPlayerToPopulation(p8)
	AddPlayerToPopulation(p9)

	allPlayers := CreatePlayersArrays()
	if allPlayers == nil {
		t.Error("Expected non-nil, received nil for CreatePlayerArrays")
	}

	if len(allPlayers[0]) != 1 {
		t.Errorf("Expected 1 QB, %v found\n", len(allPlayers[0]))
	}

	if len(allPlayers[1]) != 1 {
		t.Errorf("Expected 1 RB1, %v found\n", len(allPlayers[1]))
	}

	roster := AddPlayerToSingleList(p4)
	roster = AddPlayerToSingleList(p5)
	roster = AddPlayerToSingleList(p7)
	//fmt.Printf("roster = %v\n", roster)
	if roster == nil {
		t.Error("Expected non-ni, received nil for roster")
	}
	if len(roster) != 3 {
		t.Error("Expected 3 members of roster, didn't find 3")
	}

	WRRoster := AddPlayerToWRList(p4)
	WRRoster = AddPlayerToWRList(p5)
	if len(WRRoster) != 2 {
		t.Error("Expected 2 members of WRList, didn't receive 2. Check AddPlayerToWRList function.")
	}
}
