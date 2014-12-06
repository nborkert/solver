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

	salaryForRoster := RosterSalary(roster)
	//	fmt.Printf("Salary = %v\n", salaryForRoster)
	if salaryForRoster != 61000 {
		t.Errorf("Failed salary calculation")
	}

	underSalaryCap := UnderSalaryCap(roster, 61000)
	if !underSalaryCap {
		t.Errorf("Failed salary cap calculation")
	}

	roster3 := make([]Player, 6)
	roster3[0] = playerA
	roster3[1] = playerC
	roster3[2] = playerF
	roster3[3] = playerG
	roster3[4] = playerD
	roster3[5] = playerE

	EraseRosterAfterLevel(roster3, 2)
	//	fmt.Printf("roster3 after erasing after level 2 = %v\n", roster3)

}
