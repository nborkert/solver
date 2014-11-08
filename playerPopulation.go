package solver

import (
	"strings"
	//"fmt"
	"strconv"
)

type Player struct {
	Position string
	PlayerName string
	Team string
	ProjectedPoints float64
	Salary int64
}

func AddPlayer (dataLine string) (string, float64) {
	//Expect format QB,Peyton Manning,DEN,10,30000
//	fmt.Printf("%q\n", strings.Split(dataLine, ","))
	playerData := strings.Split(dataLine, ",")
	playerToAdd := Player{}
	playerToAdd.Position = playerData[0]
	playerToAdd.PlayerName = playerData[1]
	playerToAdd.Team = playerData[2]
	playerToAdd.ProjectedPoints, _ = strconv.ParseFloat(playerData[3], 64)
	playerToAdd.Salary, _ = strconv.ParseInt(playerData[4], 0, 64)


	return playerToAdd.Position, playerToAdd.ProjectedPoints

}
