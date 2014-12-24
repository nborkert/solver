package solver

import (
	//	"fmt"
	"testing"
)

func TestValidateRoster(t *testing.T) {
	playerA := Player{"QB", "Peyton", "DEN", 5.0, 20000}
	playerB := Player{"QB", "Peyton", "DEN", 5.0, 20000}
	playerC := Player{"RB1", "McCoy", "PHI", 3.0, 20000}
	playerD := Player{"RB2", "McCoy", "PHI", 3.0, 20000}
	playerE := Player{"WR", "A", "B", 1.0, 1000}
	playerF := Player{"RB1", "McCoy", "PHI", 3.0, 20000}
	playerG := Player{"RB2", "McCoy", "PHI", 3.0, 20000}

	roster := make([]Player, 0)
	roster = append(roster, playerA)
	roster = append(roster, playerB)
	roster = append(roster, playerC)
	roster = append(roster, playerE)

	salaryForRoster := RosterSalary(roster)
	//	fmt.Printf("Salary = %v\n", salaryForRoster)
	if salaryForRoster != 61000 {
		t.Errorf("Failed to properly calculate salary. Check RosterSalary function for errors.")
	}

	underSalaryCap := UnderSalaryCap(roster, 61000)
	if !underSalaryCap {
		t.Errorf("Failed test of UnderSalaryCap check.")
	}

	points := PointsForRoster(roster)
	if points != 14.0 {
		t.Errorf("Failed test of PointsForRoster calculation.")
	}


	roster3 := make([]Player, 6)
	roster3[0] = playerA
	roster3[1] = playerC
	roster3[2] = playerF
	roster3[3] = playerG
	roster3[4] = playerD
	roster3[5] = playerE

	EraseRosterAfterLevel(roster3, 2)
	if roster3[3].Position == "RB2" || roster3[4].Position != "" {
		t.Errorf("Failed to erase roster members. Check EraseRosterAfterLevel function for errors.")
	}

}
