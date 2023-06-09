package ui

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/bbdk612/game/ui"
)

type UI struct {
	healthBar *HeathBar
	weaponBar *WeaponBar
}

func InitUI() (*UI,error){
	ui.healthBar, err :=InitHealthBar(<path to hpBar>)
	if err != nil {
		return nil, err
	}
	ui.weaponBar, err :=InitWeaponBar(<path to weapBar>)
	if err != nil {
		return nil, err
	}
	return ui, nil;
}
