package ui

import (
	"fmt"
	"image"
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
	weaponBarFile, err := os.Open(imagePath)

	if err != nil {
		return nil, err
	}

	weaponBarFileDecoded, _, err := image.Decode(weaponBarFile)

	if err != nil {
		return nil, err
	}

	weaponBarImage := ebiten.NewImageFromImage(weaponBarFileDecoded)

	fontBytes, err := os.ReadFile("./assets/font.ttf")
	if err != nil {
		return nil, err
	}

	fontParsed, err := opentype.Parse(fontBytes)

	if err != nil {
		return nil, err
	}

	font, err := opentype.NewFace(fontParsed, &opentype.FaceOptions{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingVertical,
	})

	wpB := &WeaponBar{
		startX:   250,
		startY:   150,
		Image:    weaponBarImage,
		AmmoFont: font,
	}

	return wpB, nil
}
