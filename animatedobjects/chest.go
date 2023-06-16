package animatedobjects

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type Chest struct {
	opened         bool
	ChestPlayer    *goaseprite.Player
	ChestFile      *goaseprite.File
	ChestImage     *ebiten.Image
	tilecoordinate int
}

func (ch *Chest) Open(items []string) {
	if !ch.opened {
		number := rand.Intn(len(items))
		fmt.Println(items[number])
		ch.opened = true
		ch.ChestPlayer.Play("open")
		stopIt := func() {
			ch.ChestPlayer.Play("stop")
		}
		ch.ChestPlayer.OnLoop = stopIt
	}
}

func (ch *Chest) GetCoordinates() (int, int) {
	var x int = ((ch.tilecoordinate % 16) * 16)
	var y int = ((ch.tilecoordinate / 16) * 16)

	return x, y
}

func (ch *Chest) InActiveZone(x, y int) bool {
	var centerChestX int = ((ch.tilecoordinate % 16) * 16) + 16
	var centerChestY int = ((ch.tilecoordinate / 16) * 16) + 8

	var objectCenterX int = x + 8
	var objectCenterY int = y + 8

	var distance float64 = math.Pow(float64(centerChestX-objectCenterX), 2) + math.Pow(float64(centerChestY-objectCenterY), 2)

	if distance <= 625 {
		return true
	}

	return false
}

func InitNewChest(jsonPath string, tilecoordinate int) (*Chest, error) {
	chest := &Chest{
		opened:         false,
		ChestFile:      goaseprite.Open(jsonPath),
		tilecoordinate: tilecoordinate,
	}

	chest.ChestPlayer = chest.ChestFile.CreatePlayer()

	img, _, err := ebitenutil.NewImageFromFile(chest.ChestFile.ImagePath)
	if err != nil {
		return nil, err
	}
	chest.ChestImage = img

	chest.ChestPlayer.Play("wait")

	return chest, nil
}
