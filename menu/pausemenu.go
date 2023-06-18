package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type PauseMenu struct {
	InPauseMenu          bool
	ContinuebuttonX      int
	ContinuebuttonY      int
	ContinueButtonFile   *goaseprite.File
	ContinueButtonPlayer *goaseprite.Player
	ContinueButtonImg    *ebiten.Image

	ExitToMMbuttonX      int
	ExitToMMbuttonY      int
	ExitToMMButtonImg    *ebiten.Image
	ExitToMMButtonFile   *goaseprite.File
	ExitToMMButtonPlayer *goaseprite.Player
}

func InitPauseMenu(continuebuttonJSONPath, exitToMMbuttonJSONPath string) (*PauseMenu, error) {

	pauseM := &PauseMenu{
		InPauseMenu:        false,
		ContinuebuttonX:    10,
		ContinuebuttonY:    50,
		ContinueButtonFile: goaseprite.Open(continuebuttonJSONPath),
		ExitToMMbuttonX:    10,
		ExitToMMbuttonY:    75,
		ExitToMMButtonFile: goaseprite.Open(exitToMMbuttonJSONPath),
	}

	pauseM.ContinueButtonPlayer = pauseM.ContinueButtonFile.CreatePlayer()

	ContinueImg, _, err := ebitenutil.NewImageFromFile(pauseM.ContinueButtonFile.ImagePath)

	if err != nil {
		return nil, err
	}

	pauseM.ContinueButtonImg = ContinueImg

	pauseM.ExitToMMButtonPlayer = pauseM.ExitToMMButtonFile.CreatePlayer()

	ExitToMMImg, _, err := ebitenutil.NewImageFromFile(pauseM.ExitToMMButtonFile.ImagePath)

	if err != nil {
		return nil, err
	}

	pauseM.ExitToMMButtonImg = ExitToMMImg

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

func (pm *PauseMenu) ContinueIsActive(cursorX, cursorY int) bool {
	if cursorX > pm.ContinuebuttonX && cursorY > pm.ContinuebuttonY {
		if cursorX < pm.ContinuebuttonX+80 && cursorY < pm.ContinuebuttonY+16 {
			pm.ContinueButtonPlayer.Play("Active")
			return true
		}
	}

	pm.ContinueButtonPlayer.Play("NoActive")
	return false
}

func (pm *PauseMenu) ExitToMMIsActive(cursorX, cursorY int) bool {
	if cursorX > pm.ExitToMMbuttonX && cursorY > pm.ExitToMMbuttonY {
		if cursorX < pm.ExitToMMbuttonY+80 && cursorY < pm.ExitToMMbuttonY+16 {
			pm.ExitToMMButtonPlayer.Play("Active")
			return true
		}
	}

	pm.ExitToMMButtonPlayer.Play("NoActive")
	return false
}
