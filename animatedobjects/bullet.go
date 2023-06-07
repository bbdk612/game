package github.com/bbdk612/game/animatedobjects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type Bullet struct {
	a, b                   float64
	x, y                   int
	XDirection, YDirection int
	endTile                int
	Sprite                 *goaseprite.File
	AsePlayer              *goaseprite.Player
	Image                  *ebiten.Image
}

func (b *Bullet) Move(step int) {
	if b.XDirection != b.x {
		if b.XDirection < b.x {
			b.x -= step
		} else {
			b.x += step
		}

		b.y = int(float64(b.x)*b.a + b.b)
	} else {
		if b.YDirection > b.y {
			b.y += step
		} else {
			b.y -= step
		}
	}

}

func (b *Bullet) GetCoordinates() (int, int) {
	return b.x, b.y
}

func (b *Bullet) GetCurrentTile(tilesize int) int {
	var tile int = ((b.x) / tilesize) + (b.y/16)*tilesize
	return tile
}

func (b *Bullet) GetEndTile() int {
	return b.endTile
}

func InitNewBullet(directionX, directionY int, a, b float64, startWeaponPositonX, startWeaponPositonY int, spriteJSONPath string, tilesize int) (*Bullet, error) {

	var endTile int = (directionX)/tilesize + (directionY/16)*tilesize

	bullet := &Bullet{
		a:          a,
		b:          b,
		x:          startWeaponPositonX,
		y:          startWeaponPositonY,
		Sprite:     goaseprite.Open(spriteJSONPath),
		XDirection: directionX,
		YDirection: directionY,
		endTile:    endTile,
	}

	bullet.AsePlayer = bullet.Sprite.CreatePlayer()

	img, _, err := ebitenutil.NewImageFromFile(bullet.Sprite.ImagePath)
	if err != nil {
		return nil, err
	}

	bullet.Image = img
	return bullet, nil
}
