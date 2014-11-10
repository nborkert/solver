package solver

import (
	"testing"
)

func ValidateRosterTest(t* testing.T) {
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

	var salaryCap int64
	salaryCap = 50000
	if UnderSalaryCap(roster, salaryCap) {
		t.Errorf("Failed test of salary cap check for roster %v\n", roster)
	}

	fullCheck := ValidateRoster(goodRoster)
	if fullCheck == nil {
		t.Errorf("Failed test of ValidateRoster for roster %v\n", goodRoster)
	}
}
