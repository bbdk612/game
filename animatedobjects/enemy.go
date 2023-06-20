package animatedobjects

import (
	"fmt"
	"game/weapons"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/solarlune/goaseprite"
)

type Monster struct {
	SeeMH     bool
	Step      float64
	Position  Vector
	Sprite    *goaseprite.File
	AsePlayer *goaseprite.Player
	Image     *ebiten.Image
	Route     []Vector
	Weapon    *weapons.Weapon
}

type Vector struct {
	x, y float64
}

func (v *Vector) Normalize() {
	length := math.Sqrt(v.x*v.x + v.y*v.y)
	v.x /= length
	v.y /= length
}

func (ms *Monster) GetCoordinates() (float64, float64) {
	return ms.Position.x, ms.Position.y
}

func (ms *Monster) DoesHeSeeMH(x, y float64) {
	dist := math.Abs(distance(ms.Position.x, x, ms.Position.y, y))
	if dist < 300 {
		ms.SeeMH = true
	} else {
		ms.SeeMH = false
	}
}

func (ms *Monster) TileCoordinate(tilesize int, x float64, y float64) int {
	TileCoor := (int(x) / tilesize) + (int(y)/16)*16
	return TileCoor
}

func (ms *Monster) CanIGo(direction Vector, chunk []int) (bool, bool) {
	MoveX, MoveY := false, false
	Tile := ms.TileCoordinate(16, ms.Position.x, ms.Position.y)
	if direction.x < 0 {
		if int(ms.Position.x)%16 != 0 {
			if int(ms.Position.y)%16 != 0 {
				if (chunk[Tile] == 1) && (chunk[Tile+16] == 1) {
					MoveX = true
				}
			}
			if chunk[Tile] == 1 {
				MoveX = true
			}
		} else if int(ms.Position.y)%16 != 0 {
			if (chunk[Tile-1] == 1) && (chunk[Tile-1+16] == 1) {
				MoveX = true
			}
		} else if chunk[Tile-1] == 1 {
			MoveX = true
		}
	} else if direction.x > 0 {
		if int(ms.Position.y)%16 != 0 {
			if (chunk[Tile+1] == 1) && (chunk[Tile+1+16] == 1) {
				MoveX = true
			}
		} else if chunk[Tile+1] == 1 {
			MoveX = true
		}
	}
	if direction.y < 0 {
		if int(ms.Position.y)%16 != 0 {
			if int(ms.Position.x)%16 != 0 {
				if (chunk[Tile] == 1) && (chunk[Tile+1] == 1) {
					MoveY = true
					return MoveX, MoveY
				}
			}
			if chunk[Tile] == 1 {
				MoveY = true
				return MoveX, MoveY
			}
		} else if int(ms.Position.x)%16 != 0 {
			if (chunk[Tile-16] == 1) && (chunk[Tile-16+1] == 1) {
				MoveY = true
				return MoveX, MoveY
			}
		} else if chunk[Tile-16] == 1 {
			MoveY = true
			return MoveX, MoveY
		}
	} else if direction.y > 0 {
		if int(ms.Position.x)%16 != 0 {
			if (chunk[Tile+16] == 1) && (chunk[Tile+1+16] == 1) {
				MoveY = true
				return MoveX, MoveY
			}
		} else if chunk[Tile+16] == 1 {
			MoveY = true
			return MoveX, MoveY
		}
	}
	fmt.Println(chunk[Tile], chunk[Tile+16], chunk[Tile-1], chunk[Tile-1+16], chunk[Tile+1], chunk[Tile+1+16])
	return MoveX, MoveY
}

func (ms *Monster) Actions(MHx, MHy float64, chunk []int) [](*weapons.Bullet) {
	ms.DoesHeSeeMH(MHx, MHy)
	direction := Vector{MHx - ms.Position.x, MHy - ms.Position.y}
	direction.Normalize()
	if ms.SeeMH {
		MoveX, MoveY := ms.CanIGo(direction, chunk)
		dist := distance(ms.Position.x, MHx, ms.Position.y, MHy)
		if dist > 70 {
			if MoveX {
				ms.Position.x += direction.x * ms.Step
			}
			if MoveY {
				ms.Position.y += direction.y * ms.Step
			}
		} else {
			Bullets, err := ms.Weapon.Shoot(int(MHx), int(MHy), 16)
			if err != nil {
				log.Fatal(err)
			}
			return Bullets
		}
	}
	return nil
	// else {
	// 	ms.Patrol(chunk)
	// }
}

// func (ms *Monster) Patrol(chunk []int) {
// 	target := ms.Route[0]
// 	direction := Vector{target.x - ms.Position.x, target.y - ms.Position.y}
// 	direction.Normalize()
// 	MoveX, MoveY := ms.CanIGo(direction, chunk)
// 	if MoveX {
// 		ms.Position.x += direction.x * ms.Step
// 	}
// 	if MoveY {
// 		ms.Position.y += direction.y * ms.Step
// 	}
// 	fmt.Println(math.Abs(ms.Position.x-target.x), math.Abs(ms.Position.y-target.y), MoveX, MoveY)
// 	if math.Abs(ms.Position.x-target.x) <= 20 && math.Abs(ms.Position.y-target.y) > 20 && !MoveY {
// 		ms.Route = append(ms.Route[1:], ms.Route[0])
// 	}
// 	if math.Abs(ms.Position.x-target.x) > 20 && math.Abs(ms.Position.y-target.y) <= 20 && !MoveX {
// 		ms.Route = append(ms.Route[1:], ms.Route[0])
// 	}
// 	if math.Abs(ms.Position.x-target.x) <= 20 && math.Abs(ms.Position.y-target.y) <= 20 {
// 		ms.Route = append(ms.Route[1:], ms.Route[0])
// 	}
// 	if !MoveY && !MoveX {
// 		ms.Route = append(ms.Route[1:], ms.Route[0])
// 	}
// }

func distance(x1, x2, y1, y2 float64) float64 {
	dist := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return dist
}

func InitMonsters(step int, tilesize int, tilecoordinate int, xCount int) (*Monster, error) {

	var x float64 = float64((tilecoordinate % xCount) * tilesize)
	var y float64 = float64((tilecoordinate / xCount) * tilesize)
	weapon, err := weapons.InitNewWeapon(int(x+8), int(y+8), "./weapons/assets/shotgun.json")
	monster := &Monster{
		SeeMH:    false,
		Step:     float64(step),
		Position: Vector{x: x, y: y},
		Sprite:   goaseprite.Open("./assets/mainhero.json"),
		Route: []Vector{
			{x: float64(rand.Intn(208-32) + 32), y: float64(rand.Intn(208-32) + 32)},
			{x: float64(rand.Intn(208-32) + 32), y: float64(rand.Intn(208-32) + 32)},
			{x: float64(rand.Intn(208-32) + 32), y: float64(rand.Intn(208-32) + 32)},
			{x: float64(rand.Intn(208-32) + 32), y: float64(rand.Intn(208-32) + 32)},
		},
		Weapon: weapon,
	}
	monster.AsePlayer = monster.Sprite.CreatePlayer()

	img, _, err := ebitenutil.NewImageFromFile(monster.Sprite.ImagePath)
	if err != nil {
		return nil, err
	}

	monster.Image = img
	return monster, nil
}
