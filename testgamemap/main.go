package main

import (
	"fmt"

  "github.com/hajimehoshi/ebiten/v2"
 // "github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/bbdk612/game/gamemap"
)

type Game struct {
  Map gamemap.GameMap
  layers [][]int 
}

func (G *Game) Update() error {
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  
}

func main() {
  fmt.Println("hello, world")
}
 

