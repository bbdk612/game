package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type DeathScreen struct {
	InDeathScreen bool

	ReturnToMMbuttonX      int
	ReturnToMMbuttonY      int
	ReturnToMMButtonImg    *ebiten.Image
	ReturnToMMButtonFile   *goaseprite.File
	ReturnToMMButtonPlayer *goaseprite.Player
}

func InitDeathScreen(returnToMMbuttonJSONPath string) (*DeathScreen, error) {

	DS := &DeathScreen{
		InDeathScreen:        false,
		ReturnToMMbuttonX:    10,
		ReturnToMMbuttonY:    75,
		ReturnToMMButtonFile: goaseprite.Open(returnToMMbuttonJSONPath),
	}

	DS.ReturnToMMButtonPlayer = DS.ReturnToMMButtonFile.CreatePlayer()

	ExitToMMImg, _, err := ebitenutil.NewImageFromFile(DS.ReturnToMMButtonFile.ImagePath)

	if err != nil {
		return nil, err
	}

	DS.ReturnToMMButtonImg = ExitToMMImg

	return DS, nil
}

func (ds *DeathScreen) DeathScreenReturnToMMGame(mm *MainMenu) {
	mm.InMainMenu = true
	ds.InDeathScreen = false
}
func (ds *DeathScreen) GetDathScreenStartCoordinate() (int, int) {
	extX := ds.ReturnToMMbuttonX
	extY := ds.ReturnToMMbuttonY
	return extX, extY
}

func (ds *DeathScreen) ReturnToMMIsActive(cursorX, cursorY int) bool {
	if cursorX > ds.ReturnToMMbuttonX && cursorY > ds.ReturnToMMbuttonY {
		if cursorX < ds.ReturnToMMbuttonY+80 && cursorY < ds.ReturnToMMbuttonY+16 {
			ds.ReturnToMMButtonPlayer.Play("Active")
			return true
		}
	}

	ds.ReturnToMMButtonPlayer.Play("NoActive")
	return false
}
