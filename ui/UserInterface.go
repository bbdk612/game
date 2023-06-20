package ui

type UI struct {
	HpBar *HealthBar
	WpBar *WeaponBar
	MiniM *MiniMap
}

func InitUI() (*UI, error) {
	hpBar, err := InitHealthBar("./assets/healthpoint.png")
	if err != nil {
		return nil, err
	}
	wpBar, err := InitWeaponBar("./assets/startWeapon.png")
	if err != nil {
		return nil, err
	}
	miniM, err := InitMiniMap("./gamemap/assets/common.png", "./gamemap/assets/shop.png", "./gamemap/assets/chest.png", "./gamemap/assets/boss.png", "./gamemap/assets/current.png")
	if err != nil {
		return nil, err
	}
	useri := &UI{
		HpBar: hpBar,
		WpBar: wpBar,
		MiniM: miniM,
	}
	return useri, nil
}
