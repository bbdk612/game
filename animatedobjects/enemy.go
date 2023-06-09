package animatedobjects

import (
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
	Health    int
	Points    int
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

func (ms *Monster) Damage(damage int) {
	ms.Health -= damage
}

func (ms *Monster) TileCoordinate(tilesize int, x float64, y float64) int {
	TileCoor := (int(x) / tilesize) + (int(y)/16)*16
	return TileCoor
}

func (ms *Monster) CanIGo(direction Vector, chunk []int, coords [][]float64) (bool, bool) {
	MoveX, MoveY := false, false
	Tile := ms.TileCoordinate(16, ms.Position.x, ms.Position.y)
	x := ms.Position.x + direction.x*ms.Step
	y := ms.Position.y + direction.y*ms.Step
	for _, coordins := range coords {
		if ms.Position.x != coordins[0] && ms.Position.y != coordins[1] {
			dist := distance(coordins[0], x, coordins[1], y)
			if dist < 16 {
				return MoveX, MoveY
			}
		}
	}
	if direction.x < 0 {
		ms.AsePlayer.Play("left")
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
		ms.AsePlayer.Play("right")
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
	return MoveX, MoveY
}

func (ms *Monster) Actions(MHx, MHy float64, chunk []int, Coords [][]float64) [](*weapons.Bullet) {
	ms.DoesHeSeeMH(MHx, MHy)
	ms.Weapon.CalculateAngle(int(MHx), int(MHy))
	direction := Vector{MHx - ms.Position.x, MHy - ms.Position.y}
	direction.Normalize()

	ms.AsePlayer.Play("left")
	if ms.SeeMH {
		MoveX, MoveY := ms.CanIGo(direction, chunk, Coords)
		dist := distance(ms.Position.x, MHx, ms.Position.y, MHy)
		if dist > 70 {
			if MoveX {
				ms.Position.x += direction.x * ms.Step
			}
			if MoveY {
				ms.Position.y += direction.y * ms.Step
			}
			ms.Weapon.ChangePosition(int(ms.Position.x)+8, int(ms.Position.y)+8)
		} else {
			Bullets, err := ms.Weapon.Shoot(int(MHx), int(MHy), 16)
			if err != nil {
				log.Fatal(err)
			}
			return Bullets
		}
	}
	return nil
}

func distance(x1, x2, y1, y2 float64) float64 {
	dist := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return dist
}

func InitMonsters(tilecoordinate int, EnemyType string, points int) (*Monster, error) {

	var x float64 = float64((tilecoordinate % 16) * 16)
	var y float64 = float64((tilecoordinate / 16) * 16)
	var sprite string
	var health int
	var weapon *weapons.Weapon
	switch EnemyType {
	case "CommonEnemy":
		weapon, _ = weapons.InitNewWeapon(int(x)+8, int(y)+8, "./weapons/assets/enemy.json")
		sprite = "./assets/Enemy.json"
		health = 100
	case "Boss":
		weapon, _ = weapons.InitNewWeapon(int(x)+8, int(y)+8, "./assets/shotgun.json")
		sprite = "./assets/Boss.json"
		health = 300
	}
	weapon.CurrentAmmo = int(math.Inf(1))
	monster := &Monster{
		SeeMH:    false,
		Step:     2,
		Position: Vector{x: x, y: y},
		Sprite:   goaseprite.Open(sprite),
		Route: []Vector{
			{x: float64(rand.Intn(208-32) + 32), y: float64(rand.Intn(208-32) + 32)},
			{x: float64(rand.Intn(208-32) + 32), y: float64(rand.Intn(208-32) + 32)},
			{x: float64(rand.Intn(208-32) + 32), y: float64(rand.Intn(208-32) + 32)},
			{x: float64(rand.Intn(208-32) + 32), y: float64(rand.Intn(208-32) + 32)},
		},
		Weapon: weapon,
		Health: health,
		Points: points,
	}
	monster.AsePlayer = monster.Sprite.CreatePlayer()

	img, _, err := ebitenutil.NewImageFromFile(monster.Sprite.ImagePath)
	if err != nil {
		return nil, err
	}

	monster.Image = img
	return monster, nil
}
