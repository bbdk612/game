package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type VictoryScreen struct {
	InVictoryScreen bool

	GoToNextLevelbuttonX      int
	GoToNextLevelbuttonY      int
	GoToNextLevelButtonImg    *ebiten.Image
	GoToNextLevelButtonFile   *goaseprite.File
	GoToNextLevelButtonPlayer *goaseprite.Player
}

func InitVictoryScreen(goToNextLevelbuttonJSONPath string) (*VictoryScreen, error) {

	VS := &VictoryScreen{
		InVictoryScreen:         false,
		GoToNextLevelbuttonX:    48,
		GoToNextLevelbuttonY:    75,
		GoToNextLevelButtonFile: goaseprite.Open(goToNextLevelbuttonJSONPath),
	}

	VS.GoToNextLevelButtonPlayer = VS.GoToNextLevelButtonFile.CreatePlayer()

	GoToNextLevelImg, _, err := ebitenutil.NewImageFromFile(VS.GoToNextLevelButtonFile.ImagePath)

	if err != nil {
		return nil, err
	}

	VS.GoToNextLevelButtonImg = GoToNextLevelImg

	return VS, nil
}

func (vs *VictoryScreen) VictoryScreenGoToNextLevel() {
	vs.InVictoryScreen = false
}
func (vs *VictoryScreen) GetVictoryScreenStartCoordinate() (int, int) {
	nxtX := vs.GoToNextLevelbuttonX
	nxtY := vs.GoToNextLevelbuttonY
	return nxtX, nxtY
}

func (vs *VictoryScreen) GoToNextLevelIsActive(cursorX, cursorY int) bool {
	if cursorX > vs.GoToNextLevelbuttonX+36 && cursorY > vs.GoToNextLevelbuttonY+15 {
		if cursorX < vs.GoToNextLevelbuttonY+36+80 && cursorY < vs.GoToNextLevelbuttonY+15+16 {
			vs.GoToNextLevelButtonPlayer.Play("Active")
			return true
		}
	}

	vs.GoToNextLevelButtonPlayer.Play("NoActive")
	return false
}
