package ui

import (
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type LevelBar struct {
	Font font.Face
	x, y int
}

func (lb *LevelBar) Magic(CurrentLevel int) string {
	return fmt.Sprintf("Level %d", CurrentLevel)
}

func (lb *LevelBar) GetCoordinates() (int, int) {
	return lb.x, lb.y
}

func InitLevelBar() (*LevelBar, error) {
	fontBytes, err := os.ReadFile("./assets/font.ttf")
	if err != nil {
		return nil, err
	}

	fontParsed, err := opentype.Parse(fontBytes)

	if err != nil {
		return nil, err
	}

	font, err := opentype.NewFace(fontParsed, &opentype.FaceOptions{
		Size:    10,
		DPI:     150,
		Hinting: font.HintingVertical,
	})

	lb := &LevelBar{
		Font: font,
		x:    200,
		y:    14,
	}

	return lb, nil
}
