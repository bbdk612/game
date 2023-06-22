package ui

import (
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type ProgressBar struct {
	Font         font.Face
	x, y         int
	LevelCounter int
	ScoreCounter int
}

func (pb *ProgressBar) Magic() string {
	return fmt.Sprintf("Level %d Score %6d", pb.LevelCounter, pb.ScoreCounter)
}

func (pb *ProgressBar) GetCoordinates() (int, int) {
	return pb.x, pb.y
}

func InitProgressBar() (*ProgressBar, error) {
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

	pb := &ProgressBar{
		Font:         font,
		x:            200,
		y:            14,
		LevelCounter: 1,
		ScoreCounter: 0,
	}

	return pb, nil
}
