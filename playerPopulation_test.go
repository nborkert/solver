package solver

import (
	"testing"
)

func TestAddPlayer(t *testing.T) {
	newPlayer := "QB,Peyton Manning,DEN,18,35000"
	positionAdded, points := AddPlayer(newPlayer)
	if positionAdded != "QB" {
		t.Errorf("Expected QB, received %v", positionAdded)
	}
	if points != 18 {
		t.Errorf("Expected 18, recieved %v", points)
	}

	player2 := "RB,Yo,SEA,32.33,40000"

	newPositionAdded, newPoints := AddPlayer(player2)
	if newPositionAdded != "RB" {
		t.Error("Expected RB, received %v", newPositionAdded)
	}
	if newPoints != 32.33 {
		t.Error("Expected 32.33, received %v", newPoints)
	}
	
}
