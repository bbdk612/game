package ui

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type HealthBar struct {
	startX, startY int
	HealthNumber   int
	Image          *ebiten.Image
}

func InitHealthBar(imagePath string) (*HealthBar, error) {
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
		startX: 5,
		startY: 5,
		Image:  healthBarImage,
	}

	return hpB, nil
}

func (hpB *HealthBar) GetHpbStartCoordinate() (int, int) {
	stX := hpB.startX
	stY := hpB.startY
	return stX, stY
}
