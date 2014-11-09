package solver

import (
	"strings"
	"strconv"
)

type Player struct {
	Position string
	PlayerName string
	Team string
	ProjectedPoints float64
	Salary int64
}

var playersByPosition map[string][]Player
//map of position names to a slice of Player structs

func CreatePlayer (dataLine string) Player {
	//Expect format QB,Peyton Manning,DEN,10,30000. Projected points could be decimal format
	playerData := strings.Split(dataLine, ",")
	playerToAdd := Player{}
	playerToAdd.Position = playerData[0]
	playerToAdd.PlayerName = playerData[1]
	playerToAdd.Team = playerData[2]
	playerToAdd.ProjectedPoints, _ = strconv.ParseFloat(playerData[3], 64)
	playerToAdd.Salary, _ = strconv.ParseInt(playerData[4], 0, 64)

	return playerToAdd
}

func AddPlayerToPopulation (playerToAdd Player) int {
	if playersByPosition == nil {
		playersByPosition = make(map[string][]Player)
	}
	playersByPosition[playerToAdd.Position] = append(playersByPosition[playerToAdd.Position], playerToAdd)
	return len(playersByPosition[playerToAdd.Position])
}
