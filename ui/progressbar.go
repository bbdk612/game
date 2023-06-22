package ui

import (
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type ProgressBar struct {
	Font           font.Face
	LevelX, LevelY int
	ScoreX, ScoreY int
	LevelCounter   int
	ScoreCounter   int
}

func (pb *ProgressBar) GetLevel() string {
	return fmt.Sprintf("Level %d", pb.LevelCounter)
}
func (pb *ProgressBar) GetScore() string {
	return fmt.Sprintf("Score %06d", pb.ScoreCounter)
}

func (pb *ProgressBar) GetLevelCoordinates() (int, int) {
	return pb.LevelX, pb.LevelY
}
func (pb *ProgressBar) GetScoreCoordinates() (int, int) {
	return pb.ScoreX, pb.ScoreY
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
		LevelX:       200,
		LevelY:       13,
		ScoreX:       5,
		ScoreY:       253,
		LevelCounter: 1,
		ScoreCounter: 0,
	}

	return pb, nil
}
