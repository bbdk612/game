package gamemap

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type RoomData struct {
	TileMap []int
	id      int
}

func GetRoomIDList() []int {
	IDList := []int{}
	IDListRead, err := os.ReadFile("./gamemap/assets/101.json")
	RD := [](*RoomData){}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(IDListRead, &RD)
	for i := 0; i < len(RD); i++ {
		IDList = append(IDList, RD[i].id)
	}
	return IDList
}

func JsonFileDecodeCurrentRoom(currentRoomID int) *RoomData {
	data, err := os.ReadFile("./gamemap/assets/101.json")
	RD := [](*RoomData){}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &RD)
	for i := 0; i < len(RD); i++ {
		if RD[i].id == currentRoomID {
			return RD[i]
		}
	}
	log.Fatal(err)
	return nil
}

func (rd *RoomData) GetCurrentRoomTileMap() []int {
	return rd.TileMap
}
