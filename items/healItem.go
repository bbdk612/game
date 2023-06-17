package items

import (
	"game/animatedobjects"
)

type Heal struct {
	// Image *ebiten.Image
	X, Y float64
}

func (h *Heal) Use(mh *animatedobjects.MainHero) {
	mh.Health++
}

func (h *Heal) InActiveArea(x, y int) {

}
