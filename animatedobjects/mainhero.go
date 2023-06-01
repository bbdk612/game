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


// TODO: Fix a moving on level
func (mh *MainHero) CanIGo(direction string, chunk []int) (bool) {
	switch direction {
		case "left":
			if mh.x % 16 != 0 {
				if mh.y % 16 != 0 {
					if (chunk[mh.tilecoordinate] == 1) && (chunk[mh.tilecoordinate + 16] == 1) {
						return true
					}
				}
				if chunk[mh.tilecoordinate] == 1 {
					return true
				}
			} else if mh.y % 16 != 0 {
				if (chunk[mh.tilecoordinate - 1] == 1) && (chunk[mh.tilecoordinate - 1 + 16] == 1) {
					return true
				}
			} else if chunk[mh.tilecoordinate - 1] == 1 {
				return true
			}


			return false

		case "right":
			if mh.y % 16 != 0 {
				fmt.Println((chunk[mh.tilecoordinate + 1]), (chunk[mh.tilecoordinate + 1 + 16]))
				if (chunk[mh.tilecoordinate + 1] == 1) && (chunk[mh.tilecoordinate + 1 + 16] == 1) {
					return true
				}
			} else if chunk[mh.tilecoordinate + 1] == 1 {
				return true
			}

			return false

		case "top":
			if mh.y % 16 != 0 {
				if mh.x % 16 != 0 {
					if (chunk[mh.tilecoordinate] == 1) && (chunk[mh.tilecoordinate + 1] == 1) {
						return true
					}
				}

				if chunk[mh.tilecoordinate] == 1 {
					return true
				}
			} else if mh.x % 16 != 0 {
				if (chunk[mh.tilecoordinate - 16] == 1) && (chunk[mh.tilecoordinate - 16 + 1] == 1) {
					return true
				}
			} else if chunk[mh.tilecoordinate - 16] == 1 {
				return true
			}

			return false

		case "down":
			if mh.x % 16 != 0 {
				if (chunk[mh.tilecoordinate + 16] == 1) && (chunk[mh.tilecoordinate + 1 + 16] == 1){
						return true
					}
			} else if chunk[mh.tilecoordinate + 16] == 1 {
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
