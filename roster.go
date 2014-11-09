package solver

type Roster struct {
	Players []Player
}

func HowManyAtPosition(position string) int {
	return len(playersByPosition[position])
}

func GetPositions() []string {
	positions := make([]string, 1)
	for position := range playersByPosition {
		positions = append(positions, position)
	}
	return positions
}

//Returns int of number of rosters created from entire population without validation of constraints
func CreateRosters() int {
	//loop over all keys of player population positions and players of those positions
	//for position := range playersByPosition {

	//}
	return 0
}

