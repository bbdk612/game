package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
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

	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
	}
  

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

	g := &Game{
		Map: M,
	}

	ebiten.SetWindowSize(256*3, 256*3)
	ebiten.SetWindowTitle("test of Gamemap")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
