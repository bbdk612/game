package ui

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/bbdk612/game/ui"
)

type HealthBar struct {
	startX, startY                   int
	HealthNumber           int
	Image                  *ebiten.Image
}
func InitHealthBar (imagePath string) (*HealthBar, error){
	healthBarFile, err := os.Open(imagePath)

	if err != nil {
		return nil, err
	}

	healthBarFileDecoded, _, err := image.Decode(healthBarFile)

	if err != nil {
		return nil, err
	}

	healthBarImage := ebiten.NewImageFromImage(healthBarFileDecoded)

	hpB := &HealthBar{
		oX:    1,
		oY:    1,
		HealthNumber: 6,
		Image: healthBarImage,
	}

	return hpB, nil
}

func (hpB *HealthBar) Damage(x, y, charX, charY  int){
	if x=charX && y= charY{
		hpB.HealthNumber = hpB.HealthNumber - 1;
	}
}
