package animatedobjects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type Bullet struct {
	a, b                   float64
	x, y                   int
	XDirection, YDirection int
	Sprite                 *goaseprite.File
	AsePlayer              *goaseprite.Player
	Image                  *ebiten.Image
	step                   int
}

func (b *Bullet) Move(step int) {
	if b.XDirection != b.x {
		b.x += step
		b.y = int(float64(b.x)*b.a + b.b)
	} else {
		b.y += step
	}

}

func (b *Bullet) GetCoordinates() (int, int) {
	return b.x, b.y
}

func (b *Bullet) GetCurrentTile(tilesize int) int {
	var tile int = (((b.x) / tilesize) + (b.y/16)*tilesize) % 256
	return tile
}

func InitNewBullet(directionX, directionY int, a, b float64, startWeaponPositonX, startWeaponPositonY int, spriteJSONPath string, tilesize int) (*Bullet, error) {
	var step int = 2
	if (directionY < startWeaponPositonX) || (directionY < startWeaponPositonY) {
		step = -step
	}
	bullet := &Bullet{
		a:          a,
		b:          b,
		x:          startWeaponPositonX,
		y:          startWeaponPositonY,
		Sprite:     goaseprite.Open(spriteJSONPath),
		XDirection: directionX,
		YDirection: directionY,
		step:       step,
	}

	bullet.AsePlayer = bullet.Sprite.CreatePlayer()

	img, _, err := ebitenutil.NewImageFromFile(bullet.Sprite.ImagePath)
	if err != nil {
		return nil, err
	}

	bullet.Image = img
	return bullet, nil
}
