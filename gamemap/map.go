package gamemap

import (
	"errors"
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameMap struct {
	mapX                    int
	ampY                    int
	chunks                  [][]int
	roadsTo                 []map[string]int
	currentChunk            int
	TileSize                int
	SreenWidth, SreenHeight int
	tileset                 *ebiten.Image
}

func (GM *GameMap) CheckDirection(direction string) (int, bool) {
	chunk, ok := GM.roadsTo[GM.currentChunk][direction]
	if ok {
		return chunk, true
	} else {
		return -1, false
	}
}

func (GM *GameMap) GetCurrentChunk() []int {
	return GM.chunks[GM.currentChunk]
}

func (GM *GameMap) ChangeCurrentChunk(chunk int) error {
	if chunk > len(GM.chunks) {
		return errors.New("Chunk is out of range")
	}
	GM.currentChunk = chunk
	return nil
}

func (GM *GameMap) GetTile(tileNumber int) *ebiten.Image {
	w := GM.tileset.Bounds().Dx()
	tileXCount := w / GM.TileSize

	tileStartX := (tileNumber % tileXCount) * GM.TileSize
	tileStartY := (tileNumber / tileXCount) * GM.TileSize

	return GM.tileset.SubImage(image.Rect(tileStartX, tileStartY, tileStartX+GM.TileSize, tileStartY+GM.TileSize)).(*ebiten.Image)
}

func NewGameMap(chunks [][]int, currentChunk int, roadsTo []map[string]int, sreenWidth int, sreenHeight int) (*GameMap, error) {
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
