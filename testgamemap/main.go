package main

import (
	"fmt"

	"github.com/bbdk612/game/gamemap"
)

const (
  screenWidth = 240
  screenHeight = 240
)

const (
  tileSize = 16
)

type Game struct {
  Map gamemap.GameMap
  layers [][]int 
}
 

