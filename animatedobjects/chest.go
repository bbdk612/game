package animatedobjects

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"game/items"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type Chest struct {
	Opened         bool
	x, y           int
	ChestPlayer    *goaseprite.Player
	ChestFile      *goaseprite.File
	ChestImage     *ebiten.Image
	tilecoordinate int
}

func (ch *Chest) randomItem() string {
	items := []string{
		"heal",
		"gun",
		"shotgun",
	}

	return items[rand.Intn(len(items))]
}

func (ch *Chest) Open() *items.Item {
	if !ch.Opened {
		fmt.Println("Chest good open: ", ch)
		ch.Opened = true
		ch.ChestPlayer.Play("open")
		ch.ChestPlayer.OnLoop = func() {
			ch.ChestPlayer.Play("stop")
		}
		//r := ch.randomItem()
		r := "heal"
		var jsonPath string
		switch r {
		case "heal":
			jsonPath = "./assets/heal.json"
		case "gun":
			jsonPath = "./assets/gun.json"
		case "shotgun":
			jsonPath = "./assets/shotgun.json"
		}

		spawnX, spawnY := ch.spawnCoordinates()

		i, err := items.InitItem(r, jsonPath, float64(spawnX), float64(spawnY))
		if err != nil {
			log.Fatal(err)
		}

		return i
	}
	return nil

}

func (ch *Chest) spawnCoordinates() (int, int) {
	x, y := ch.GetCoordinates()
	return x + 12, y + 4
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
		Opened:         false,
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
