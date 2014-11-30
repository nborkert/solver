package solver

import (
	//"fmt"
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

	//fmt.Printf("testing Simplex solver for roster: %v\n", roster)
	//ret := Solve()
	//if ret == nil {
	//	t.Errorf("Failed test of Simplex solver for roster %v\n", roster)
	//	}

	//CREATE ARRAYS, call CreateSimplexTableaux, then show() to compare tableaus to Simplex.java
	/*
		c := []float64{13.0, 23.0}
		b := []float64{480.0, 160.0, 1190.0}

		//A := [][]float64{{5.0, 15.0}, {4.0, 4.0}, {35.0, 20.0},}
		A := make([][]float64, 3)
		for i := range A {
			A[i] = make([]float64, 2)
		}
		A[0][0] = 5.0
		A[0][1] = 15.0
		A[1][0] = 4.0
		A[1][1] = 4.0
		A[2][0] = 35.0
		A[2][1] = 20.0
	*/
	//	createSimplexTableaux(A, b, c)
	//	show()
	roster = CreateSimplexRoster()
}
