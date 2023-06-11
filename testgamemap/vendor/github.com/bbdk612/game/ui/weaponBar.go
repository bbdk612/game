package ui

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type WeaponBar struct {
	startX, startY int
	Image          *ebiten.Image
}

func InitWeaponBar(imagePath string) (*WeaponBar, error) {
	weaponBarFile, err := os.Open(imagePath)

	if err != nil {
		return nil, err
	}

	weaponBarFileDecoded, _, err := image.Decode(weaponBarFile)

	if err != nil {
		return nil, err
	}

	weaponBarImage := ebiten.NewImageFromImage(weaponBarFileDecoded)

	wpB := &WeaponBar{
		oX:    1,
		oY:    1,
		Image: healthBarImage,
	}

	return wpB, nil
}