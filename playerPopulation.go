package solver

import (
	"strconv"
	"strings"
)

type Player struct {
	Position        string
	PlayerName      string
	Team            string
	ProjectedPoints float64
	Salary          int
}

//map of position names to a slice of Player structs
var playersByPosition map[string][]Player

/* we need an array of arrays of Player structs to have reliable behavior when creating rosters later
This should be mentally considered as a two-dimensional layout of players by position like:
QBs [Peyton Manning] [Tom Brady] [Carson Palmer]
RB1 [Lesean McCoy]   [Demarco Murray]
RB2 [Lesean McCoy]   [Demarco Murray]
etc. for each position and player in this order: QB, RB1, RB2, WR1, WR2, WR3, TE, K, D
*/
var AllPlayers [][]Player

var SingleList []Player

func AddPlayerToSingleList(toAdd Player) []Player {
	SingleList = append(SingleList, toAdd)
	return SingleList
}

var WRList []Player

func AddPlayerToWRList(toAdd Player) []Player {
	WRList = append(WRList, toAdd)
	return WRList
}

func CreatePlayersArrays() [][]Player {
	/* Below commented-out code does not return consistent ordering of arrays.
	This is by design in Go.
	for position := range playersByPosition {
		AllPlayers = append(AllPlayers, playersByPosition[position])
	}
	*/
	AllPlayers = append(AllPlayers, playersByPosition["QB"])
	AllPlayers = append(AllPlayers, playersByPosition["RB1"])
	AllPlayers = append(AllPlayers, playersByPosition["RB2"])
	AllPlayers = append(AllPlayers, playersByPosition["WR1"])
	AllPlayers = append(AllPlayers, playersByPosition["WR2"])
	AllPlayers = append(AllPlayers, playersByPosition["WR3"])
	AllPlayers = append(AllPlayers, playersByPosition["TE"])
	AllPlayers = append(AllPlayers, playersByPosition["K"])
	AllPlayers = append(AllPlayers, playersByPosition["D"])

	return AllPlayers

}

func CreatePlayer(dataLine string) Player {
	//Expect format Position,PlayerName,Team,ProjectedPoints,Salary. ProjectedPoints could be decimal format.
	//Example: QB,Peyton Manning,DEN,10,30000
	playerData := strings.Split(dataLine, ",")
	playerToAdd := Player{}
	playerToAdd.Position = playerData[0]
	playerToAdd.PlayerName = playerData[1]
	playerToAdd.Team = playerData[2]
	playerToAdd.ProjectedPoints, _ = strconv.ParseFloat(playerData[3], 64)
	playerToAdd.Salary, _ = strconv.Atoi(playerData[4])
	return playerToAdd
}

func AddPlayerToPopulation(playerToAdd Player) int {
	if playersByPosition == nil {
		playersByPosition = make(map[string][]Player)
	}
	playersByPosition[playerToAdd.Position] = append(playersByPosition[playerToAdd.Position], playerToAdd)
	return len(playersByPosition[playerToAdd.Position])
}
