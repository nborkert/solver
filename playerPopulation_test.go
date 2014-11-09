package solver

import (
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

	player2 := "RB,Yo,SEA,32.33,40000"

	playerAdded2 := CreatePlayer(player2)
	if playerAdded2.Position != "RB" {
		t.Error("Expected RB, received %v", playerAdded2.Position)
	}
	if playerAdded2.ProjectedPoints != 32.33 {
		t.Error("Expected 32.33, received %v", playerAdded2.ProjectedPoints)
	}

	testResult := AddPlayerToPopulation(playerAdded2)
	if testResult != 1 {
		t.Error("Expected 1, received %v", testResult)
	}

}
