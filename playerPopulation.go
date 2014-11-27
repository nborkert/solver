package solver

import (
	"strconv"
	"strings"
	"fmt"
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
etc. for each position and player
*/
var AllPlayers [][]Player

func CreatePlayersArrays() [][]Player {
	/* Below commented-out code does not return consistent ordering of arrays.
	This is by design in Go.
	for position := range playersByPosition {
		AllPlayers = append(AllPlayers, playersByPosition[position])
	}
	*/
	fmt.Printf("Making AllPlayers, current %v\n", AllPlayers)
	AllPlayers := make([][]Player, 9)
	
	fmt.Printf("Making AllPlayers, current %v\n", AllPlayers)
	AllPlayers[0] = playersByPosition["QB"]


	fmt.Printf("Making AllPlayers, current %v\n", AllPlayers)
	AllPlayers[1] = playersByPosition["RB1"]
	AllPlayers[2] = playersByPosition["RB2"]
	AllPlayers[3] = playersByPosition["WR1"]
	AllPlayers[4] = playersByPosition["WR2"]
	AllPlayers[5] = playersByPosition["WR3"]
	AllPlayers[6] = playersByPosition["TE"]
	AllPlayers[7] = playersByPosition["K"]
	AllPlayers[8] = playersByPosition["D"]
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
	//playerToAdd.Salary, _ = strconv.ParseInt(playerData[4], 0, 64)

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
