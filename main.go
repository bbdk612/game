package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"game/animatedobjects"
	"game/gamemap"
	"game/menu"
	"game/ui"
)

// Game struct contains a game objects
type Game struct {
	Bullets          [](*animatedobjects.Bullet)
	CurrentRoom      *gamemap.GameMap
	RoomList         [](*gamemap.GameMap)
	MapOptions       *gamemap.GameMapOptions
	RD               *gamemap.RoomData
	CurrentRoomTiles []int
	MH               *animatedobjects.MainHero
	UI               *ui.UI
	MM               *menu.MainMenu
	PM               *menu.PauseMenu
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
			if ebiten.IsKeyPressed(ebiten.KeyEscape) {
				g.PM.InPauseMenu = true
			}

			if IsMoveKeyPressed() {
				if ebiten.IsKeyPressed(ebiten.KeyA) {
					g.MH.AsePlayer.Play("walk")
					if g.MH.GetTileCoor()%16 == 0 {
						g.CurrentRoom = g.CurrentRoom.ChangeCurrentRoom("left")
						g.RD = gamemap.JsonFileDecodeCurrentRoom(g.CurrentRoom.RoomID)
						g.CurrentRoomTiles = g.RD.GetCurrentRoomTileMap()
						g.CurrentRoomTiles = g.CurrentRoom.DeleteDoors(g.CurrentRoomTiles)
						g.MH.SetTileCoor(g.MH.GetTileCoor() + 14)
					} else {
						g.MH.Move("left", g.RD.GetCurrentRoomTileMap())
					}
				}
				if ebiten.IsKeyPressed(ebiten.KeyD) {
					g.MH.AsePlayer.Play("walk")
					if (g.MH.GetTileCoor()+1)%16 == 0 {
						g.CurrentRoom = g.CurrentRoom.ChangeCurrentRoom("right")
						g.RD = gamemap.JsonFileDecodeCurrentRoom(g.CurrentRoom.RoomID)
						fmt.Println("good: ", g.RD)
						g.CurrentRoomTiles = g.RD.GetCurrentRoomTileMap()
						g.CurrentRoomTiles = g.CurrentRoom.DeleteDoors(g.CurrentRoomTiles)
						g.MH.SetTileCoor(g.MH.GetTileCoor() - 14)
					} else {
						g.MH.Move("right", g.RD.GetCurrentRoomTileMap())
					}
				}
				if ebiten.IsKeyPressed(ebiten.KeyW) {
					g.MH.AsePlayer.Play("walk")
					if _, y := g.MH.GetCoordinates(); y == 0 {
						g.CurrentRoom = g.CurrentRoom.ChangeCurrentRoom("top")
						g.RD = gamemap.JsonFileDecodeCurrentRoom(g.CurrentRoom.RoomID)
						g.CurrentRoomTiles = g.RD.GetCurrentRoomTileMap()
						g.CurrentRoomTiles = g.CurrentRoom.DeleteDoors(g.CurrentRoomTiles)
						x, _ := g.MH.GetCoordinates()
						g.MH.SetCoordinates(x, 224)
					} else {
						g.MH.Move("top", g.RD.GetCurrentRoomTileMap())
					}
				}
				if ebiten.IsKeyPressed(ebiten.KeyS) {
					g.MH.AsePlayer.Play("walk")
					if (g.MH.GetTileCoor() > 240) && (g.MH.GetTileCoor() < 256) {
						g.CurrentRoom = g.CurrentRoom.ChangeCurrentRoom("down")
						g.RD = gamemap.JsonFileDecodeCurrentRoom(g.CurrentRoom.RoomID)
						g.CurrentRoomTiles = g.RD.GetCurrentRoomTileMap()
						g.CurrentRoomTiles = g.CurrentRoom.DeleteDoors(g.CurrentRoomTiles)
						x, _ := g.MH.GetCoordinates()
						g.MH.SetCoordinates(x, 16)
					} else {
						g.MH.Move("down", g.RD.GetCurrentRoomTileMap())
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
					if g.CurrentRoomTiles[bullet.GetCurrentTile(16)] != 1 {
						bullet = nil
						g.Bullets[i] = nil
						continue
					}

				}
			}

			cursorX, cursorY := ebiten.CursorPosition()
			g.MH.GetCurrentWeapon().CalculateAngle(cursorX, cursorY)
			if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
				bull, err := g.MH.GetCurrentWeapon().Shoot(cursorX, cursorY, "./assets/bullet.json", 16)
				if err != nil {
					log.Fatal(err)
				}

				if bull != nil {
					g.Bullets = append(g.Bullets, bull)
				}
			}

			if ebiten.IsKeyPressed(ebiten.KeyH) {
				g.MH.Health = 6
			}
		} else {
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				//charX, charY := g.MH.GetCoordinates()
				g.PM.PauseMenuContinueGame()
			}
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
				//charX, charY := g.MH.GetCoordinates()
				g.PM.PauseMenuExitToMMGame(g.MM)
			}
		}
	} else {
		cursorX, cursorY := ebiten.CursorPosition()
		if g.MM.StartIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			//charX, charY := g.MH.GetCoordinates()
			g.MM.MenuStartGame()
			g.CurrentRoom, g.RoomList = g.CurrentRoom.GenerateMap(10, 0, 0, 0)
			fmt.Println("good: ", g.CurrentRoom.RoomID)
			g.RD = gamemap.JsonFileDecodeCurrentRoom(g.CurrentRoom.RoomID)
			g.CurrentRoomTiles = g.RD.GetCurrentRoomTileMap()
			g.CurrentRoomTiles = g.CurrentRoom.DeleteDoors(g.CurrentRoomTiles)
		}
		if g.MM.ExitIsActive(cursorX, cursorY) && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			//charX, charY := g.MH.GetCoordinates()
			g.MM.MenuExitGame()
		}
	}

	g.MH.AsePlayer.Update(float32(1.0 / 60.0))

	return nil
}

// Draw drawing a game objects on the screen
func (g *Game) Draw(screen *ebiten.Image) {
	if !(g.MM.InMainMenu) {
		if !(g.PM.InPauseMenu) {
			//drawing a map
			xCount := (g.MapOptions.ScreenWidth / g.MapOptions.TileSize)

			for tileCoordinate, tileNumber := range g.CurrentRoomTiles {
				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64((tileCoordinate%xCount)*g.MapOptions.TileSize), float64((tileCoordinate/xCount)*g.MapOptions.TileSize))

				screen.DrawImage(g.MapOptions.GetTile(tileNumber), options)
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
			for i := 1; i < g.MH.Health; i++ {
				opHPBar := &ebiten.DrawImageOptions{}
				opHPBar.GeoM.Translate(float64(hpbX), float64(hpbY))
				screen.DrawImage(g.UI.HpBar.Image, opHPBar)
				hpbX = hpbX + 10
			}
			//WeaponBar
			wpbX, wpbY := g.UI.WpBar.GetWpbStartCoordinate()
			text.Draw(screen, g.UI.WpBar.GetAmmo(g.MH.GetCurrentWeapon().GetAmmo()), g.UI.WpBar.AmmoFont, wpbX, wpbY, color.White)
		} else {
			//Main menu
			stX, stY, extX, extY := g.MM.GetMainMStartCoordinate()
			opForContinueButton := &ebiten.DrawImageOptions{}
			opForExitToMMButton := &ebiten.DrawImageOptions{}

			opForContinueButton.GeoM.Translate(float64(stX), float64(stY))
			opForExitToMMButton.GeoM.Translate(float64(extX), float64(extY))

			subContinue := g.PM.ContinueButtonImg.SubImage(image.Rect(g.PM.ContinueButtonPlayer.CurrentFrameCoords()))
			screen.DrawImage(subContinue.(*ebiten.Image), opForContinueButton)

			subExitMM := g.PM.ExitToMMButtonImg.SubImage(image.Rect(g.PM.ExitToMMButtonPlayer.CurrentFrameCoords()))
			screen.DrawImage(subExitMM.(*ebiten.Image), opForContinueButton)
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
	return g.MapOptions.ScreenWidth, g.MapOptions.ScreenHeight
}

func main() {
	fmt.Println("hello, world")

	M, err := gamemap.InitGameMap("./gamemap/assets/tileset.png", 256, 256)
	if err != nil {
		fmt.Println(err)
	}

	mh, err := animatedobjects.InitMainHero(34, 16, 16, 4)

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
	g := &Game{
		MapOptions: M,
		MH:         mh,
		UI:         ui,
		MM:         Menu,
		PM:         pauseM,
	}
	ebiten.SetWindowSize(256*2, 256*2)
	ebiten.SetWindowTitle("test of Gamemap")
	g.MH.AsePlayer.PlaySpeed = 0.5
	g.MH.AsePlayer.Play("stop")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
