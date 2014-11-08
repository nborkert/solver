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


}
