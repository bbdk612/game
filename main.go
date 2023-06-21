package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"game/animatedobjects"
	"game/gamemap"
	"game/menu"
	"game/ui"
	"game/weapons"
)

// Game struct contains a game objects
type Game struct {
	Bullets  [](*weapons.Bullet)
	GM       *gamemap.GameMap
	MH       *animatedobjects.MainHero
	UI       *ui.UI
	MM       *menu.MainMenu
	PM       *menu.PauseMenu
	DS       *menu.DeathScreen
	MenuRoll time.Time
}

// this function do all things for start game
func (g *Game) startGame() {
	g.MM.MenuStartGame()
	//generate map
	g.GM.CurrentRoom, g.GM.MiniMapPlan, g.GM.RoomList = g.GM.CurrentRoom.GenerateMap(12, 1, 1, 1)
	//set start room
	g.GM.RD = gamemap.JsonFileDecodeCurrentRoom(g.GM.CurrentRoom.RoomID, "./gamemap/assets/roomlist.json")
	g.GM.CurrentRoomTiles = g.GM.RD.GetCurrentRoomTileMap()
	g.GM.CurrentRoomTiles = g.GM.CurrentRoom.DeleteDoors(g.GM.CurrentRoomTiles)
	//set main hero properties
	g.MH.Health = g.MH.MaxHealth
	g.MenuRoll = time.Now()
}

// IsMoveKeyPressed checks on pressing a moving Key
func IsMoveKeyPressed() bool {
	moveKeys := [4](ebiten.Key){ebiten.KeyA, ebiten.KeyD, ebiten.KeyW, ebiten.KeyS}
	for _, moveKey := range moveKeys {
		if ebiten.IsKeyPressed(moveKey) {
			return true
		}
	}

	return false
}

// Update This Function Updates game data for game objects
func (g *Game) Update() error {
	if !(g.MM.InMainMenu) {
		if !(g.PM.InPauseMenu) {
			if !(g.DS.InDeathScreen) {
				if g.MH.Health <= 0 {
					g.DS.InDeathScreen = true
				}
				currTime := time.Now()
				dur, err := time.ParseDuration("300ms")
				if err != nil {
					return err
				}

				rlbck := dur.Milliseconds()

				if time.Duration(time.Duration(currTime.Sub(g.MenuRoll))).Milliseconds() > rlbck {
					if ebiten.IsKeyPressed(ebiten.KeyEscape) {
						g.PM.InPauseMenu = true
					}

					if IsMoveKeyPressed() {
						if ebiten.IsKeyPressed(ebiten.KeyA) {
							g.MH.AsePlayer.Play("walk")
							if g.MH.GetTileCoor()%16 == 0 {
								g.GM.CurrentRoom, g.GM.RD, g.GM.CurrentRoomTiles = g.GM.CurrentRoom.ChangeCurrentRoom("left")
								g.MH.SetTileCoor(g.MH.GetTileCoor() + 14)
							} else {
								g.MH.Move("left", g.GM.RD.GetCurrentRoomTileMap())
							}
						}
						if ebiten.IsKeyPressed(ebiten.KeyD) {
							g.MH.AsePlayer.Play("walk")
							if (g.MH.GetTileCoor()+1)%16 == 0 {
								g.GM.CurrentRoom, g.GM.RD, g.GM.CurrentRoomTiles = g.GM.CurrentRoom.ChangeCurrentRoom("right")
								g.MH.SetTileCoor(g.MH.GetTileCoor() - 14)
							} else {
								g.MH.Move("right", g.GM.RD.GetCurrentRoomTileMap())
							}
						}
						if ebiten.IsKeyPressed(ebiten.KeyW) {
							g.MH.AsePlayer.Play("walk")
							if _, y := g.MH.GetCoordinates(); y == 0 {
								g.GM.CurrentRoom, g.GM.RD, g.GM.CurrentRoomTiles = g.GM.CurrentRoom.ChangeCurrentRoom("top")
								x, _ := g.MH.GetCoordinates()
								g.MH.SetCoordinates(x, 224)
							} else {
								g.MH.Move("top", g.GM.RD.GetCurrentRoomTileMap())
							}
						}
						if ebiten.IsKeyPressed(ebiten.KeyS) {
							g.MH.AsePlayer.Play("walk")
							if (g.MH.GetTileCoor() > 240) && (g.MH.GetTileCoor() < 256) {
								g.GM.CurrentRoom, g.GM.RD, g.GM.CurrentRoomTiles = g.GM.CurrentRoom.ChangeCurrentRoom("down")
								x, _ := g.MH.GetCoordinates()
								g.MH.SetCoordinates(x, 16)
							} else {
								g.MH.Move("down", g.GM.RD.GetCurrentRoomTileMap())
							}
						}
					} else {
						g.MH.AsePlayer.Play("stop")
					}
					if ebiten.IsKeyPressed(ebiten.KeyR) {
						g.MH.GetCurrentWeapon().Reload()
					}
					for _, bullet := range g.Bullets {
						if bullet != nil {
							bullet.AsePlayer.Play("fly")
							bullet.Move()
						}
					}
					for i, bullet := range g.Bullets {
						if bullet != nil {
							mhX, mhY := g.MH.GetCoordinates()
							bullX, bullY := bullet.GetCoordinates()
							if (bullX >= float64(mhX)) && (bullY >= float64(mhY)) {
								if (bullX <= float64(mhX+16)) && (bullY <= float64(mhY+16)) {
									bullet = nil
									g.Bullets[i] = nil
									g.MH.Damage()
									continue
								}
							}
							if g.GM.CurrentRoomTiles[bullet.GetCurrentTile(16)] != 1 {
								bullet = nil
								g.Bullets[i] = nil
								continue
							}

						}
					}

					cursorX, cursorY := ebiten.CursorPosition()
					g.MH.GetCurrentWeapon().CalculateAngle(cursorX, cursorY)
					if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
						bull, err := g.MH.GetCurrentWeapon().Shoot(cursorX, cursorY, 16)
						if err != nil {
							log.Fatal(err)
						}

						if bull != nil {
							g.Bullets = append(g.Bullets, bull...)
						}
					}

					if ebiten.IsKeyPressed(ebiten.KeyH) {
						g.MH.Health = 6
					}
					if ebiten.IsKeyPressed(ebiten.KeyK) {
						g.MH.Health = 0
					}

				}
			} else {
				currTime := time.Now()
				dur, err := time.ParseDuration("300ms")
				if err != nil {
					return err
				}

				rlbck := dur.Milliseconds()

				if time.Duration(time.Duration(currTime.Sub(g.MenuRoll))).Milliseconds() > rlbck {
					cursorX, cursorY := ebiten.CursorPosition()
					if g.DS.ReturnToMMIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
						g.DS.DeathScreenReturnToMMGame(g.MM)
						g.MenuRoll = time.Now()
					}
				}
			}
		} else {
			currTime := time.Now()
			dur, err := time.ParseDuration("300ms")
			if err != nil {
				return err
			}

			rlbck := dur.Milliseconds()

			if time.Duration(time.Duration(currTime.Sub(g.MenuRoll))).Milliseconds() > rlbck {
				cursorX, cursorY := ebiten.CursorPosition()
				if g.PM.ContinueIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
					g.PM.PauseMenuContinueGame()
					g.MenuRoll = time.Now()
				}
				if g.PM.ExitToMMIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
					g.PM.PauseMenuExitToMMGame(g.MM)
					g.MenuRoll = time.Now()
				}
			}
		}
	} else {
		currTime := time.Now()
		dur, err := time.ParseDuration("300ms")
		if err != nil {
			return err
		}

		rlbck := dur.Milliseconds()

		if time.Duration(time.Duration(currTime.Sub(g.MenuRoll))).Milliseconds() > rlbck {
			cursorX, cursorY := ebiten.CursorPosition()
			if g.MM.StartIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				g.startGame()
			}
			if g.MM.ExitIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				g.MM.MenuExitGame()
			}
		}
	}

	g.MH.AsePlayer.Update(float32(1.0 / 60.0))
	return nil
}

// Draw drawing a game objects on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	if !(g.MM.InMainMenu) {
		if !(g.PM.InPauseMenu) {
			if !(g.DS.InDeathScreen) {
				//drawing a map
				xCount := (g.GM.ScreenWidth / g.GM.TileSize)

				for tileCoordinate, tileNumber := range g.GM.CurrentRoomTiles {
					options := &ebiten.DrawImageOptions{}
					options.GeoM.Translate(float64((tileCoordinate%xCount)*g.GM.TileSize), float64((tileCoordinate/xCount)*g.GM.TileSize))

					screen.DrawImage(g.GM.GetTile(tileNumber), options)
				}

				// drawing a personage

				optionsForMainHero := &ebiten.DrawImageOptions{}

				x, y := g.MH.GetCoordinates()
				optionsForMainHero.GeoM.Translate(float64(x), float64(y))

				sub := g.MH.Image.SubImage(image.Rect(g.MH.AsePlayer.CurrentFrameCoords()))

				screen.DrawImage(sub.(*ebiten.Image), optionsForMainHero)

				// drawing a gun
				optionsForWeapon := &ebiten.DrawImageOptions{}

				optionsForWeapon.GeoM.Rotate(g.MH.GetCurrentWeapon().GetAngle())
				oX, oY := g.MH.GetCurrentWeapon().GetOCoordinates()

				optionsForWeapon.GeoM.Translate(float64(oX), float64(oY))

				screen.DrawImage(g.MH.GetCurrentWeapon().Image, optionsForWeapon)

				//Draw a Bullets

				for _, bullet := range g.Bullets {
					if bullet != nil {
						opBullet := &ebiten.DrawImageOptions{}
						bX, bY := bullet.GetCoordinates()
						opBullet.GeoM.Translate(bX, bY)

						frame := bullet.Image.SubImage(image.Rect(bullet.AsePlayer.CurrentFrameCoords()))
						screen.DrawImage(frame.(*ebiten.Image), opBullet)
					}

				}
				// UI
				//HeathBar
				hpbX, hpbY := g.UI.HpBar.GetHpbStartCoordinate()
				opHPBar := &ebiten.DrawImageOptions{}
				opHPBar.GeoM.Translate(float64(hpbX), float64(hpbY))
				for i := 0; i < g.MH.Health; i++ {
					opHPBar.GeoM.Translate(float64(10), float64(0))
					screen.DrawImage(g.UI.HpBar.Image, opHPBar)
				}
				//WeaponBar
				wpbX, wpbY := g.UI.WpBar.GetWpbStartCoordinate()
				text.Draw(screen, g.UI.WpBar.GetAmmo(g.MH.GetCurrentWeapon().GetAmmo()), g.UI.WpBar.AmmoFont, wpbX, wpbY, color.White)
			} else {
				//DeathScreen
				stX, stY := g.DS.GetDathScreenStartCoordinate()
				opForReturnToMMButton := &ebiten.DrawImageOptions{}

				opForReturnToMMButton.GeoM.Translate(float64(84), float64(stY+15))

				subReturnMM := g.DS.ReturnToMMButtonImg.SubImage(image.Rect(g.DS.ReturnToMMButtonPlayer.CurrentFrameCoords()))
				screen.DrawImage(subReturnMM.(*ebiten.Image), opForReturnToMMButton)

				text.Draw(screen, "You Died", g.UI.WpBar.AmmoFont, stX+50, stY, color.White)

			}

		} else {
			//Pause menu
			stX, stY, extX, extY := g.PM.GetPauseMStartCoordinate()
			opForContinueButton := &ebiten.DrawImageOptions{}
			opForExitToMMButton := &ebiten.DrawImageOptions{}

			opForContinueButton.GeoM.Translate(float64(stX), float64(stY))
			opForExitToMMButton.GeoM.Translate(float64(extX), float64(extY))

			subContinue := g.PM.ContinueButtonImg.SubImage(image.Rect(g.PM.ContinueButtonPlayer.CurrentFrameCoords()))
			screen.DrawImage(subContinue.(*ebiten.Image), opForContinueButton)

			subExitMM := g.PM.ExitToMMButtonImg.SubImage(image.Rect(g.PM.ExitToMMButtonPlayer.CurrentFrameCoords()))
			screen.DrawImage(subExitMM.(*ebiten.Image), opForExitToMMButton)

			//Mini Map
			mmX, mmY := g.UI.MiniM.GetMiniMapStartCoordinate()
			startmmX := mmX
			currroomX, currroomY := g.GM.CurrentRoom.GetCurrentRoomCoordinate()
			for i := 0; i < len(g.GM.MiniMapPlan); i++ {
				mmX = startmmX
				for j := 0; j < len(g.GM.MiniMapPlan); j++ {
					opMiniM := &ebiten.DrawImageOptions{}
					opMiniM.GeoM.Translate(float64(mmX), float64(mmY))
					if g.GM.MiniMapPlan[j][i] != 0 {
						if j == currroomX && i == currroomY {
							screen.DrawImage(g.UI.MiniM.CurrentRoomImage, opMiniM)
						} else {
							if (g.GM.MiniMapPlan[j][i] > 100 && g.GM.MiniMapPlan[j][i] < 200) || (g.GM.MiniMapPlan[j][i] > 500 && g.GM.MiniMapPlan[j][i] < 600) {
								screen.DrawImage(g.UI.MiniM.CommonRoomImage, opMiniM)
							}
							if g.GM.MiniMapPlan[j][i] > 200 && g.GM.MiniMapPlan[j][i] < 300 {
								screen.DrawImage(g.UI.MiniM.ChestRoomImage, opMiniM)
							}
							if g.GM.MiniMapPlan[j][i] > 300 && g.GM.MiniMapPlan[j][i] < 400 {
								screen.DrawImage(g.UI.MiniM.BossRoomImage, opMiniM)
							}
							if g.GM.MiniMapPlan[j][i] > 400 && g.GM.MiniMapPlan[j][i] < 500 {
								screen.DrawImage(g.UI.MiniM.ShopRoomImage, opMiniM)
							}
						}
					}
					mmX = mmX + 9
				}
				mmY = mmY - 9
			}
		}
	} else {
		//Main menu
		stX, stY, extX, extY := g.MM.GetMainMStartCoordinate()
		opForStartButton := &ebiten.DrawImageOptions{}
		opForExitButton := &ebiten.DrawImageOptions{}
		opForStartButton.GeoM.Translate(float64(stX), float64(stY))
		opForExitButton.GeoM.Translate(float64(extX), float64(extY))

		subImageStartBtn := g.MM.StartButtonImg.SubImage(image.Rect(g.MM.StartButtonPlayer.CurrentFrameCoords()))
		screen.DrawImage(subImageStartBtn.(*ebiten.Image), opForStartButton)
		subImageExitBtn := g.MM.ExitButtonImg.SubImage(image.Rect(g.MM.ExitButtonPlayer.CurrentFrameCoords()))
		screen.DrawImage(subImageExitBtn.(*ebiten.Image), opForExitButton)

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.GM.ScreenWidth, g.GM.ScreenHeight
}

func main() {
	fmt.Println("hello, world")

	M, err := gamemap.InitGameMap("./gamemap/assets/tileset.png", 256, 256)
	if err != nil {
		fmt.Println(err)
	}

	mh, err := animatedobjects.InitMainHero(34, 16, 16, 2)

	if err != nil {
		log.Fatal(err)
	}

	ui, err := ui.InitUI()

	if err != nil {
		log.Fatal(err)
	}

	Menu, err := menu.InitMenu("./assets/start_button.json", "./assets/exitButton.json")

	if err != nil {
		log.Fatal(err)
	}

	pauseM, err := menu.InitPauseMenu("./assets/cotinue.json", "./assets/exitToMM.json")

	if err != nil {
		log.Fatal(err)
	}
	deathS, err := menu.InitDeathScreen("./assets/exitToMM.json")

	if err != nil {
		log.Fatal(err)
	}
	g := &Game{
		GM: M,
		MH: mh,
		UI: ui,
		MM: Menu,
		PM: pauseM,
		DS: deathS,
	}
	ebiten.SetWindowSize(256*2, 256*2)
	ebiten.SetWindowTitle("test of Gamemap")
	g.MH.AsePlayer.PlaySpeed = 0.5
	g.MH.AsePlayer.Play("stop")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
