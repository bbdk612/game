package items

import (
	"encoding/json"
	"fmt"
	ao "game/animatedobjects"
	"image"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Heal struct {
	Name      string
	Image     *ebiten.Image
	ImagePath string
	X, Y      float64
	health    int
}

func (h *Heal) Use(mh *ao.MainHero) {
	mh.Health += h.health
}

func (h *Heal) InActiveArea(MHx, MHy int) bool {
	var MHCenterX float64 = float64(MHx + 8)
	var MHCenterY float64 = float64(MHy + 8)

	var ItemCenterX float64 = h.X + 4
	var ItemCenterY float64 = h.Y + 4

	var distance float64 = math.Pow(MHCenterX-ItemCenterX, 2) + math.Pow(MHCenterY-ItemCenterY, 2)

	if distance <= 256 {
		return true
	}

	return false
}

func (h *Heal) InitHealItem(JSONPath string, x, y float64) (*Heal, error) {
	data, err := os.ReadFile(JSONPath)

	if err != nil {
		return nil, err
	}

	heal := &Heal{}
	jsonErr := json.Unmarshal(data, heal)

	if jsonErr != nil {
		return nil, jsonErr
	}

	heal.X = x
	heal.Y = y

	imageBytes, err := os.Open(heal.ImagePath)
	if err != nil {
		return nil, err
	}

	imageDecoded, _, err := image.Decode(imageBytes)

	if err != nil {
		return nil, err
	}

	image := ebiten.NewImageFromImage(imageDecoded)
	heal.Image = image
	fmt.Println(heal)
	return heal, nil
}
