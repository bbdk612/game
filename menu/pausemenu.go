package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"os"
)

type PauseMenu struct {
	InPauseMenu       bool
	ContinuebuttonX   int
	ContinuebuttonY   int
	ContinuebuttonImg *ebiten.Image
	ExitToMMbuttonX   int
	ExitToMMbuttonY   int
	ExitToMMbuttonImg *ebiten.Image
}

func InitPauseMenu(continuebuttonImagePath, exitToMMbuttonImagePath string) (*PauseMenu, error) {
	continuebuttonFile, err := os.Open(continuebuttonImagePath)

	if err != nil {
		return nil, err
	}

	continuebuttonFileDecoded, _, err := image.Decode(continuebuttonFile)

	if err != nil {
		return nil, err
	}

	continuebuttonImage := ebiten.NewImageFromImage(continuebuttonFileDecoded)

	exitToMMbuttonFile, err := os.Open(exitToMMbuttonImagePath)

	if err != nil {
		return nil, err
	}

	exitToMMbuttonFileDecoded, _, err := image.Decode(exitToMMbuttonFile)

	if err != nil {
		return nil, err
	}

	exitToMMbuttonImage := ebiten.NewImageFromImage(exitToMMbuttonFileDecoded)

	pauseM := &PauseMenu{
		InPauseMenu:       false,
		ContinuebuttonX:   10,
		ContinuebuttonY:   50,
		ContinuebuttonImg: continuebuttonImage,
		ExitToMMbuttonX:   25,
		ExitToMMbuttonY:   75,
		ExitToMMbuttonImg: exitToMMbuttonImage,
	}

	return pauseM, nil
}

func (pm *PauseMenu) PauseMenuContinueGame() {
	pm.InPauseMenu = false
}

func (pm *PauseMenu) PauseMenuExitToMMGame(mm *MainMenu) {
	mm.InMainMenu = true
	pm.InPauseMenu = false
}
func (pm *PauseMenu) GetPauseMStartCoordinate() (int, int, int, int) {
	stbX := pm.ContinuebuttonX
	stbY := pm.ContinuebuttonY
	extX := pm.ExitToMMbuttonX
	extY := pm.ExitToMMbuttonY
	return stbX, stbY, extX, extY
}
