package ui

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type WeaponBar struct {
	startX, startY int
	Image          *ebiten.Image
	AmmoFont       font.Face
}

func (wb *WeaponBar) GetAmmo(currAmmo, maxAmmo int) string {
	return fmt.Sprintf("Ammo %d/%d", currAmmo, maxAmmo)
}

func (wb *WeaponBar) GetWpbStartCoordinate() (int, int) {
	return wb.startX, wb.startY
}

func InitWeaponBar(imagePath string) (*WeaponBar, error) {
	//health bar image
	weaponBarImage, err := DecodeImage(imagePath)

	if err != nil {
		return nil, err
	}

	fontBytes, err := os.ReadFile("./assets/font.ttf")
	if err != nil {
		return nil, err
	}

	fontParsed, err := opentype.Parse(fontBytes)

	if err != nil {
		return nil, err
	}

	font, err := opentype.NewFace(fontParsed, &opentype.FaceOptions{
		Size:    10,
		DPI:     150,
		Hinting: font.HintingVertical,
	})

	wpB := &WeaponBar{
		startX:   176,
		startY:   250,
		Image:    weaponBarImage,
		AmmoFont: font,
	}

	return wpB, nil
}
