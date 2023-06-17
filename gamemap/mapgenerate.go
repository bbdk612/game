package gamemap

import (
	"crypto/x509"
	"fmt"
	"math/rand"
)

type GameMap struct {
	MapX             int
	MapY             int
	RoomID           int
	LeftDestination  *GameMap
	UpDestination    *GameMap
	RightDestination *GameMap
	DownDestination  *GameMap
}
type Neighbors struct {
	X int
	Y int
}

func (GM *GameMap) GetCurrentRoomID() int {
	return GM.RoomID
}
func (GM *GameMap) ChangeCurrentRoom(direction string) *GameMap {
	switch direction {
	case "left":
		CurrentRoom := GM.LeftDestination
		return CurrentRoom

	case "right":
		CurrentRoom := GM.RightDestination
		return CurrentRoom

	case "top":
		CurrentRoom := GM.UpDestination
		return CurrentRoom

	case "down":
		CurrentRoom := GM.DownDestination
		return CurrentRoom
	}
	return nil
}

func (gm *GameMap) DeleteDoors() int {
	currentRoomID := gm.GetCurrentRoomID
	//take json file
	//delete doors
	if gm.LeftDestination == nil {
		x = 1
		y = 1
	}
	if gm.RightDestination == nil {
		x = 1
		y = 1
	}
	if gm.UpDestination == nil {
		x = 1
		y = 1
	}
	if gm.DownDestination == nil {
		x = 1
		y = 1
	}
	return currentRoom
}

func (gm *GameMap) CloseDoors() int {
	currentRoomID := gm.GetCurrentRoomID
	//take json file
	//delete doors
	if !(gm.LeftDestination == nil) {
		currentRoom[x] = 5
		currentRoom[y] = 5
	}
	if !(gm.RightDestination == nil) {
		currentRoom[x] = 5
		currentRoom[y] = 5
	}
	if !(gm.UpDestination == nil) {
		currentRoom[x] = 5
		currentRoom[y] = 5
	}
	if !(gm.DownDestination == nil) {
		currentRoom[x] = 5
		currentRoom[y] = 5
	}
	return currentRoom
}

func (gm *GameMap) OpenDoors() int {
	currentRoomID := gm.GetCurrentRoomID
	//take json file
	//delete doors
	if !(gm.LeftDestination == nil) {
		currentRoom[x] = 2
		currentRoom[y] = 2
	}
	if !(gm.RightDestination == nil) {
		currentRoom[x] = 2
		currentRoom[y] = 2
	}
	if !(gm.UpDestination == nil) {
		currentRoom[x] = 2
		currentRoom[y] = 2
	}
	if !(gm.DownDestination == nil) {
		currentRoom[x] = 2
		currentRoom[y] = 2
	}
	return currentRoom
}

func (gm *GameMap) GenerateMap(numberOfCommonRooms, numberOfBossRooms, numberOfShopRooms, numberOfChestRooms int) {
	//minimap generation
	minimap := [][]int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			minimap[i][j] = 0
		}
	}
	currentPointX := rand.Intn(5) + 2
	currentPointY := rand.Intn(5) + 2
	minimap[currentPointX][currentPointY] = 101
	potencial := [](*Neighbors){}
	numberOfRooms := numberOfBossRooms + numberOfChestRooms + numberOfCommonRooms + numberOfShopRooms
	for i := 0; i < numberOfRooms; i++ {
		//Left potencial
		if ((currentPointX - 1) == 0) && ((currentPointY) == 0) {
			potencialNeighbor := &Neighbors{
				X: currentPointX - 1,
				Y: currentPointY,
			}
			potencial = append(potencial, potencialNeighbor)
		}
		//Up potencial
		if ((currentPointX) == 0) && ((currentPointY + 1) == 0) {
			potencialNeighbor := &Neighbors{
				X: currentPointX,
				Y: currentPointY + 1,
			}
			potencial = append(potencial, potencialNeighbor)
		}
		//Right potencial
		if ((currentPointX + 1) == 0) && ((currentPointY) == 0) {
			potencialNeighbor := &Neighbors{
				X: currentPointX + 1,
				Y: currentPointY,
			}
			potencial = append(potencial, potencialNeighbor)
		}
		//Down potencial
		if ((currentPointX) == 0) && ((currentPointY - 1) == 0) {
			potencialNeighbor := &Neighbors{
				X: currentPointX,
				Y: currentPointY - 1,
			}
			potencial = append(potencial, potencialNeighbor)
		}
		//Choose New Point
		rand1 := rand.Intn(len(potencial) - 1)
		currentPointX = potencial[rand1].X
		currentPointY = potencial[rand1].Y
		potencial = append(potencial[:rand1-1], potencial[rand1+1:])
		if numberOfCommonRooms != 0 {
			rand2 := rand.Intn(len(IDList))
			minimap[currentPointX][currentPointY] = IDList[rand2]
			numberOfCommonRooms = numberOfCommonRooms - 1
		} else {
			if numberOfChestRooms != 0 {
				rand2 := rand.Intn(len(IDList))
				minimap[currentPointX][currentPointY] = IDList[rand2]
				numberOfChestRooms = numberOfChestRooms - 1
			} else {
				if numberOfShopRooms != 0 {
					rand2 := rand.Intn(len(IDList))
					minimap[currentPointX][currentPointY] = IDList[rand2]
					numberOfShopRooms = numberOfShopRooms - 1
				} else {

					if numberOfBossRooms != 0 {
						rand2 := rand.Intn(len(IDList))
						minimap[currentPointX][currentPointY] = IDList[rand2]
						numberOfBossRooms = numberOfBossRooms - 1

					}
				}
			}
		}
	}
	//map GenerateMap
	GameRoomList := [](*GameMap){}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if minimap[i][j] != 0 {
				NewRoom := &GameMap{
					MapX:   i - 1,
					MapY:   j,
					RoomID: minimap[i-1][j],
				}
				GameRoomList = append(GameRoomList, NewRoom)
			}
		}
	}
	for i := 0; i < len(GameRoomList); i++ {
		for j := 0; j < len(GameRoomList); j++ {
			//Left Doors
			if (GameRoomList[i].MapX-1 == GameRoomList[j].MapX) && (GameRoomList[i].MapY == GameRoomList[j].MapY) {
				GameRoomList[i].LeftDestination = GameRoomList[j]
			}
			//Up Doors
			if (GameRoomList[i].MapX == GameRoomList[j].MapX) && (GameRoomList[i].MapY+1 == GameRoomList[j].MapY) {
				GameRoomList[i].UpDestination = GameRoomList[j]
			}
			//Right Doors
			if (GameRoomList[i].MapX+1 == GameRoomList[j].MapX) && (GameRoomList[i].MapY == GameRoomList[j].MapY) {
				GameRoomList[i].RightDestination = GameRoomList[j]
			}
			//Down Doors
			if (GameRoomList[i].MapX == GameRoomList[j].MapX) && (GameRoomList[i].MapY-1 == GameRoomList[j].MapY) {
				GameRoomList[i].DownDestination = GameRoomList[j]
			}
		}
	}
}