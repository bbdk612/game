package gamemap

import{

	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
}

type GameMapOptions struct {
	TileSize                int
	SreenWidth, SreenHeight int
	tileset                 *ebiten.Image
}

func (GM *GameMapOptions) GetTile(tileNumber int) *ebiten.Image {
	w := GM.tileset.Bounds().Dx()
	tileXCount := w / GM.TileSize

	tileStartX := (tileNumber % tileXCount) * GM.TileSize
	tileStartY := (tileNumber / tileXCount) * GM.TileSize

	return GM.tileset.SubImage(image.Rect(tileStartX, tileStartY, tileStartX+GM.TileSize, tileStartY+GM.TileSize)).(*ebiten.Image)
}

func InitGameMap(tilesetImgPath string) (*GameMapOptions, error) {
	tilesetFile, err := os.Open(tilesetImgPath)
	if err != nil {
		return nil, err
	}

	tileset, _, err := image.Decode(tilesetFile)

	if err != nil {
		return nil, err
	}

	tilesImage := ebiten.NewImageFromImage(tileset)

	GM := &GameMapOptions{
		TileSize:     16,
		SreenWidth:   sreenWidth,
		SreenHeight:  sreenHeight,
		tileset:      tilesImage,
	}
	return GM, nil
}
