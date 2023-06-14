package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"game/animatedobjects"
	"game/gamemap"
	"game/menu"
	"game/ui"
)

type Game struct {
	Bullets [](*animatedobjects.Bullet)
	Map     *gamemap.GameMap
	MH      *animatedobjects.MainHero
	UI      *ui.UI
	MM      *menu.MainMenu
	PM      *menu.PauseMenu
}

func IsMoveKeyPressed() bool {
	moveKeys := [4](ebiten.Key){ebiten.KeyA, ebiten.KeyD, ebiten.KeyW, ebiten.KeyS}
	for _, moveKey := range moveKeys {
		if ebiten.IsKeyPressed(moveKey) {
			return true
		}
	}

	return false
}

func (G *Game) Update() error {
	if !(G.MM.InMainMenu) {
		if !(G.PM.InPauseMenu) {
			if ebiten.IsKeyPressed(ebiten.KeyEscape) {
				G.PM.InPauseMenu = true
			}

			if IsMoveKeyPressed() {
				if ebiten.IsKeyPressed(ebiten.KeyA) {
					G.MH.AsePlayer.Play("walk")
					if G.MH.GetTileCoor()%16 == 0 {
						if chunk, ok := G.Map.CheckDirection("left"); ok {
							G.Map.ChangeCurrentChunk(chunk)
							G.MH.SetTileCoor(G.MH.GetTileCoor() + 15)
						}
					} else if G.MH.CanIGo("left", G.Map.GetCurrentChunk()) {
						G.MH.Move("left", G.Map.GetCurrentChunk())
					}
				}
				if ebiten.IsKeyPressed(ebiten.KeyD) {
					G.MH.AsePlayer.Play("walk")
					if (G.MH.GetTileCoor()+1)%16 == 0 {
						if chunk, ok := G.Map.CheckDirection("right"); ok {
							G.Map.ChangeCurrentChunk(chunk)
							G.MH.SetTileCoor(G.MH.GetTileCoor() - 15)
						}
					} else {
						G.MH.Move("right", G.Map.GetCurrentChunk())
					}
				}
				if ebiten.IsKeyPressed(ebiten.KeyW) {
					G.MH.AsePlayer.Play("walk")
					if _, y := G.MH.GetCoordinates(); y == 0 {
						if chunk, ok := G.Map.CheckDirection("top"); ok {
							G.Map.ChangeCurrentChunk(chunk)
							G.MH.SetTileCoor(256 - (G.MH.GetTileCoor() - 2))
						}
					} else {
						G.MH.Move("top", G.Map.GetCurrentChunk())
					}
				}
				if ebiten.IsKeyPressed(ebiten.KeyS) {
					G.MH.AsePlayer.Play("walk")
					if (G.MH.GetTileCoor() > 240) && (G.MH.GetTileCoor() < 256) {
						if chunk, ok := G.Map.CheckDirection("down"); ok {
							G.Map.ChangeCurrentChunk(chunk)
							x, _ := G.MH.GetCoordinates()
							G.MH.SetCoordinates(x, 0)
						}
					} else {
						G.MH.Move("down", G.Map.GetCurrentChunk())
					}
				}
			}
		} else {
			G.MH.AsePlayer.Play("stop")
		}

		if ebiten.IsKeyPressed(ebiten.KeyR) {
			G.MH.GetCurrentWeapon().Reload()
		}
		for _, bullet := range G.Bullets {
			if bullet != nil {
				bullet.AsePlayer.Play("fly")
				bullet.Move()
			}
		}
		chunk := G.Map.GetCurrentChunk()
		for i, bullet := range G.Bullets {
			if bullet != nil {
				mhX, mhY := G.MH.GetCoordinates()
				bullX, bullY := bullet.GetCoordinates()
				if (bullX >= float64(mhX)) && (bullY >= float64(mhY)) {
					if (bullX <= float64(mhX+16)) && (bullY <= float64(mhY+16)) {
						bullet = nil
						G.Bullets[i] = nil
						G.MH.Damage()
						continue
					}
				}
				if chunk[bullet.GetCurrentTile(16)] != 1 {
					bullet = nil
					G.Bullets[i] = nil
					continue
				}

			}
		}

		cursorX, cursorY := ebiten.CursorPosition()
		G.MH.GetCurrentWeapon().CalculateAngle(cursorX, cursorY)
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			bull, err := G.MH.GetCurrentWeapon().Shoot(cursorX, cursorY, "./assets/bullet.json", 16)
			if err != nil {
				log.Fatal(err)
			}

			if bull != nil {
				G.Bullets = append(G.Bullets, bull)
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyH) {
			G.MH.Health = 6
		} else {
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				//charX, charY := G.MH.GetCoordinates()
				G.PM.PauseMenuContinueGame()
			}
			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
				//charX, charY := G.MH.GetCoordinates()
				G.PM.PauseMenuExitToMMGame()
			}
		}
	} else {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			//charX, charY := G.MH.GetCoordinates()
			G.MM.MenuStartGame()
		}
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
			//charX, charY := G.MH.GetCoordinates()
			G.MM.MenuExitGame()
		}
	}

	G.MH.AsePlayer.Update(float32(1.0 / 60.0))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !(g.MM.InMainMenu) {
		if !(g.PM.InPauseMenu) {
			//drawing a map
			xCount := (g.Map.SreenWidth / g.Map.TileSize)

			currentChunk := g.Map.GetCurrentChunk()

			for tileCoordinate, tileNumber := range currentChunk {
				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64((tileCoordinate%xCount)*g.Map.TileSize), float64((tileCoordinate/xCount)*g.Map.TileSize))

				screen.DrawImage(g.Map.GetTile(tileNumber), options)
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
			for i := 1; i < g.UI.HpBar.HealthNumber; i++ {
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
			screen.DrawImage(g.PM.ContinuebuttonImg, opForContinueButton)
			msg1 := fmt.Sprintf("Continue")
			ebitenutil.DebugPrintAt(screen, msg1, stX+100, stY)
			screen.DrawImage(g.PM.ExitToMMbuttonImg, opForExitToMMButton)
			msg2 := fmt.Sprintf("Exit To Main Menu")
			ebitenutil.DebugPrintAt(screen, msg2, extX+100, extY)
		}
	} else {
		//Main menu
		stX, stY, extX, extY := g.MM.GetMainMStartCoordinate()
		opForStartButton := &ebiten.DrawImageOptions{}
		opForExitButton := &ebiten.DrawImageOptions{}
		opForStartButton.GeoM.Translate(float64(stX), float64(stY))
		opForExitButton.GeoM.Translate(float64(extX), float64(extY))
		screen.DrawImage(g.MM.StartbuttonImg, opForStartButton)
		msg1 := fmt.Sprintf("Start")
		ebitenutil.DebugPrintAt(screen, msg1, stX+100, stY)
		screen.DrawImage(g.MM.ExitbuttonImg, opForExitButton)
		msg2 := fmt.Sprintf("Exit")
		ebitenutil.DebugPrintAt(screen, msg2, extX+100, extY)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Map.SreenWidth, g.Map.SreenHeight
}

func main() {
	fmt.Println("hello, world")
	chunks := [][]int{
		{
			4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
			4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 4,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 4,
			4, 1, 2, 2, 3, 3, 2, 2, 2, 2, 3, 3, 2, 2, 1, 4,
			4, 1, 2, 2, 3, 3, 2, 2, 2, 2, 3, 3, 2, 2, 1, 4,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1,
			4, 1, 2, 2, 3, 3, 2, 2, 2, 2, 3, 3, 2, 2, 1, 4,
			4, 1, 2, 2, 3, 3, 2, 2, 2, 2, 3, 3, 2, 2, 1, 4,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 4,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 4,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 4,
			4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4,
			4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
		},
		{
			4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
			4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 1, 1, 1, 1, 1, 1, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 1, 1, 2, 1, 1, 1, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 4, 4, 1, 2, 1, 4, 4, 4, 4, 4,
			4, 4, 4, 4, 4, 4, 4, 4, 1, 2, 1, 4, 4, 4, 4, 4,
			4, 4, 4, 4, 4, 4, 4, 4, 1, 2, 1, 4, 4, 4, 4, 4,
			4, 4, 4, 4, 4, 4, 4, 4, 1, 2, 1, 4, 4, 4, 4, 4,
		},
		{
			4, 4, 4, 4, 4, 4, 4, 4, 1, 2, 1, 4, 4, 4, 4, 4,
			4, 4, 4, 4, 4, 4, 4, 4, 1, 2, 1, 4, 4, 4, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 1, 1, 2, 1, 1, 1, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 2, 2, 2, 2, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 1, 1, 1, 1, 1, 1, 1, 1, 4, 4,
			4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
			4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
			4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
			4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
		},
	}

	for i := range chunks {
		for j := range chunks[i] {
			chunks[i][j]--
		}
	}

	roadsTo := []map[string]int{
		{
			"right": 1,
		},
		{
			"left": 0,
			"down": 2,
		},
		{
			"top": 1,
		},
	}

	M, err := gamemap.NewGameMap(chunks, 0, roadsTo, 256, 256)
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

	menu, err := menu.InitMenu("./assets/healthpoint.png", "./assets/healthpoint.png")

	if err != nil {
		log.Fatal(err)
	}

	pauseM, err := menu.InitPauseMenu("./assets/healthpoint.png", "./assets/healthpoint.png")

	if err != nil {
		log.Fatal(err)
	}
	g := &Game{
		Map: M,
		MH:  mh,
		UI:  ui,
		MM:  menu,
		PM:  pauseM,
	}
	ebiten.SetWindowSize(256*3, 256*3)
	ebiten.SetWindowTitle("test of Gamemap")
	g.MH.AsePlayer.PlaySpeed = 0.5
	g.MH.AsePlayer.Play("stop")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
