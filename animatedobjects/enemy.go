package animatedobjects

import (
	"fmt"
	"math"

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
	//Weapon

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
	//если дисанция меньше определенного расстояния(хз какого)
	if dist < 80 {
		ms.SeeMH = true
	} else {
		ms.SeeMH = false
	}
}

func (ms *Monster) TileCoordinate(tilesize int, x float64, y float64) int {
	TileCoor := (int(x) / tilesize) + (int(y)/16)*16
	return TileCoor
}

func (ms *Monster) CanIGo(direction Vector, chunk []int) bool {
	x := ms.Position.x + direction.x*ms.Step
	y := ms.Position.y + direction.y*ms.Step
	nextTile1 := ms.TileCoordinate(16, x, y)
	nextTile2 := ms.TileCoordinate(16, x+16, y)
	nextTile3 := ms.TileCoordinate(16, x+16, y+16)
	nextTile4 := ms.TileCoordinate(16, x, y+16)
	fmt.Println(chunk[nextTile1], chunk[nextTile2], chunk[nextTile3], chunk[nextTile4])
	if chunk[nextTile1] == 1 && chunk[nextTile2] == 1 && chunk[nextTile3] == 1 && chunk[nextTile4] == 1 {
		return true
	}
	return false
}

func (ms *Monster) Actions(MHx, MHy float64, chunk []int) {
	ms.DoesHeSeeMH(MHx, MHy)
	direction := Vector{MHx - ms.Position.x, MHy - ms.Position.y}
	direction.Normalize()
	if ms.SeeMH && ms.CanIGo(direction, chunk) {
		dist := distance(ms.Position.x, MHx, ms.Position.y, MHy)
		if dist > 50 {
			ms.Position.x += direction.x * ms.Step
			ms.Position.y += direction.y * ms.Step
		} else {
			fmt.Println("fire")
		}
		//если дисанция меньше определенного расстояния(хз какого)
		// if dist < 10 {
		// 	//действия с хп главного героя
		// 	fmt.Println()
		// }
	} else {
		ms.Patrol(chunk)
	}
}

func (ms *Monster) Patrol(chunk []int) {
	target := ms.Route[0]
	direction := Vector{target.x - ms.Position.x, target.y - ms.Position.y}
	direction.Normalize()
	if ms.CanIGo(direction, chunk) {
		ms.Position.x += direction.x * ms.Step
		ms.Position.y += direction.y * ms.Step
		if math.Abs(ms.Position.x-target.x) < 1 && math.Abs(ms.Position.y-target.y) < 1 {
			ms.Route = append(ms.Route[1:], ms.Route[0])
		}
	}
}

func distance(x1, x2, y1, y2 float64) float64 {
	dist := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return dist
}

func InitMonsters(step int, tilesize int, tilecoordinate int, xCount int) (*Monster, error) {

	var x float64 = float64((tilecoordinate % xCount) * tilesize)
	var y float64 = float64((tilecoordinate / xCount) * tilesize)

	monster := &Monster{
		SeeMH:    false,
		Step:     float64(step),
		Position: Vector{x: x, y: y},
		Sprite:   goaseprite.Open("./assets/mainhero.json"),
		Route: []Vector{
			{x: 20, y: 20},
			{x: 70, y: 45},
			{x: 30, y: 100},
		},
	}
	monster.AsePlayer = monster.Sprite.CreatePlayer()

	img, _, err := ebitenutil.NewImageFromFile(monster.Sprite.ImagePath)
	if err != nil {
		return nil, err
	}

	monster.Image = img
	return monster, nil
}
