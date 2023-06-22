package ui

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type UI struct {
	HpBar *HealthBar
	WpBar *WeaponBar
	MiniM *MiniMap
	LB    *LevelBar
}

func DecodeImage(imagePath string) (*ebiten.Image, error) {
	imageFile, err := os.Open(imagePath)

	if err != nil {
		return nil, err
	}

	imageFileDecoded, _, err := image.Decode(imageFile)

	if err != nil {
		return nil, err
	}

	Image := ebiten.NewImageFromImage(imageFileDecoded)
	return Image, err
}

func InitUI() (*UI, error) {
	hpBar, err := InitHealthBar("./assets/healthpoint.png")
	if err != nil {
		return nil, err
	}
	wpBar, err := InitWeaponBar("./assets/startWeapon.png")
	if err != nil {
		return nil, err
	}
	miniM, err := InitMiniMap("./gamemap/assets/common.png", "./gamemap/assets/shop.png", "./gamemap/assets/chest.png", "./gamemap/assets/boss.png", "./gamemap/assets/current.png")
	if err != nil {
		return nil, err
	}

	lb, err := InitLevelBar()

	useri := &UI{
		HpBar: hpBar,
		WpBar: wpBar,
		MiniM: miniM,
		LB:    lb,
	}
	return useri, nil
}
