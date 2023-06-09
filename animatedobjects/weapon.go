package animatedobjects

import (
	"image"
	_ "image/png"
	"math"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Weapon struct {
	oX, oY       int
	angle        float64
	xEnd, yEnd   int
	Image        *ebiten.Image
	rollbackTime time.Time
	Bullets      [](*Bullet)
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

func (w *Weapon) Shoot(directionX, directionY int, spritePath string, tilesize int) error {
	if w.currentAmmo != 0 {
		rlbkDur, err := time.ParseDuration("500ms")
		if err != nil {
			return err
		}
		rollbk := rlbkDur.Milliseconds()

		currTime := time.Now()
		if currTime.Sub(w.rollbackTime).Milliseconds() >= rollbk {

			var deltaX float64 = float64(directionX - w.oX)
			var deltaY float64 = float64(directionY - w.oY)

			var a float64 = deltaY / deltaX
			var b float64 = float64(w.oY) - a*float64(w.oX)
			var startX int = w.oX
			var startY int = w.oY
			if deltaX < 0 {
				startX -= 8
				startY = int(float64(startX)*a + b)
			} else if deltaX > 0 {
				startX += 8
				startY = int(float64(startX)*a + b)
			} else if deltaX == 0 {
				if deltaY > 0 {
					startY += 8
				} else {
					startY += 8
				}
			}

			bullet, err := InitNewBullet(directionX, directionY, a, b, startX, startY, spritePath, 16)

			if err != nil {
				return err
			}

			w.Bullets = append(w.Bullets, bullet)
			w.currentAmmo--
			w.rollbackTime = time.Now()
		}
	}
	return nil
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
