package solver

import (
	"testing"
	"sort"
)

//Sort is performed with max at beginning of array after sort
func TestSortStackByVal(t *testing.T) {
	//PlayerStack defined with player names, team, points, value, salary
	stackA := PlayerStack{"AA", "DEN", 5.0, 0.5, 0}
	stackB := PlayerStack{"BB", "DEN", 100.0, 10.0, 0}
	stackC := PlayerStack{"CC", "PHI", 20, 20.0, 0}

	stacks := make([]PlayerStack, 0)
	stacks = append(stacks, stackA)
	stacks = append(stacks, stackB)
	stacks = append(stacks, stackC)

	sort.Sort(ByVal(stacks))

	if stacks[0].PlayerNames != "CC" {
		t.Errorf("Did not find stack CC at highest value after sorting ByVal")
	}


	sort.Sort(ByPoints(stacks))

	if stacks[0].PlayerNames != "BB" {
		t.Errorf("Did not find stack BB at highest projected points after sorting ByPoints")
	}
}
