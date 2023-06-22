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
	Bullets      [](*weapons.Bullet)
	GM           *gamemap.GameMap
	MH           *animatedobjects.MainHero
	MS           [](*animatedobjects.Monster)
	UI           *ui.UI
	AllM         *menu.AllMenus
	MenuRoll     time.Time
	Enemies      [](*animatedobjects.Monster)
	LevelCounter int
}

// this function do all things for start game
func (g *Game) startGame() {
	g.AllM.MM.MenuStartGame()
	//generate map
	g.GM.CurrentRoom, g.GM.MiniMapPlan, g.GM.RoomList = g.GM.CurrentRoom.GenerateMap(12, 1, 1, 2)
	//set start room
	g.GM.RD = gamemap.JsonFileDecodeCurrentRoom(g.GM.CurrentRoom.RoomID, "./gamemap/assets/roomlist.json")
	g.GM.CurrentRoomTiles = g.GM.RD.GetCurrentRoomTileMap()
	g.GM.CurrentRoomTiles = g.GM.CurrentRoom.DeleteDoors(g.GM.CurrentRoomTiles)
	g.MS = [](*animatedobjects.Monster){}
	g.Bullets = [](*weapons.Bullet){}
	//set main hero propertiesgo
	g.LevelCounter = 1
	g.MH.Health = g.MH.MaxHealth
	g.MenuRoll = time.Now()
}

func (g *Game) GenerateNextLevel() {
	//generate map
	g.GM.CurrentRoom, g.GM.MiniMapPlan, g.GM.RoomList = g.GM.CurrentRoom.GenerateMap(12, 1, 1, 2)
	//set start room
	g.GM.RD = gamemap.JsonFileDecodeCurrentRoom(g.GM.CurrentRoom.RoomID, "./gamemap/assets/roomlist.json")
	g.GM.CurrentRoomTiles = g.GM.RD.GetCurrentRoomTileMap()
	g.GM.CurrentRoomTiles = g.GM.CurrentRoom.DeleteDoors(g.GM.CurrentRoomTiles)
	g.MS = [](*animatedobjects.Monster){}
	g.Bullets = [](*weapons.Bullet){}
	g.LevelCounter = g.LevelCounter + 1
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
	if !(g.AllM.MM.InMainMenu) {
		if !(g.AllM.PM.InPauseMenu) {
			if !(g.AllM.DS.InDeathScreen) {
				if !(g.AllM.VS.InVictoryScreen) {
					if g.MH.Health <= 0 {
						g.AllM.DS.InDeathScreen = true
					}
					if g.MS != nil {
						MonsterInRoom := false
						for i := 0; i < len(g.MS); i++ {
							if g.MS[i] != nil {
								MonsterInRoom = true
								break
							}
						}
						if !MonsterInRoom && !g.GM.CurrentRoom.RoomIsCleaned {
							g.MS = [](*animatedobjects.Monster){}
							g.GM.CurrentRoom.RoomIsCleaned = true
							if g.GM.CurrentRoom.WayToNextLevel != nil {
								g.GM.CurrentRoom.WayToNextLevel.SpawnWayNextLevel()
							}
							g.GM.CurrentRoom.ChangeDoorsState(g.GM.CurrentRoomTiles, 1)
						}
					}
					x, y := g.MH.GetCoordinates()
					var Coordinates [][]float64
					Coordinates = append(Coordinates, []float64{float64(x), float64(y)})
					for _, monster := range g.MS {
						if monster != nil {
							MSx, MSy := monster.GetCoordinates()
							Coordinates = append(Coordinates, []float64{MSx, MSy})
						}
					}
					currTime := time.Now()
					dur, err := time.ParseDuration("300ms")
					if err != nil {
						return err
					}
					for i, bullet := range g.Bullets {
						if bullet != nil {
							bullX, bullY := bullet.GetCoordinates()
							for j, coordinate := range Coordinates {
								if (bullX >= coordinate[0]) && (bullY >= coordinate[1]) {
									if (bullX <= coordinate[0]+16) && (bullY <= coordinate[1]+16) {
										var remBull bool = false
										if j == 0 {
											g.MH.Damage()
											remBull = true

										} else if g.MS[j-1] != nil {
											fmt.Println(bullet.Damage)
											g.MS[j-1].Damage(bullet.Damage)
											remBull = true
											if g.MS[j-1].Health <= 0 {
												g.MS[j-1] = nil
											}
										}

										if remBull {
											bullet = nil
											g.Bullets[i] = nil
										}

										break
									}
								}
								if g.GM.CurrentRoomTiles[bullet.GetCurrentTile(16)] != 1 {
									bullet = nil
									g.Bullets[i] = nil
									break
								}
							}
						}
					}
					rlbck := dur.Milliseconds()

					if time.Duration(time.Duration(currTime.Sub(g.MenuRoll))).Milliseconds() > rlbck {
						if ebiten.IsKeyPressed(ebiten.KeyEscape) {
							g.AllM.PM.InPauseMenu = true
						}

						if IsMoveKeyPressed() {
							if ebiten.IsKeyPressed(ebiten.KeyA) {
								g.MH.AsePlayer.Play("walkleft")
								if g.MH.GetTileCoor()%16 == 0 {
									g.GM.CurrentRoom, g.GM.RD, g.GM.CurrentRoomTiles, g.MS = g.GM.CurrentRoom.ChangeCurrentRoom("left")
									g.Bullets = [](*weapons.Bullet){}
									g.MH.SetTileCoor(g.MH.GetTileCoor() + 14)
								} else {
									g.MH.Move("left", g.GM.RD.GetCurrentRoomTileMap(), Coordinates)
								}
							}
							if ebiten.IsKeyPressed(ebiten.KeyD) {
								g.MH.AsePlayer.Play("walkright")
								if (g.MH.GetTileCoor()+1)%16 == 0 {
									g.GM.CurrentRoom, g.GM.RD, g.GM.CurrentRoomTiles, g.MS = g.GM.CurrentRoom.ChangeCurrentRoom("right")
									g.Bullets = [](*weapons.Bullet){}
									g.MH.SetTileCoor(g.MH.GetTileCoor() - 14)
								} else {
									g.MH.Move("right", g.GM.RD.GetCurrentRoomTileMap(), Coordinates)
								}
							}
							if ebiten.IsKeyPressed(ebiten.KeyW) {
								if _, y := g.MH.GetCoordinates(); y == 0 {
									g.GM.CurrentRoom, g.GM.RD, g.GM.CurrentRoomTiles, g.MS = g.GM.CurrentRoom.ChangeCurrentRoom("top")
									g.Bullets = [](*weapons.Bullet){}
									x, _ := g.MH.GetCoordinates()
									g.MH.SetCoordinates(x, 224)
								} else {
									g.MH.Move("top", g.GM.RD.GetCurrentRoomTileMap(), Coordinates)
								}
							}
							if ebiten.IsKeyPressed(ebiten.KeyS) {
								if (g.MH.GetTileCoor() > 240) && (g.MH.GetTileCoor() < 256) {
									g.GM.CurrentRoom, g.GM.RD, g.GM.CurrentRoomTiles, g.MS = g.GM.CurrentRoom.ChangeCurrentRoom("down")
									g.Bullets = [](*weapons.Bullet){}
									x, _ := g.MH.GetCoordinates()
									g.MH.SetCoordinates(x, 16)
								} else {
									g.MH.Move("down", g.GM.RD.GetCurrentRoomTileMap(), Coordinates)
								}
							}
						} else {
							g.MH.AsePlayer.Play("wait")
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
						for _, monster := range g.MS {
							if monster != nil {
								bullets := monster.Actions(float64(x), float64(y), g.GM.CurrentRoomTiles, Coordinates)

								g.Bullets = append(g.Bullets, bullets...)

								monster.AsePlayer.Update(float32(1.0 / 60.0))
							}
						}

						if ebiten.IsKeyPressed(ebiten.KeyH) {
							g.MH.Health = 6
						}
						if ebiten.IsKeyPressed(ebiten.KeyK) {
							g.MH.Health = 0
						}
						if g.GM.CurrentRoom.Chest != nil {
							//fmt.Println("Chest good update: ", g.GM.CurrentRoom.Chest)
							g.GM.CurrentRoom.Chest.ChestPlayer.Update(float32(1.0 / 60.0))
						}
						if g.GM.CurrentRoom.WayToNextLevel != nil {
							g.GM.CurrentRoom.WayToNextLevel.WNLPlayer.Update(float32(1.0 / 60.0))
						}
						if g.GM.CurrentRoom.Chest != nil && (g.GM.CurrentRoom.Chest.InActiveZone(g.MH.GetCoordinates())) && (ebiten.IsKeyPressed(ebiten.KeyE)) {
							g.GM.CurrentRoom.Chest.Open()
							g.MenuRoll = time.Now()
						}
						if g.GM.CurrentRoom.WayToNextLevel != nil && (g.GM.CurrentRoom.WayToNextLevel.InActiveZone(g.MH.GetCoordinates())) && (ebiten.IsKeyPressed(ebiten.KeyE)) {
							g.GM.CurrentRoom.WayToNextLevel.WNLPlayer.Play("open")
							g.GM.CurrentRoom.WayToNextLevel.WNLPlayer.OnLoop = func() {
								animatedobjects.GoToNextLevel(g.AllM.VS)
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
						if g.AllM.VS.GoToNextLevelIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
							g.AllM.VS.VictoryScreenGoToNextLevel()
							g.GenerateNextLevel()
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
					if g.AllM.DS.ReturnToMMIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
						g.AllM.DS.DeathScreenReturnToMMGame(g.AllM.MM)
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
				if g.AllM.PM.ContinueIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
					g.AllM.PM.PauseMenuContinueGame()
					g.MenuRoll = time.Now()
				}
				if g.AllM.PM.ExitToMMIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
					g.AllM.PM.PauseMenuExitToMMGame(g.AllM.MM)
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
			if g.AllM.MM.StartIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				g.startGame()
			}
			if g.AllM.MM.ExitIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				g.AllM.MM.MenuExitGame()
			}
		}
	}
	g.MH.AsePlayer.Update(float32(1.0 / 60.0))
	return nil
}

// Draw drawing a game objects on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	if !(g.AllM.MM.InMainMenu) {
		if !(g.AllM.PM.InPauseMenu) {
			if !(g.AllM.DS.InDeathScreen) {
				if !(g.AllM.VS.InVictoryScreen) {
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

					text.Draw(screen, g.UI.WpBar.GetAmmo(g.MH.GetCurrentWeapon().GetAmmo()), g.UI.WpBar.AmmoFont, wpbX, wpbY, color.White)
					for _, monster := range g.MS {
						if monster != nil {
							optionsForMonster := &ebiten.DrawImageOptions{}
							opForMonstWeap := &ebiten.DrawImageOptions{}
							MSx, MSy := monster.GetCoordinates()
							optionsForMonster.GeoM.Translate(MSx, MSy)
							MSsub := monster.Image.SubImage(image.Rect(monster.AsePlayer.CurrentFrameCoords()))
							screen.DrawImage(MSsub.(*ebiten.Image), optionsForMonster)

							weapx, weapy := monster.Weapon.GetOCoordinates()
							opForMonstWeap.GeoM.Rotate(monster.Weapon.GetAngle())
							opForMonstWeap.GeoM.Translate(float64(weapx), float64(weapy))
							screen.DrawImage(monster.Weapon.Image, opForMonstWeap)
						}
					}
					//drawing chest
					if g.GM.CurrentRoom.Chest != nil {
						chestX, chestY := g.GM.CurrentRoom.Chest.GetCoordinates()
						optionsForChest := &ebiten.DrawImageOptions{}

						optionsForChest.GeoM.Translate(float64(chestX), float64(chestY))
						subChest := g.GM.CurrentRoom.Chest.ChestImage.SubImage(image.Rect(g.GM.CurrentRoom.Chest.ChestPlayer.CurrentFrameCoords()))
						screen.DrawImage(subChest.(*ebiten.Image), optionsForChest)
					}
					//drawing way to nest level
					if g.GM.CurrentRoom.WayToNextLevel != nil {
						wnlX, wnlY := g.GM.CurrentRoom.WayToNextLevel.GetCoordinates()
						optionsForWNL := &ebiten.DrawImageOptions{}

						optionsForWNL.GeoM.Translate(float64(wnlX), float64(wnlY))
						subWNL := g.GM.CurrentRoom.WayToNextLevel.WNLImage.SubImage(image.Rect(g.GM.CurrentRoom.WayToNextLevel.WNLPlayer.CurrentFrameCoords()))
						screen.DrawImage(subWNL.(*ebiten.Image), optionsForWNL)
					}
				} else {
					//VictoryScreen
					stX, stY := g.AllM.VS.GetVictoryScreenStartCoordinate()
					opForGoNextLevelButton := &ebiten.DrawImageOptions{}

					opForGoNextLevelButton.GeoM.Translate(float64(84), float64(stY+15))

					subGoNextLevel := g.AllM.VS.GoToNextLevelButtonImg.SubImage(image.Rect(g.AllM.VS.GoToNextLevelButtonPlayer.CurrentFrameCoords()))
					screen.DrawImage(subGoNextLevel.(*ebiten.Image), opForGoNextLevelButton)

					text.Draw(screen, fmt.Sprintf("You've reached level %d", g.LevelCounter), g.UI.WpBar.AmmoFont, stX+10, stY, color.White)

				}
			} else {
				//DeathScreen
				stX, stY := g.AllM.DS.GetDeathScreenStartCoordinate()
				opForReturnToMMButton := &ebiten.DrawImageOptions{}

				opForReturnToMMButton.GeoM.Translate(float64(84), float64(stY+15))

				subReturnMM := g.AllM.DS.ReturnToMMButtonImg.SubImage(image.Rect(g.AllM.DS.ReturnToMMButtonPlayer.CurrentFrameCoords()))
				screen.DrawImage(subReturnMM.(*ebiten.Image), opForReturnToMMButton)

				text.Draw(screen, "You Died", g.UI.WpBar.AmmoFont, stX+50, stY, color.White)

			}

		} else {
			//Pause menu
			stX, stY, extX, extY := g.AllM.PM.GetPauseMStartCoordinate()
			opForContinueButton := &ebiten.DrawImageOptions{}
			opForExitToMMButton := &ebiten.DrawImageOptions{}

			opForContinueButton.GeoM.Translate(float64(stX), float64(stY))
			opForExitToMMButton.GeoM.Translate(float64(extX), float64(extY))

			subContinue := g.AllM.PM.ContinueButtonImg.SubImage(image.Rect(g.AllM.PM.ContinueButtonPlayer.CurrentFrameCoords()))
			screen.DrawImage(subContinue.(*ebiten.Image), opForContinueButton)

			subExitMM := g.AllM.PM.ExitToMMButtonImg.SubImage(image.Rect(g.AllM.PM.ExitToMMButtonPlayer.CurrentFrameCoords()))
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
		stX, stY, extX, extY := g.AllM.MM.GetMainMStartCoordinate()
		opForStartButton := &ebiten.DrawImageOptions{}
		opForExitButton := &ebiten.DrawImageOptions{}
		opForStartButton.GeoM.Translate(float64(stX), float64(stY))
		opForExitButton.GeoM.Translate(float64(extX), float64(extY))

		subImageStartBtn := g.AllM.MM.StartButtonImg.SubImage(image.Rect(g.AllM.MM.StartButtonPlayer.CurrentFrameCoords()))
		screen.DrawImage(subImageStartBtn.(*ebiten.Image), opForStartButton)
		subImageExitBtn := g.AllM.MM.ExitButtonImg.SubImage(image.Rect(g.AllM.MM.ExitButtonPlayer.CurrentFrameCoords()))
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

	enemies := [](*animatedobjects.Monster){}

	mh, err := animatedobjects.InitMainHero(34, 16, 16, 2)

	if err != nil {
		log.Fatal(err)
	}

	ui, err := ui.InitUI()

	if err != nil {
		log.Fatal(err)
	}
	allM, err := menu.InitAllMenus()
	if err != nil {
		log.Fatal(err)
	}
	g := &Game{
		GM:   M,
		MH:   mh,
		MS:   enemies,
		UI:   ui,
		AllM: allM,
	}
	ebiten.SetWindowSize(256*2, 256*2)
	ebiten.SetWindowTitle("test of Gamemap")
	g.MH.AsePlayer.PlaySpeed = 0.5
	g.MH.AsePlayer.Play("wait")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
