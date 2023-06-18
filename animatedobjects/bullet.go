package animatedobjects

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type Bullet struct {
	a, b      float64
	X, Y      float64
	Sprite    *goaseprite.File
	AsePlayer *goaseprite.Player
	Image     *ebiten.Image
	step      float64

	CalculateNextStep func(float64, float64, float64, float64, float64) (float64, float64)
}

func (b *Bullet) GetCoordinates() (float64, float64) {
	return b.X, b.Y
}

func (b *Bullet) GetCurrentTile(tilesize int) int {
	var tile int = ((int(b.X) / tilesize) + ((int(b.Y) / 16) * tilesize)) % 256
	if tile < 0 {
		tile = 0
	}
	return tile
}

func (b *Bullet) Move() {
	b.X, b.Y = b.CalculateNextStep(b.X, b.Y, b.a, b.b, b.step)
}

func InitNewBullet(a, b float64, step float64, startWeaponPositonX, startWeaponPositonY float64, spriteJSONPath string, tilesize int) (*Bullet, error) {
	bullet := &Bullet{
		a:      a,
		b:      b,
		X:      startWeaponPositonX,
		Y:      startWeaponPositonY,
		Sprite: goaseprite.Open(spriteJSONPath),
		step:   step,
	}

	bullet.AsePlayer = bullet.Sprite.CreatePlayer()

	img, _, err := ebitenutil.NewImageFromFile(bullet.Sprite.ImagePath)
	if err != nil {
		return nil, err
	}

	bullet.Image = img
	return bullet, nil
}
