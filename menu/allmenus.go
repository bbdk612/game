package menu

import (
	"log"
)

type AllMenus struct {
	DS *DeathScreen
	MM *MainMenu
	PM *PauseMenu
	VS *VictoryScreen
}

func InitAllMenus() (*AllMenus, error) {
	Menu, err := InitMenu("./assets/start_button.json", "./assets/exitButton.json")

	if err != nil {
		log.Fatal(err)
	}

	pauseM, err := InitPauseMenu("./assets/cotinue.json", "./assets/exitToMM.json")

	if err != nil {
		log.Fatal(err)
	}
	deathS, err := InitDeathScreen("./assets/exitToMM.json")

	if err != nil {
		log.Fatal(err)

	}
	victoryS, err := InitVictoryScreen("./assets/gonext.json")

	if err != nil {
		log.Fatal(err)

	}
	allM := &AllMenus{
		DS: deathS,
		MM: Menu,
		PM: pauseM,
		VS: victoryS,
	}
	return allM, nil
}
