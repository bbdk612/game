package ui

type UI struct {
	HpBar *HealthBar
	WpBar *WeaponBar
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
	useri := &UI{
		HpBar: hpBar,
		WpBar: wpBar,
	}
	return useri, nil
}
