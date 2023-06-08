package ui

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/bbdk612/game/ui"
)

type UI struct {
	heathBar *HeathBar
}

func (ui *UI) InitUI() {
	InitHealthBar(<path to hpBar>)

}
