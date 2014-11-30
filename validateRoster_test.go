package solver

import (
//	"fmt"
	"testing"
)

func TestValidateRoster(t *testing.T) {
	playerA := Player{"QB", "Peyton", "DEN", 5.0, 20000}
	playerB := Player{"QB", "Peyton", "DEN", 5.0, 20000}
	playerC := Player{"RB1", "McCoy", "PHI", 3.0, 20000}

	playerF := Player{"RB1", "McCoy", "PHI", 3.0, 20000}

	playerG := Player{"RB2", "McCoy", "PHI", 3.0, 20000}

	playerD := Player{"RB2", "McCoy", "PHI", 3.0, 20000}

	playerE := Player{"WR", "A", "B", 1.0, 1000}

	roster := make([]Player, 0)
	roster = append(roster, playerA)
	//roster = append(roster, playerB)
	roster = append(roster, playerC)
	roster = append(roster, playerE)

	roster = append(roster, playerB)

//	fmt.Printf("testing roster for dup: %v\n", roster)
	ret := DuplicatePlayersFound(roster)
	if !ret {
		t.Errorf("Failed test of duplicate player roster validation for roster %v\n", roster)
	}

	salaryForRoster := RosterSalary(roster)
//	fmt.Printf("Salary = %v\n", salaryForRoster)
	if salaryForRoster != 61000 {
		t.Errorf("Failed salary calculation")
	}

	underSalaryCap := UnderSalaryCap(roster, 61000)
	if !underSalaryCap {
		t.Errorf("Failed salary cap calculation")
	}

	roster2 := make([]Player, 0)
	roster2 = append(roster2, playerC)
	roster2 = append(roster2, playerD)
	roster2 = append(roster2, playerA)
	dupCheck := DuplicatePlayersFound(roster2)
//	fmt.Printf("Roster2 = %v\n", roster2)
	if !dupCheck {
		t.Errorf("Failed test of duplicate player roster validation for roster %v\n", roster2)
	}

	goodRoster := make([]Player, 0)
	goodRoster = append(goodRoster, playerA)
	goodRoster = append(goodRoster, playerC)
	goodRoster = append(goodRoster, playerE)
	goodDupCheck := DuplicatePlayersFound(goodRoster)
//	fmt.Printf("goodRoster = %v\n", goodRoster)
	if goodDupCheck {
		t.Errorf("Failed test of dup player roster validation for roster %v\n", goodRoster)
	}

	roster3 := make([]Player, 6)
	roster3[0] = playerA
	roster3[1] = playerC
	roster3[2] = playerF
	roster3[3] = playerG
	roster3[4] = playerD
	roster3[5] = playerE

//	fmt.Printf("testing roster3 for dup: %v\n", roster3)
	if !DuplicatePlayersFound(roster3) {
		t.Errorf("AHHHH %v\n", roster3)
	}

	EraseRosterAfterLevel(roster3, 2)
//	fmt.Printf("roster3 after erasing after level 2 = %v\n", roster3)

}
