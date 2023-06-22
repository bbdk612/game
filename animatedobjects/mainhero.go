package animatedobjects

import (
	// "fmt"

	"game/weapons"
	"math"

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
	MaxHealth      int
	Image          *ebiten.Image
	Weapons        [3](*weapons.Weapon)
	step           int
	CurrentWeapon  int
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
	if mh.Health > mh.MaxHealth {
		mh.Health = mh.MaxHealth
	}
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

func (mh *MainHero) GetCurrentWeapon() *weapons.Weapon {
	return mh.Weapons[mh.CurrentWeapon]
}

func (mh *MainHero) Move(direction string, chunk []int, coords [][]float64) {
	if mh.CanIGo(direction, chunk) {
		switch direction {
		case "left":
			for _, i := range coords[1:] {
				distX := math.Abs(float64(mh.x-mh.step) - i[0])
				distY := math.Abs(float64(mh.y) - i[1])
				if distX < 20 && distY < 20 {
					return
				}
			}
			mh.SetCoordinates(mh.x-mh.step, mh.y)

		case "right":
			for _, i := range coords[1:] {
				distX := math.Abs(float64(mh.x+mh.step) - i[0])
				distY := math.Abs(float64(mh.y) - i[1])
				if distX < 20 && distY < 20 {
					return
				}
			}
			mh.SetCoordinates(mh.x+mh.step, mh.y)

		case "top":
			for _, i := range coords[1:] {
				distX := math.Abs(float64(mh.x) - i[0])
				distY := math.Abs(float64(mh.y-mh.step) - i[1])
				if distX < 20 && distY < 20 {
					return
				}
			}
			mh.SetCoordinates(mh.x, mh.y-mh.step)

		case "down":
			for _, i := range coords[1:] {
				distX := math.Abs(float64(mh.x) - i[0])
				distY := math.Abs(float64(mh.y+mh.step) - i[1])
				if distX < 20 && distY < 20 {
					return
				}
			}
			mh.SetCoordinates(mh.x, mh.y+mh.step)
		}

	}
}

func InitMainHero(tilecoordinate int, tilesize int, xCount int, step int) (*MainHero, error) {
	var x int = (tilecoordinate % xCount) * tilesize
	var y int = (tilecoordinate / xCount) * tilesize

	startWeapon, err := weapons.InitNewWeapon(x+8, y+8, "./weapons/assets/shotgun.json")
	var weapons [3](*weapons.Weapon)
	weapons[0] = startWeapon
	if err != nil {
		return nil, err
	}

	mainhero := &MainHero{
		Sprite:         goaseprite.Open("./assets/mainheroNew.json"),
		tilecoordinate: tilecoordinate,
		x:              x,
		y:              y,
		Weapons:        weapons,
		CurrentWeapon:  0,
		step:           step,
		MaxHealth:      6,
	}

	mainhero.AsePlayer = mainhero.Sprite.CreatePlayer()

	img, _, err := ebitenutil.NewImageFromFile(mainhero.Sprite.ImagePath)
	if err != nil {
		return nil, err
	}

	mainhero.Image = img

	return mainhero, nil
}
