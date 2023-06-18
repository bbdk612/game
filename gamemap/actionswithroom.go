package gamemap

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type RoomData struct {
	Data []int
	Id   int
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
