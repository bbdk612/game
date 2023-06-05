package animatedobjects

import (
	"image"
	_ "image/png"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Weapon struct {
	oX, oY     int
	angle      float64
	xEnd, yEnd int
	Image      *ebiten.Image
}

func (w *Weapon) CalculateAngle(x, y int) {
	var aY float64 = (float64(w.oY) - float64(y))
	var aX float64 = (float64(w.oX) - float64(x))
	var bY float64 = (float64(w.oY) - float64(w.yEnd))
	var bX float64 = (float64(w.oX) - float64(w.xEnd))

	w.angle = (-math.Atan2(bY, bX) + math.Atan2(aY, aX)) + math.Pi
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
		w.ChangePosition(w.oX - step, w.oY)

	case "right":
		w.ChangePosition(w.oX+step, w.oY)

	case "top":
		w.ChangePosition(w.oX, w.oY-step)

	case "down":
		w.ChangePosition(w.oX, w.oY+step)
	}
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
		oX:    x,
		oY:    y,
		xEnd:  x + 8,
		yEnd:  y + 8,
		Image: weaponImage,
		angle: 0.0,
	}

	return w, nil
}
