package animatedobjects

import (
	"game/menu"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type WayToNextlevel struct {
	IsSpawned      bool
	x, y           int
	WNLPlayer      *goaseprite.Player
	WNLFile        *goaseprite.File
	WNLImage       *ebiten.Image
	tilecoordinate int
}

func (wnl *WayToNextlevel) SpawnWayNextLevel() {
	wnl.IsSpawned = true
}

func GoToNextLevel(vs *menu.VictoryScreen) {
	vs.InVictoryScreen = true
}
func (wnl *WayToNextlevel) GetCoordinates() (int, int) {
	var x int = ((wnl.tilecoordinate % 16) * 16)
	var y int = ((wnl.tilecoordinate / 16) * 16)

	return x, y
}

func (wnl *WayToNextlevel) InActiveZone(x, y int) bool {
	var centerChestX int = ((wnl.tilecoordinate % 16) * 16) + 16
	var centerChestY int = ((wnl.tilecoordinate / 16) * 16) + 8

	var objectCenterX int = x + 8
	var objectCenterY int = y + 8

	var distance float64 = math.Pow(float64(centerChestX-objectCenterX), 2) + math.Pow(float64(centerChestY-objectCenterY), 2)

	if distance <= 625 {
		return true
	}

	return false
}

func InitNewWayToNextLevel(jsonPath string, tilecoordinate int) (*WayToNextlevel, error) {
	wnl := &WayToNextlevel{
		IsSpawned:      false,
		WNLFile:        goaseprite.Open(jsonPath),
		tilecoordinate: tilecoordinate,
	}

	wnl.WNLPlayer = wnl.WNLFile.CreatePlayer()
	img, _, err := ebitenutil.NewImageFromFile(wnl.WNLFile.ImagePath)
	if err != nil {
		return nil, err
	}
	wnl.WNLImage = img

	wnl.WNLPlayer.Play("wait")

	return wnl, nil
}
