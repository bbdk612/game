package ui

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/bbdk612/game/ui"
)

type UI struct {
	healthBar *HeathBar
}

func InitUI() (*UI,error){
	ui.healthBar, err :=InitHealthBar(<path to hpBar>)
	if err != nil {
		return nil, err
	}
	return ui, nil;
}
