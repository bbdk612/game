package gamemap

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"os"
)

type GameMap struct {
	TileSize                  int
	ScreenWidth, ScreenHeight int
	tileset                   *ebiten.Image
	CurrentRoom               *GameRoom
	RoomList                  [](*GameRoom)
	MapOptions                *GameMap
	RD                        *RoomData
	MiniMapPlan               [10][10]int
	CurrentRoomIDMiniMap      int
	CurrentRoomTiles          []int
}

func (GM *GameMap) GetTile(tileNumber int) *ebiten.Image {
	w := GM.tileset.Bounds().Dx()
	tileXCount := w / GM.TileSize

	tileStartX := (tileNumber % tileXCount) * GM.TileSize
	tileStartY := (tileNumber / tileXCount) * GM.TileSize

	return GM.tileset.SubImage(image.Rect(tileStartX, tileStartY, tileStartX+GM.TileSize, tileStartY+GM.TileSize)).(*ebiten.Image)
}

func InitGameMap(tilesetImgPath string, screenWidth, screenHeight int) (*GameMap, error) {
	tilesetFile, err := os.Open(tilesetImgPath)
	if err != nil {
		return nil, err
	}

	tileset, _, err := image.Decode(tilesetFile)

	if err != nil {
		return nil, err
	}

	tilesImage := ebiten.NewImageFromImage(tileset)

	GM := &GameMap{
		TileSize:     16,
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
		tileset:      tilesImage,
	}
	return GM, nil
}
