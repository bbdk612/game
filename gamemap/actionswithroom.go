package gamemap

import (
	"encoding/json"
	"fmt"
	"game/animatedobjects"
	"log"
	"os"
)

type RoomData struct {
	Data                 []int
	Id                   int
	NumberOfMonsters     int
	MonsterStartTiles    []int
	WeHaveChest          bool
	WeHaveWayToNextLevel bool
}

func SetCurrentRoom(CurrentRoom *GameRoom) (*GameRoom, *RoomData, []int, []*animatedobjects.Monster) {
	RD := JsonFileDecodeCurrentRoom(CurrentRoom.RoomID, "./gamemap/assets/roomlist.json")
	CurrentRoomTiles := RD.GetCurrentRoomTileMap()
	CurrentRoomTiles = CurrentRoom.DeleteDoors(CurrentRoomTiles)
	ListOfMonsters := [](*animatedobjects.Monster){}
	if RD.WeHaveChest && CurrentRoom.Chest == nil {
		ch, err := animatedobjects.InitNewChest("./assets/chest.json", 135)
		if err != nil {
			log.Fatal(err)
		}
		ch.ChestPlayer.Play("wait")
		CurrentRoom.Chest = ch
	}
	if RD.WeHaveWayToNextLevel && CurrentRoom.WayToNextLevel == nil {
		wnl, err := animatedobjects.InitNewWayToNextLevel("./assets/chest.json", 135)
		if err != nil {
			log.Fatal(err)
		}
		wnl.WNLPlayer.Play("wait")
		CurrentRoom.WayToNextLevel = wnl
	}
	if !(CurrentRoom.RoomIsCleaned) {
		if RD.NumberOfMonsters > 0 {
			ListOfMonsters = SpawnMonsters(RD)
			CurrentRoomTiles = CurrentRoom.ChangeDoorsState(CurrentRoomTiles, 4)
		} else {
			CurrentRoom.RoomIsCleaned = true
		}
	}
	return CurrentRoom, RD, CurrentRoomTiles, ListOfMonsters
}

func SpawnMonsters(RD *RoomData) []*animatedobjects.Monster {
	ListOfMonsters := [](*animatedobjects.Monster){}
	ListOfMonsters = nil
	for i := 0; i < RD.NumberOfMonsters; i++ {
		ms, er := animatedobjects.InitMonsters(RD.MonsterStartTiles[i])
		if er != nil {
			log.Fatal(er)
		}
		ms.AsePlayer.Play("left")
		ListOfMonsters = append(ListOfMonsters, ms)
	}
	return ListOfMonsters
}

func (GR *GameRoom) ChangeCurrentRoom(direction string) (*GameRoom, *RoomData, []int, []*animatedobjects.Monster) {
	switch direction {
	case "left":
		CurrentRoom := GR.LeftDestination
		return SetCurrentRoom(CurrentRoom)

	case "right":
		CurrentRoom := GR.RightDestination
		return SetCurrentRoom(CurrentRoom)

	case "top":
		CurrentRoom := GR.UpDestination
		return SetCurrentRoom(CurrentRoom)

	case "down":
		CurrentRoom := GR.DownDestination
		return SetCurrentRoom(CurrentRoom)
	}
	return nil, nil, nil, nil
}

func (GR *GameRoom) DeleteDoors(currentRoom []int) []int {
	//delete doors
	if GR.LeftDestination == nil {
		currentRoom[112] = 0
		currentRoom[128] = 0
	}
	if GR.UpDestination == nil {
		currentRoom[7] = 0
		currentRoom[8] = 0
	}
	if GR.RightDestination == nil {
		currentRoom[127] = 0
		currentRoom[143] = 0
	}
	if GR.DownDestination == nil {
		currentRoom[247] = 0
		currentRoom[248] = 0
	}
	return currentRoom
}

func (GR *GameRoom) ChangeDoorsState(currentRoom []int, doorState int) []int {
	//change state
	if !(GR.LeftDestination == nil) {
		currentRoom[112] = doorState
		currentRoom[128] = doorState
	}
	if !(GR.UpDestination == nil) {
		currentRoom[7] = doorState
		currentRoom[8] = doorState
	}
	if !(GR.RightDestination == nil) {
		currentRoom[127] = doorState
		currentRoom[143] = doorState
	}
	if !(GR.DownDestination == nil) {
		currentRoom[247] = doorState
		currentRoom[248] = doorState
	}
	return currentRoom
}

func GetRoomIDList(filePath string) []int {
	IDList := []int{}
	IDListRead, err := os.ReadFile(filePath)
	RD := [](*RoomData){}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(IDListRead, &RD)
	for i := 0; i < len(RD); i++ {
		IDList = append(IDList, RD[i].Id)
	}
	return IDList
}

func JsonFileDecodeCurrentRoom(currentRoomID int, filePath string) *RoomData {
	data, err := os.ReadFile(filePath)
	RD := [](*RoomData){}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &RD)
	for i := 0; i < len(RD); i++ {
		if RD[i].Id == currentRoomID {
			return RD[i]
		}
	}
	fmt.Println("error")
	return nil
}

func (rd *RoomData) GetCurrentRoomTileMap() []int {
	return rd.Data
}
func (GR *GameRoom) GetCurrentRoomID() int {
	return GR.RoomID
}

func (GR *GameRoom) GetCurrentRoomCoordinate() (int, int) {
	stX := GR.MapX
	stY := GR.MapY
	return stX, stY
}
