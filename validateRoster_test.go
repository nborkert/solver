package solver

import (
	"testing"
)

func TestInsureNoDuplicatePlayers(t* testing.T) {
	playerA := Player{"QB", "Peyton", "DEN", 5.0, 20000}
	playerB := Player{"QB", "Peyton", "DEN", 5.0, 20000}
	playerC := Player{"RB", "McCoy", "PHI", 3.0, 20000}

	roster := make([]Player, 0)
	roster = append(roster, playerA)
	roster = append(roster, playerB)
	roster = append(roster, playerC)

	ret := NoDuplicatePlayersFound(roster)
	if ret {
		t.Errorf("Failed test of duplicate player roster validation for roster %v\n", roster)
	}

	goodRoster := make([]Player, 0)
	goodRoster = append(goodRoster, playerA)
	goodRoster = append(goodRoster, playerC)
	goodDupCheck := NoDuplicatePlayersFound(goodRoster)
	if !goodDupCheck {
		t.Errorf("Failed test of dup player roster validation for roster %v\n", goodRoster)
	}




}
