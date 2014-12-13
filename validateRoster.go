package solver

const salaryCap int = 60000
const minWinningRosterSalary int = 59500 //This is the minimum salary expected for winning rosters
const minPlayerSalary int = 4500         //4500 on real data
const maxPlayerSalary int = 9000         //9000 on real data

func EraseRosterAfterLevel(roster []Player, level int) {
	for i := range roster {
		if i > level {
			roster[i] = Player{"", "", "", 0.0, 0}
		}
	}
}

func UnderSalaryCap(roster []Player, max int) bool {
	total := RosterSalary(roster)
	if total > max {
		return false
	}
	return true
}

func RosterSalary(roster []Player) int {
	var total int
	total = 0
	for _, player := range roster {
		total += player.Salary
	}
	return total
}

func PointsForRoster(roster []Player) float64 {
	var points float64 = 0.0
	for _, player := range roster {
		points += player.ProjectedPoints
	}
	return points
}
