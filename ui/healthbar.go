package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type HealthBar struct {
	startX, startY int
	HealthNumber   int
	Image          *ebiten.Image
}

func InitHealthBar(imagePath string) (*HealthBar, error) {
	//health bar image
	healthBarImage, err := DecodeImage(imagePath)

	if err != nil {
		return nil, err
	}

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
