package menu

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type MainMenu struct {
	InMainMenu        bool
	StartbuttonX      int
	StartbuttonY      int
	StartButtonFile   *goaseprite.File
	StartButtonPlayer *goaseprite.Player
	StartButtonImg    *ebiten.Image
	ExitbuttonX       int
	ExitbuttonY       int
	ExitButtonFile    *goaseprite.File
	ExitButtonPlayer  *goaseprite.Player
	ExitButtonImg     *ebiten.Image
	MainImage         *ebiten.Image
}

func InitMenu(startbuttonJSONPath, exitbuttonJSONPath string) (*MainMenu, error) {

	mainM := &MainMenu{
		InMainMenu:      true,
		StartbuttonX:    10,
		StartbuttonY:    50,
		StartButtonFile: goaseprite.Open(startbuttonJSONPath),
		ExitbuttonX:     10,
		ExitbuttonY:     75,
		ExitButtonFile:  goaseprite.Open(exitbuttonJSONPath),
	}

	mainM.StartButtonPlayer = mainM.StartButtonFile.CreatePlayer()

	startimg, _, err := ebitenutil.NewImageFromFile(mainM.StartButtonFile.ImagePath)

	if err != nil {
		return nil, err
	}

	mainM.StartButtonImg = startimg

	mainM.ExitButtonPlayer = mainM.ExitButtonFile.CreatePlayer()

	exitimg, _, err := ebitenutil.NewImageFromFile(mainM.ExitButtonFile.ImagePath)

	if err != nil {
		return nil, err
	}

	mainM.ExitButtonImg = exitimg

	imagefile, err := os.Open("./assets/back.png")
	if err != nil {
		return nil, err
	}

	imagedecoded, _, err := image.Decode(imagefile)
	if err != nil {
		return nil, err
	}

	mainM.MainImage = ebiten.NewImageFromImage(imagedecoded)

	return mainM, nil
}

func (mm *MainMenu) MenuStartGame() {
	mm.InMainMenu = false
}

func (mm *MainMenu) MenuExitGame() {
	os.Exit(0)
}
func (mm *MainMenu) GetMainMStartCoordinate() (int, int, int, int) {
	stbX := mm.StartbuttonX
	stbY := mm.StartbuttonY
	extX := mm.ExitbuttonX
	extY := mm.ExitbuttonY
	return stbX, stbY, extX, extY
}

func (mm *MainMenu) StartIsActive(cursorX, cursorY int) bool {
	if cursorX > mm.StartbuttonX && cursorX < mm.StartbuttonX+48 {
		if cursorY > mm.StartbuttonY && cursorY < mm.StartbuttonY+16 {
			mm.StartButtonPlayer.Play("Active")
			return true
		}
	}
	mm.StartButtonPlayer.Play("NoActive")
	return false
}
func (mm *MainMenu) ExitIsActive(cursorX, cursorY int) bool {
	if cursorX > mm.ExitbuttonX && cursorX < mm.ExitbuttonX+48 {
		if cursorY > mm.ExitbuttonY && cursorY < mm.ExitbuttonY+16 {
			mm.ExitButtonPlayer.Play("Active")
			return true
		}
	}
	mm.ExitButtonPlayer.Play("NoActive")
	return false
}
