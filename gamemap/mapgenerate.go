package gamemap

import (
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

func (gm *GameMap) DeleteDoors(currentRoom []int) []int {
	//delete doors
	if gm.LeftDestination == nil {
		currentRoom[112] = 0
		currentRoom[128] = 0
	}
	if gm.UpDestination == nil {
		currentRoom[7] = 0
		currentRoom[8] = 0
	}
	if gm.RightDestination == nil {
		currentRoom[127] = 0
		currentRoom[143] = 0
	}
	if gm.DownDestination == nil {
		currentRoom[247] = 0
		currentRoom[248] = 0
	}
	return currentRoom
}

func (gm *GameMap) CloseDoors(currentRoom []int) []int {
	//delete doors
	if !(gm.LeftDestination == nil) {
		currentRoom[112] = 5
		currentRoom[128] = 5
	}
	if !(gm.UpDestination == nil) {
		currentRoom[7] = 5
		currentRoom[8] = 5
	}
	if !(gm.RightDestination == nil) {
		currentRoom[127] = 5
		currentRoom[143] = 5
	}
	if !(gm.DownDestination == nil) {
		currentRoom[247] = 5
		currentRoom[248] = 5
	}
	return currentRoom
}

func (gm *GameMap) OpenDoors(currentRoom []int) []int {
	//delete doors
	if !(gm.LeftDestination == nil) {
		currentRoom[112] = 2
		currentRoom[128] = 2
	}
	if !(gm.UpDestination == nil) {
		currentRoom[7] = 2
		currentRoom[8] = 2
	}
	if !(gm.RightDestination == nil) {
		currentRoom[127] = 2
		currentRoom[143] = 2
	}
	if !(gm.DownDestination == nil) {
		currentRoom[247] = 2
		currentRoom[248] = 2
	}
	return currentRoom
}

func (gm *GameMap) GenerateMap(numberOfCommonRooms, numberOfBossRooms, numberOfShopRooms, numberOfChestRooms int) (*GameMap, [](*GameMap)) {
	//minimap generation
	minimap := [10][10]int{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			minimap[i][j] = 0
		}
	}
	//currentPointX := rand.Intn(5) + 2
	//currentPointY := rand.Intn(5) + 2
	currentPointX := 5
	currentPointY := 5
	minimap[currentPointX][currentPointY] = 101
	potencial := [](*Neighbors){}
	numberOfRooms := numberOfBossRooms + numberOfChestRooms + numberOfCommonRooms + numberOfShopRooms
	//get ID List
	CommonRoomsIDList := GetRoomIDList("./gamemap/assets/commonrooms.json")
	IDList := GetRoomIDList("./gamemap/assets/commonrooms.json")
	for i := 0; i < numberOfRooms; i++ {
		//fmt.Println(currentPointX)
		//fmt.Println(currentPointY)
		//Left potencial
		if (currentPointX > 0) && (minimap[currentPointX-1][currentPointY] == 0) {
			potencialNeighbor := &Neighbors{
				X: currentPointX - 1,
				Y: currentPointY,
			}
			potencial = append(potencial, potencialNeighbor)
			//fmt.Println("Left")
		}
		//Up potencial
		if (currentPointY < 9) && (minimap[currentPointX][currentPointY+1] == 0) {
			potencialNeighbor := &Neighbors{
				X: currentPointX,
				Y: currentPointY + 1,
			}
			potencial = append(potencial, potencialNeighbor)
			//fmt.Println("Up")
		}
		//Right potencial
		if (currentPointX < 9) && (minimap[currentPointX+1][currentPointY] == 0) {
			potencialNeighbor := &Neighbors{
				X: currentPointX + 1,
				Y: currentPointY,
			}
			potencial = append(potencial, potencialNeighbor)
			//fmt.Println("Right")
		}
		//Down potencial
		if (currentPointY > 0) && (minimap[currentPointX][currentPointY-1] == 0) {
			potencialNeighbor := &Neighbors{
				X: currentPointX,
				Y: currentPointY - 1,
			}
			potencial = append(potencial, potencialNeighbor)
			//fmt.Println("Down")
		}
		//fmt.Println(IDList)
		//Choose New Point
		fmt.Println(len(potencial))
		rand1 := rand.Intn(len(potencial)-2) + 1
		currentPointX = potencial[rand1].X
		currentPointY = potencial[rand1].Y
		potencial = append(potencial[:rand1-1], potencial[rand1+1:]...)
		if numberOfCommonRooms != 0 {
			rand2 := rand.Intn(len(CommonRoomsIDList))
			minimap[currentPointX][currentPointY] = CommonRoomsIDList[rand2]
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
					MapX:   i,
					MapY:   j,
					RoomID: minimap[i][j],
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
	StartRoom := &GameMap{}
	for i := 0; i < len(GameRoomList); i++ {
		if GameRoomList[i].RoomID == 101 {
			StartRoom = GameRoomList[i]
		}
	}
	return StartRoom, GameRoomList
}
