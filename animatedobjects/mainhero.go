package animatedobject

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type MainHero struct {
	tilenumber int
	Sprite *goaseprite.File
	AsePlayer *goaseprite.Player
	Image  *ebiten.Image
}

func InitMainHero(tilenumber int) (*MainHero, error) {
	mainhero := &MainHero{
		Sprite: goaseprite.Open("./assets/mainhero.json"),
	}

	mainhero.AsePlayer = mainhero.Sprite.CreatePlayer()

	img, _, err := ebitenutil.NewImageFromFile(mainhero.Sprite.ImagePath)
	if err != nil {
		return nil, err
	}

	mainhero.Image = img

	return mainhero, nil
}
