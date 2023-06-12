package menu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"os"
)

type MainMenu struct {
	InMainMenu     bool
	StartbuttonX   int
	StartbuttonY   int
	StartbuttonImg *ebiten.Image
	ExitbuttonX    int
	ExitbuttonY    int
	ExitbuttonImg  *ebiten.Image
}

func InitMenu(startbuttonImagePath, exitbuttonImagePath string) (*MainMenu, error) {
	startbuttonFile, err := os.Open(startbuttonImagePath)

	if err != nil {
		return nil, err
	}

	startbuttonFileDecoded, _, err := image.Decode(startbuttonFile)

	if err != nil {
		return nil, err
	}

	startbuttonImage := ebiten.NewImageFromImage(startbuttonFileDecoded)

	exitbuttonFile, err := os.Open(exitbuttonImagePath)

	if err != nil {
		return nil, err
	}

	exitbuttonFileDecoded, _, err := image.Decode(exitbuttonFile)

	if err != nil {
		return nil, err
	}

	exitbuttonImage := ebiten.NewImageFromImage(exitbuttonFileDecoded)

	mainM := &MainMenu{
		InMainMenu:     true,
		StartbuttonX:   10,
		StartbuttonY:   50,
		StartbuttonImg: startbuttonImage,
		ExitbuttonX:    25,
		ExitbuttonY:    75,
		ExitbuttonImg:  exitbuttonImage,
	}

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
