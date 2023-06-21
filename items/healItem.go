package items

import (
	"encoding/json"
	"image"
	"math"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Item struct {
	Type      string
	Image     *ebiten.Image
	JsonPath  string
	ImagePath string
	X, Y      float64
}

func (h *Item) PickUp() (string, string) {
	return h.Type, h.JsonPath
}

func (h *Item) InActiveArea(MHx, MHy int) bool {
	var MHCenterX float64 = float64(MHx + 8)
	var MHCenterY float64 = float64(MHy + 8)

	var ItemCenterX float64 = h.X + 4
	var ItemCenterY float64 = h.Y + 4

	var distance float64 = math.Pow(MHCenterX-ItemCenterX, 2) + math.Pow(MHCenterY-ItemCenterY, 2)

	if distance <= 625 {
		return true
	}

	return false
}

func InitItem(Type string, JSONPath string, x, y float64) (*Item, error) {
	data, err := os.ReadFile(JSONPath)

	if err != nil {
		return nil, err
	}

	heal := &Item{}
	jsonErr := json.Unmarshal(data, heal)
	heal.JsonPath = JSONPath
	heal.Type = Type
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
	return heal, nil
}
