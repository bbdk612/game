package gamemap

import (
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
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
type GameMapOptions struct {
	TileSize                int
	SreenWidth, SreenHeight int
	tileset                 *ebiten.Image
}
type Neighbors struct {
	X int
	Y int
}

func (GM *GameMap) CheckDirection(direction string) (int, bool) {
	chunk, ok := GM.roadsTo[GM.currentChunk][direction]
	if ok {
		return chunk, true
	} else {
		return -1, false
	}
}

func (GM *GameMap) GetCurrentRoomID() int {
	return GM.RoomID
}

func (GM *GameMapOptions) GetTile(tileNumber int) *ebiten.Image {
	w := GM.tileset.Bounds().Dx()
	tileXCount := w / GM.TileSize

	tileStartX := (tileNumber % tileXCount) * GM.TileSize
	tileStartY := (tileNumber / tileXCount) * GM.TileSize

	return GM.tileset.SubImage(image.Rect(tileStartX, tileStartY, tileStartX+GM.TileSize, tileStartY+GM.TileSize)).(*ebiten.Image)
}

func InitGameMap(chunks [][]int, currentChunk int, roadsTo []map[string]int, sreenWidth int, sreenHeight int) (*GameMap, error) {
	tilesetFile, err := os.Open("./assets/tileset.png")
	if err != nil {
		return nil, err
	}

	tileset, _, err := image.Decode(tilesetFile)

	if err != nil {
		return nil, err
	}

	tilesImage := ebiten.NewImageFromImage(tileset)

	GM := &GameMap{
		chunks:       chunks,
		roadsTo:      roadsTo,
		SreenWidth:   sreenWidth,
		SreenHeight:  sreenHeight,
		TileSize:     16,
		currentChunk: currentChunk,
		tileset:      tilesImage,
	}
	return GM, nil
}

func GenerateMap(numberOfCommonRooms, numberOfBossRooms, numberOfShopRooms, numberOfChestRooms int) {
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
	potencial := [](*Neighbors)
	numberOfRooms := numberOfBossRooms + numberOfChestRooms + numberOfCommonRooms + numberOfShopRooms
	for i = 0; i < numberOfRooms; i++ {
		//Left potencial
		if ((currentPointX - 1) == 0) && ((currentPointY) == 0) {
			potencialNeighbor := Neighbors{
				X: currentPointX - 1,
				Y: currentPointY,
			}
			potencial := append(potencial, potencialNeighbor)
		}
		//Up potencial
		if ((currentPointX) == 0) && ((currentPointY + 1) == 0) {
			potencialNeighbor := Neighbors{
				X: currentPointX,
				Y: currentPointY + 1,
			}
			potencial := append(potencial, potencialNeighbor)
		}
		//Right potencial
		if ((currentPointX + 1) == 0) && ((currentPointY) == 0) {
			potencialNeighbor := Neighbors{
				X: currentPointX + 1,
				Y: currentPointY,
			}
			potencial := append(potencial, potencialNeighbor)
		}
		//Down potencial
		if ((currentPointX) == 0) && ((currentPointY-1) == 0) == 0 {
			potencialNeighbor := Neighbors{
				X: currentPointX,
				Y: currentPointY - 1,
			}
			potencial := append(potencial, potencialNeighbor)
		}
		//Choose New Point
		rand1 := rand.Intn(len(potencial))
		currentPointX = potencial[rand1].X
		currentPointY = potencial[rand1].Y
		splice(potencial, rand1, 1)
		if numberOfCommonRooms != 0 {
			rand2 := random(0, len(IDList))
			minimap[currentPointX][currentPointY] = IDList[rand2]
			numberOfCommonRooms = numberOfCommonRooms - 1
		} else {
			if numberOfChestRooms != 0 {
				rand2 := random(0, len(IDList))
				minimap[currentPointX][currentPointY] = IDList[rand2]
				numberOfChestRooms = numberOfChestRooms - 1
			} else {
				if numberOfShopRooms != 0 {
					rand2 := random(0, len(IDList))
					minimap[currentPointX][currentPointY] = IDList[rand2]
					numberOfShopRooms = numberOfShopRooms - 1
				} else {

					if numberOfBossRooms != 0 {
						rand2 := random(0, len(IDList))
						minimap[currentPointX][currentPointY] = IDList[rand2]
						numberOfBossRooms = numberOfBossRooms - 1

					}
				}
			}
		}
	}
	//map GenerateMap
	GameRoomList := []*GameMap
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if minimap[i][j] != 0 {
				NewRoom := *GameMap{
					MapX:   i - 1,
					MapY:   j,
					RoomID: minimap[i-1][j],
				}
				GameRoomList = append(GameRoomList, NewRoom)
			}
		}
	}
	for i := 0; i < len(GameRoomList)-1; i++ {
		for j := i + 1; j < len(GameRoomList); j++ {
			//Left Doors
			if (GameRoomList[i].MapX-1 == GameRoomList[j].MapX) && (GameRoomList[i].MapY == GameRoomList[j].MapY) {
				GameRoomList[i].LeftDestination = GameRoomList[j]
				GameRoomList[j].RigthDestination = GameRoomList[i]
			}
			//Up Doors
			if (GameRoomList[i].MapX == GameRoomList[j].MapX) && (GameRoomList[i].MapY+1 == GameRoomList[j].MapY) {
				GameRoomList[i].UpDestination = GameRoomList[j]
				GameRoomList[j].DownDestination = GameRoomList[i]
			}
			//Right Doors
			if (GameRoomList[i].MapX+1 == GameRoomList[j].MapX) && (GameRoomList[i].MapY == GameRoomList[j].MapY) {
				GameRoomList[i].RigthDestination = GameRoomList[j]
				GameRoomList[j].LeftDestination = GameRoomList[i]
			}
			//Down Doors
			if (GameRoomList[i].MapX == GameRoomList[j].MapX) && (GameRoomList[i].MapY-1 == GameRoomList[j].MapY) {
				GameRoomList[i].DownDestination = GameRoomList[j]
				GameRoomList[j].UpDestination = GameRoomList[i]
			}
		}
	}
}
