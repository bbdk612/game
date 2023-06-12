package mainmenu

import/ (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"os"
)

type MainMenu struct {
	inMainMenu     bool
	startbuttonX   int
	startbuttonY   int
	startbuttonImg *ebiten.Image
	exitbuttonX    int
	exitbuttonY    int
	exitbuttonImg  *ebiten.Image
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

	mainM := &HealthBar{
		inMainMenu:     true,
		startbuttonX:   10,
		startbuttonY:   50,
		startbuttonImg: startbuttonImage,
		exitbuttonX:    25,
		exitbuttonY:    75,
		exitbuttonImg:  exitbuttonImage,
	}

	return mainM, nil
}

func (mm *MainMenu) MenuStartGame() {
	mm.inMainMenu := false
}

func MenuExitGame() {
	os.exit
}
func (mm *MainMenu) GetMainMStartCoordinate() (int, int, int, int) {
	stbX := mm.startbuttonX
	stbY := mm.startbuttonY
	extX := mm.exitbuttonX
	extY := mm.exitbuttonY
	return stbX, stbY, extX, extY
}
