package solver

import (
	"fmt"
	"testing"
)

func TestValidateRoster(t *testing.T) {
	playerA := Player{"QB", "Peyton", "DEN", 5.0, 20000}
	playerB := Player{"QB", "Peyton", "DEN", 5.0, 20000}
	playerC := Player{"RB1", "McCoy", "PHI", 3.0, 20000}
	playerD := Player{"RB2", "McCoy", "PHI", 3.0, 20000}
	playerE := Player{"WR", "A", "B", 1.0, 1000}

	roster := make([]Player, 0)
	roster = append(roster, playerA)
	//roster = append(roster, playerB)
	roster = append(roster, playerC)
	roster = append(roster, playerE)

	roster = append(roster, playerB)

	fmt.Printf("testing roster for dup: %v\n", roster)
	ret := DuplicatePlayersFound(roster)
	if !ret {
		t.Errorf("Failed test of duplicate player roster validation for roster %v\n", roster)
	}

	roster2 := make([]Player, 0)
	roster2 = append(roster2, playerC)
	roster2 = append(roster2, playerD)
	roster2 = append(roster2, playerA)
	dupCheck := DuplicatePlayersFound(roster2)
	fmt.Printf("Roster2 = %v\n", roster2)
	if !dupCheck {
		t.Errorf("Failed test of duplicate player roster validation for roster %v\n", roster2)
	}

	goodRoster := make([]Player, 0)
	goodRoster = append(goodRoster, playerA)
	goodRoster = append(goodRoster, playerC)
	goodRoster = append(goodRoster, playerE)
	goodDupCheck := DuplicatePlayersFound(goodRoster)
	fmt.Printf("goodRoster = %v\n", goodRoster)
	if goodDupCheck {
		t.Errorf("Failed test of dup player roster validation for roster %v\n", goodRoster)
	}
	/*
		var salaryCap int64
		salaryCap = 50000
		if UnderSalaryCap(roster, salaryCap) {
			t.Errorf("Failed test of salary cap check for roster %v\n", roster)
		}

		fullCheck := ValidateRoster(goodRoster)
		if fullCheck == nil {
			t.Errorf("Failed test of ValidateRoster for roster %v\n", goodRoster)
		}
	*/
}
