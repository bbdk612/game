package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type MiniMap struct {
	startX, startY   int
	CommonRoomImage  *ebiten.Image
	ShopRoomImage    *ebiten.Image
	ChestRoomImage   *ebiten.Image
	BossRoomImage    *ebiten.Image
	CurrentRoomImage *ebiten.Image
}

func InitMiniMap(CommonRoomimagePath, ShopRoomimagePath, ChestRoomimagePath, BossRoomimagePath, CurrentRoomimagePath string) (*MiniMap, error) {
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
	//current room image
	currentRoomImage, err := DecodeImage(CurrentRoomimagePath)

	if err != nil {
		return nil, err
	}

	mm := &MiniMap{
		startX:           10,
		startY:           240,
		CommonRoomImage:  commonRoomImage,
		ShopRoomImage:    shopRoomImage,
		ChestRoomImage:   chestRoomImage,
		BossRoomImage:    bossRoomImage,
		CurrentRoomImage: currentRoomImage,
	}

	return mm, nil
}

func (mm *MiniMap) GetMiniMapStartCoordinate() (int, int) {
	stX := mm.startX
	stY := mm.startY
	return stX, stY
}
