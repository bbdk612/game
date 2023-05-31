package animatedobjects

import (
	"fmt"

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

func (mh *MainHero) calculateTilecoordinate(tilesize int) {
	//calculate tilecoordiante on screen
	mh.tilecoordinate = ((mh.x) / tilesize) + (mh.y / 16) * 16
	fmt.Println(mh.x, mh.y, mh.tilecoordinate)
}

func (mh *MainHero) GetCoordinates() (int, int) {
	return mh.x, mh.y
}

func (mh *MainHero) SetCoordinates(x, y int) {
	mh.x = x
	mh.y = y
	mh.calculateTilecoordinate(16)
}

func (mh *MainHero) iBetweenTiles() (bool) {
	if (mh.y % 16 != 0) || (mh.x % 16 != 0) {
		return true
	}

	return false
}
// TODO: Fix a moving on level
func (mh *MainHero) CanIGo(direction string, chunk []int) (bool) {
	switch direction {
		case "left":
			fmt.Println("nextTile", mh.tilecoordinate-1)
			if chunk[mh.tilecoordinate - 1] == 1 {
				return true
			}

			return false

		case "right":
			fmt.Println("nextTile", mh.tilecoordinate+1)
			if chunk[mh.tilecoordinate + 1] == 1 {
				return true
			}

			return false

		case "top":
			fmt.Println("nextTile", mh.tilecoordinate-16)
			if chunk[mh.tilecoordinate - 16] == 1 {
				return true
			}

			return false

		case "down":
			fmt.Println("nextTile", mh.tilecoordinate+16)
			if chunk[mh.tilecoordinate + 16] == 1 {
				return true
			}

			return false

		default:
			return false
	}
}

func InitMainHero(tilecoordinate int, tilesize int, xCount int) (*MainHero, error) {
	var x int = (tilecoordinate % xCount) * tilesize
	var y int = (tilecoordinate / xCount) * tilesize
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
