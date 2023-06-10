package ui

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)
type WeaponBar struct {
	startX, startY                   int
	currentAmmo                      int
	youHeveAmmo                      int
	boardsImage                      *ebiten.Image
}

//func GetCurrentAmmo()(int int){
//	creWeap
//	return maxAmmo,currentAmmo
//}
func InitWeaponBar (weaponImagePath,boardsImagePath string,currentAmmo,maxAmmo int) (*WeaponBar, error){

	//init boards img

	boardsFile, err := os.Open(boardsImagePath)

	if err != nil {
		return nil, err
	}

	boardsFileDecoded, _, err := image.Decode(boardsFile)

	if err != nil {
		return nil, err
	}

	boardsImage := ebiten.NewImageFromImage(boardsFileDecoded)

	//struct init

	wpB := &WeaponBar{
		startX:    10,
		startY:    10,
		currentAmmo: currentAmmo,
		youHeveAmmo:     maxAmmo,
		boardsImage: boardsImage,
	}

	return wpB, nil
}
