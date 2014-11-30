package solver

import (
	"fmt"
	"testing"
)

func TestSimplex(t *testing.T) {
	playerA := Player{"QB", "Peyton", "DEN", 5.0, 20000}
	playerB := Player{"QB", "Peyton", "DEN", 5.0, 20000}
	playerC := Player{"RB1", "McCoy", "PHI", 3.0, 20000}

//	playerF := Player{"RB1", "McCoy", "PHI", 3.0, 20000}

//	playerG := Player{"RB2", "McCoy", "PHI", 3.0, 20000}

//	playerD := Player{"RB2", "McCoy", "PHI", 3.0, 20000}

	playerE := Player{"WR", "A", "B", 1.0, 1000}

	roster := make([]Player, 0)
	roster = append(roster, playerA)
	roster = append(roster, playerB)
	roster = append(roster, playerC)
	roster = append(roster, playerE)

	fmt.Printf("testing Simplex solver for roster: %v\n", roster)
	ret := Solve()
	if ret == nil {
		t.Errorf("Failed test of Simplex solver for roster %v\n", roster)
	}

}
