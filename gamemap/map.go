package gamemap

import (
	"errors"
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameMap struct {
	MapX                    int
	MapY                    int
	RoomID                  int
	LeftDestination         &GameMap
	UpDestination         &GameMap
	RightDestination         &GameMap
	DownDestination         &GameMap
	TileSize                int
	SreenWidth, SreenHeight int
	tileset                 *ebiten.Image
}
type GameMapOptions struct {
	TileSize                int
	SreenWidth, SreenHeight int
	tileset                 *ebiten.Image
}
type Neighbors struct{
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

func (GM *GameMap) GetCurrentRoomID()(int) {
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

func GenerateMap(numberOfCommonRooms,numberOfBossRooms,numberOfShopRooms,numberOfChestRooms int)
{
	minimap:= [][]int
	for int i := 0 ; i < 10; i++{
			for int j := 0 ; j < 10; j++{
				minimap[i][j]=0
			}
	}
	currentPointX := random(2,7)
	currentPointY := random(2,7)
	minimap[currentPointX][currentPointY]=101;
	potencial:= []&Neighbors
	numberOfRooms:= numberOfBossRooms+numberOfChestRooms+numberOfCommonRooms+numberOfShopRooms
	for i=0;i<numberOfRooms;i++{
		//Left potencial
		if ((currentPointX - 1)&&(currentPointY))==0{
			potencialNeighbor:=Neighbors{
				X: currentPointX-1,
				Y: currentPointY,
			}
			potencial:=append(potencial,potencialNeighbor)
		}
		//Up potencial
		if ((currentPointX)&&(currentPointY+1))==0{
			potencialNeighbor:=Neighbors{
				X: currentPointX,
				Y: currentPointY+1,
			}
			potencial:=append(potencial,potencialNeighbor)
		}
		//Right potencial
		if ((currentPointX+1)&&(currentPointY))==0{
			potencialNeighbor:=Neighbors{
				X: currentPointX+1,
				Y: currentPointY,
			}
			potencial:=append(potencial,potencialNeighbor)
		}
		//Down potencial
		if ((currentPointX)&&(currentPointY-1))==0{
			potencialNeighbor:=Neighbors{
				X: currentPointX,
				Y: currentPointY-1,
			}
			potencial:=append(potencial,potencialNeighbor)
		}
		//Choose New Point
		rand1:= random(0,potencial.lenght)
		currentPointX = potencial[rand1].X
		currentPointY = potencial[rand1].Y
		splice(potencial,rand1,1)
		if (numberOfCommonRooms!=0){
			rand2:=random(0,IDList.length)
			minimap[currentPointX][currentPointY]=IDList[rand2]
			numberOfCommonRooms=numberOfCommonRooms-1
		}else{
			if (numberOfChestRooms!=0){
			rand2:=random(0,IDList.length)
			minimap[currentPointX][currentPointY]=IDList[rand2]
			numberOfChestRooms=numberOfChestRooms-1
			}else{
				if (numberOfShopRooms!=0){
					rand2:=random(0,IDList.length)
					minimap[currentPointX][currentPointY]=IDList[rand2]
					numberOfShopRooms=numberOfShopRooms-1
				}else{

					if (numberOfBossRooms!=0){
						rand2:=random(0,IDList.length)
						minimap[currentPointX][currentPointY]=IDList[rand2]
						numberOfBossRooms=numberOfBossRooms-1

					}
				}
			}
		}
	}
}
