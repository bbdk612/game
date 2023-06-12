package animatedobjects

import (
	// "fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type MainHero struct {
	tilecoordinate int
	x, y           int
	Sprite         *goaseprite.File
	AsePlayer      *goaseprite.Player
	Health         int
	Image          *ebiten.Image
	Weapons        [2](*Weapon)
	step           int
	currentWeapon  int
}

func (mh *MainHero) calculateTilecoordinate(tilesize int) {
	//calculate tilecoordiante on screen
	mh.tilecoordinate = ((mh.x) / tilesize) + (mh.y/16)*16
}

func (mh *MainHero) GetTileCoor() int {
	return mh.tilecoordinate
}

func (mh *MainHero) SetTileCoor(tilecoor int) {
	var x int = (tilecoor % 16) * 16
	var y int = (tilecoor / 16) * 16

	mh.x = x
	mh.y = y
	mh.tilecoordinate = tilecoor
}

func (mh *MainHero) Damage() int {
	mh.Health--
	return mh.Health
}

func (mh *MainHero) Heal(heal int) int {
	mh.Health += heal
	return mh.Health
}

func (mh *MainHero) GetCoordinates() (int, int) {
	return mh.x, mh.y
}

func (mh *MainHero) SetCoordinates(x, y int) {
	mh.x = x
	mh.y = y

	mh.GetCurrentWeapon().ChangePosition(x+8, y+8)

	mh.calculateTilecoordinate(16)
}

func (mh *MainHero) CanIGo(direction string, chunk []int) bool {
	switch direction {
	case "left":
		if mh.x%16 != 0 {
			if mh.y%16 != 0 {
				if (chunk[mh.tilecoordinate] == 1) && (chunk[mh.tilecoordinate+16] == 1) {
					return true
				}
			}
			if chunk[mh.tilecoordinate] == 1 {
				return true
			}
		} else if mh.y%16 != 0 {
			if (chunk[mh.tilecoordinate-1] == 1) && (chunk[mh.tilecoordinate-1+16] == 1) {
				return true
			}
		} else if chunk[mh.tilecoordinate-1] == 1 {
			return true
		}

		return false

	case "right":
		if mh.y%16 != 0 {
			if (chunk[mh.tilecoordinate+1] == 1) && (chunk[mh.tilecoordinate+1+16] == 1) {
				return true
			}
		} else if chunk[mh.tilecoordinate+1] == 1 {
			return true
		}

		return false

	case "top":
		if mh.y%16 != 0 {
			if mh.x%16 != 0 {
				if (chunk[mh.tilecoordinate] == 1) && (chunk[mh.tilecoordinate+1] == 1) {
					return true
				}
			}

			if chunk[mh.tilecoordinate] == 1 {
				return true
			}
		} else if mh.x%16 != 0 {
			if (chunk[mh.tilecoordinate-16] == 1) && (chunk[mh.tilecoordinate-16+1] == 1) {
				return true
			}
		} else if chunk[mh.tilecoordinate-16] == 1 {
			return true
		}

		return false

	case "down":
		if mh.x%16 != 0 {
			if (chunk[mh.tilecoordinate+16] == 1) && (chunk[mh.tilecoordinate+1+16] == 1) {
				return true
			}
		} else if chunk[mh.tilecoordinate+16] == 1 {
			return true
		}

		return false

	default:
		return false
	}
}

func (mh *MainHero) GetCurrentWeapon() *Weapon {
	return mh.Weapons[mh.currentWeapon]
}

func (mh *MainHero) Move(direction string, chunk []int) {
	if mh.CanIGo(direction, chunk) {
		switch direction {
		case "left":
			mh.SetCoordinates(mh.x-mh.step, mh.y)

		case "right":
			mh.SetCoordinates(mh.x+mh.step, mh.y)

		case "top":
			mh.SetCoordinates(mh.x, mh.y-mh.step)

		case "down":
			mh.SetCoordinates(mh.x, mh.y+mh.step)
		}

	}
}

func InitMainHero(tilecoordinate int, tilesize int, xCount int, step int) (*MainHero, error) {
	var x int = (tilecoordinate % xCount) * tilesize
	var y int = (tilecoordinate / xCount) * tilesize

	startWeapon, err := InitNewWeapon(x+8, y+8, "./assets/startWeapon.png")
	var weapons [2](*Weapon)
	weapons[0] = startWeapon
	if err != nil {
		return nil, err
	}

	mainhero := &MainHero{
		Sprite:         goaseprite.Open("./assets/mainhero.json"),
		tilecoordinate: tilecoordinate,
		x:              x,
		y:              y,
		Weapons:        weapons,
		currentWeapon:  0,
		step:           step,
		Health:         6,
	}

	mainhero.AsePlayer = mainhero.Sprite.CreatePlayer()

	img, _, err := ebitenutil.NewImageFromFile(mainhero.Sprite.ImagePath)
	if err != nil {
		return nil, err
	}

	mainhero.Image = img

	return mainhero, nil
}
