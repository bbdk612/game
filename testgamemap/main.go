package main

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/bbdk612/game/animatedobjects"
	"github.com/bbdk612/game/gamemap"
)

type Game struct {
	Map *gamemap.GameMap
	MH  *animatedobjects.MainHero
}

func (G *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		G.MH.AsePlayer.Play("walk")
		if G.MH.CanIGo("left", G.Map.GetCurrentChunk()) {
			fmt.Println("ok")
			if G.MH.GetTileCoor()%16 != 0 {
				x, y := G.MH.GetCoordinates()
				G.MH.SetCoordinates(x-2, y)
			} else {
				if chunk, ok := G.Map.CheckDirection("left"); ok {
					G.Map.ChangeCurrentChunk(chunk)
					G.MH.SetTileCoor(G.MH.GetTileCoor() + 15)
				}
			}
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		G.MH.AsePlayer.Play("walk")
		if G.MH.CanIGo("right", G.Map.GetCurrentChunk()) {
			fmt.Println("ok")
			if (G.MH.GetTileCoor()+1)%16 != 0 {
				x, y := G.MH.GetCoordinates()

				G.MH.SetCoordinates(x+2, y)
			} else {
				if chunk, ok := G.Map.CheckDirection("right"); ok {
					G.Map.ChangeCurrentChunk(chunk)
					G.MH.SetTileCoor(G.MH.GetTileCoor() - 15)
				}
			}
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		G.MH.AsePlayer.Play("walk")
		if G.MH.CanIGo("top", G.Map.GetCurrentChunk()) {
			if G.MH.GetTileCoor() < 16 {
				if chunk, ok := G.Map.CheckDirection("top"); ok {
					G.Map.ChangeCurrentChunk(chunk)
					G.MH.SetTileCoor(256 - G.MH.GetTileCoor())
				}
			} else {
				x, y := G.MH.GetCoordinates()
				G.MH.SetCoordinates(x, y-2)
			}
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		G.MH.AsePlayer.Play("walk")
		if G.MH.CanIGo("down", G.Map.GetCurrentChunk()) {
			if (G.MH.GetTileCoor() > 240) && (G.MH.GetTileCoor() < 256) {
				if chunk, ok := G.Map.CheckDirection("down"); ok {
					G.Map.ChangeCurrentChunk(chunk)
					x, _ := G.MH.GetCoordinates()
					G.MH.SetCoordinates(x, 0)
				}
			} else {
				x, y := G.MH.GetCoordinates()
				G.MH.SetCoordinates(x, y+2)
			}
		}
	} else {
		G.MH.AsePlayer.Play("stop")
	}

	G.MH.AsePlayer.Update(float32(1.0 / 60.0))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	xCount := (g.Map.SreenWidth / g.Map.TileSize)

	currentChunk := g.Map.GetCurrentChunk()

	for tileCoordinate, tileNumber := range currentChunk {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64((tileCoordinate%xCount)*g.Map.TileSize), float64((tileCoordinate/xCount)*g.Map.TileSize))

		screen.DrawImage(g.Map.GetTile(tileNumber), options)
	}

	options := &ebiten.DrawImageOptions{}

	x, y := g.MH.GetCoordinates()
	options.GeoM.Translate(float64(x), float64(y))

	sub := g.MH.Image.SubImage(image.Rect(g.MH.AsePlayer.CurrentFrameCoords()))

	screen.DrawImage(sub.(*ebiten.Image), options)
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

	mh, err := animatedobjects.InitMainHero(34, 16, 16)

	if err != nil {
		log.Fatal(err)
	}

	g := &Game{
		Map: M,
		MH:  mh,
	}

	ebiten.SetWindowSize(256*3, 256*3)
	ebiten.SetWindowTitle("test of Gamemap")

	g.MH.AsePlayer.Play("stop")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
