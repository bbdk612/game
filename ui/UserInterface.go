package ui

import ()

type UI struct {
	healthBar *HeathBar
	weaponBar *WeaponBar
}

func InitUI() (*UI, error) {
	ui.healthBar, err := InitHealthBar("./assets/startWeapon.png")
	if err != nil {
		return nil, err
	}
	ui.weaponBar, err := InitWeaponBar("./assets/startWeapon.png")
	if err != nil {
		return nil, err
	}
	return ui, nil
}
