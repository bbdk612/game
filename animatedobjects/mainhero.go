package animatedobject

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type MainHero struct {
	tilecoordinate int
	x, y int
	Sprite *goaseprite.File
	AsePlayer *goaseprite.Player
	Image  *ebiten.Image
}

func (mh *MainHero) calculateTilenumber(tilesize int) {
	mh.tilecoordinate = (mh.x / tilesize) + mh.y
}



func InitMainHero(tilecoordinate int, tilesize int, xCount int) (*MainHero, error) {
	var x int = (tilecoordinate % xCount) / tilesize
	var y int = (tilecoordinate / xCount) / tilesize
	mainhero := &MainHero{
		Sprite: goaseprite.Open("./assets/mainhero.json"),
		tilecoordinate: tilecoordinate,
		x: x,
		y: y,
	}

	mainhero.AsePlayer = mainhero.Sprite.CreatePlayer()

	img, _, err := ebitenutil.NewImageFromFile(mainhero.Sprite.ImagePath)
	if err != nil {
		return nil, err
	}

	mainhero.Image = img

	return mainhero, nil
}
