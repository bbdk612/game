package ui

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type MiniMap struct {
	startX, startY  int
	CommonRoomImage *ebiten.Image
	ShopRoomImage   *ebiten.Image
	ChestRoomImage  *ebiten.Image
	BossRoomImage   *ebiten.Image
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

func InitMiniMap(CommonRoomimagePath, ShopRoomimagePath, ChestRoomimagePath, BossRoomimagePath string) (*MiniMap, error) {
	//common room image
	commonRoomImage, err := DecodeImage(CommonRoomimagePath)

	if err != nil {
		return nil, err
	}
	//shop room image
	shopRoomImage, err := DecodeImage(ShopRoomimagePath)

	if err != nil {
		return nil, err
	}
	//chest room image
	chestRoomImage, err := DecodeImage(ChestRoomimagePath)

	if err != nil {
		return nil, err
	}
	//boss room image
	bossRoomImage, err := DecodeImage(BossRoomimagePath)

	if err != nil {
		return nil, err
	}

	mm := &MiniMap{
		startX:          10,
		startY:          240,
		CommonRoomImage: commonRoomImage,
		ShopRoomImage:   shopRoomImage,
		ChestRoomImage:  chestRoomImage,
		BossRoomImage:   bossRoomImage,
	}

	return mm, nil
}

func (mm *MiniMap) GetMiniMapStartCoordinate() (int, int) {
	stX := mm.startX
	stY := mm.startY
	return stX, stY
}
