package weapons

import (
	// "fmt"
	"fmt"
	ao "game/animatedobjects"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"math"
	"os"
	"time"
)

type Weapon struct {
	oX, oY       int
	angle        float64
	xEnd, yEnd   int
	Image        *ebiten.Image
	rollbackTime time.Time
	Bullets      [](*ao.Bullet)
	currentAmmo  int
	maxAmmo      int
}

func (w *Weapon) CalculateAngle(x, y int) {
	var aY float64 = (float64(w.oY) - float64(y))
	var aX float64 = (float64(w.oX) - float64(x))
	var bY float64 = (float64(w.oY) - float64(w.yEnd))
	var bX float64 = (float64(w.oX) - float64(w.xEnd))

	w.angle = (-math.Atan2(bY, bX) + math.Atan2(aY, aX))
}

func (w *Weapon) GetAmmo() (int, int) {
	return w.currentAmmo, w.maxAmmo
}

func (w *Weapon) GetAngle() float64 {
	return w.angle
}

func (w *Weapon) GetOCoordinates() (int, int) {
	return w.oX, w.oY
}

func (w *Weapon) ChangePosition(x, y int) {
	w.oX = x
	w.oY = y

	w.xEnd = x + 8
	w.yEnd = y + 8
}

func (w *Weapon) MoveWeapon(direction string, step int) {

	switch direction {
	case "left":
		w.ChangePosition(w.oX-step, w.oY)

	case "right":
		w.ChangePosition(w.oX+step, w.oY)

	case "top":
		w.ChangePosition(w.oX, w.oY-step)

	case "down":
		w.ChangePosition(w.oX, w.oY+step)
	}
}

func (w *Weapon) Shoot(directionX, directionY int, spritePath string, tilesize int) (*ao.Bullet, error) {
	if w.currentAmmo != 0 {
		rlbkDur, err := time.ParseDuration("500ms")
		if err != nil {
			return nil, err
		}
		rollbk := rlbkDur.Milliseconds()

		currTime := time.Now()
		if currTime.Sub(w.rollbackTime).Milliseconds() >= rollbk {

			var deltaX float64 = float64(w.oX) - float64(directionX)
			var deltaY float64 = float64(w.oY) - float64(directionY)
			var startY float64 = float64(w.oY)
			var startX float64 = float64(w.oX)
			var a, b float64
			var step float64 = 2
			if deltaY != 0 {
				if deltaX != 0 {
					a = deltaY / deltaX
					b = float64(w.oY) - (float64(w.oX) * a)

					if deltaX > 0 {
						step = -2
						startX = startX - 8.0
					} else {
						startX = startX + 8.0
					}
					fmt.Println(step)
					fmt.Println()
					fmt.Println(startX + (-4))
					startY = (startX*a + b)
				} else {
					a = deltaY / deltaX
					if a > 0 {
						step = -2
						startY -= 8
					} else {
						startY += 8
					}

				}
			} else {
				if deltaX > 0 {
					step = -2
					startX -= 8
				} else {
					startX += 8
				}

			}

			bullet, err := ao.InitNewBullet(float64(directionX), float64(directionY), a, b, step, startX, startY, spritePath, 16)

			if err != nil {
				return nil, err
			}

			w.currentAmmo--
			w.rollbackTime = time.Now()
			return bullet, nil
		}
	}
	return nil, nil
}

func (w *Weapon) Reload() {
	w.currentAmmo = w.maxAmmo
}

func InitNewWeapon(x, y int, imagePath string) (*Weapon, error) {
	weaponFile, err := os.Open(imagePath)

	if err != nil {
		return nil, err
	}

	weaponFileDecoded, _, err := image.Decode(weaponFile)

	if err != nil {
		return nil, err
	}

	weaponImage := ebiten.NewImageFromImage(weaponFileDecoded)

	w := &Weapon{
		oX:          x,
		oY:          y,
		xEnd:        x + 8,
		yEnd:        y + 8,
		Image:       weaponImage,
		angle:       0.0,
		currentAmmo: 20,
		maxAmmo:     20,
	}

	return w, nil
}
