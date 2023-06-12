package animatedobjects

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type Bullet struct {
	a, b           float64
	x, y           float64
	deltaX, deltaY float64
	Sprite         *goaseprite.File
	AsePlayer      *goaseprite.Player
	Image          *ebiten.Image
	step           float64
}

func (b *Bullet) Move() {
	if math.Abs(b.a) == math.Inf(1) {
		b.y += b.step
	} else if b.a == 0 {
		b.x += b.step
	} else {
		b.x += b.step
		b.y = (b.x*b.a + b.b)
	}
}

func (b *Bullet) GetCoordinates() (float64, float64) {
	return b.x, b.y
}

func (b *Bullet) GetCurrentTile(tilesize int) int {
	var tile int = ((int(b.x) / tilesize) + ((int(b.y) / 16) * tilesize)) % 256
	if tile < 0 {
		tile = 0
	}
	return tile
}

func InitNewBullet(directionX, directionY float64, a, b float64, step float64, startWeaponPositonX, startWeaponPositonY float64, spriteJSONPath string, tilesize int) (*Bullet, error) {
	fmt.Println("step:", step, "Start:", startWeaponPositonX, startWeaponPositonY)
	fmt.Println("direction:", directionX, directionY)
	bullet := &Bullet{
		a:      a,
		b:      b,
		x:      startWeaponPositonX,
		y:      startWeaponPositonY,
		Sprite: goaseprite.Open(spriteJSONPath),
		deltaX: directionX - startWeaponPositonX,
		deltaY: directionY - startWeaponPositonY,
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
