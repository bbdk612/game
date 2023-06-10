package animatedobjects

import (
	"math"
)

type Monster struct {
	x, y int
	SeeMH bool
	Damage int
	Step float64
	Position Vector
	//Weapon 
}	

type Vector struct {
	x, y float64
}

func (v *Vector) Normalize() {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y)
	v.X /= length
	v.Y /= length
}

func (ms *Monster) DoesHeSeeMH(G *Game) {
	MHx, MHy := G.animatedobjects.MainHero.GetCoordinates
	dist = math.Abs(distance(ms.x, MHx, ms.y, MHy))
	//если дисанция меньше определенного расстояния(хз какого)
	if dist <  {
		ms.SeeMH = true
	}
	else {
		ms.SeeMH = false
	}
}

func (ms *Monster) Actions(G *Game) {
	MHx, MHy := G.animatedobjects.MainHero.GetCoordinates
	ms.DoesHeSeeMH() 
	if ms.SeeMH {
		direction := Vector{MHx - ms.x, MHy - ms.y}
		direction.Normalize()
		ms.Position.x += direction.x + ms.Step
		ms.Position.y += direction.y + ms.Step
		dist := distance(ms.x, MHx, ms.y, MHy)
		//если дисанция меньше определенного расстояния(хз какого)
		if dist <  {
			//действия с хп главного героя
		}
	}
	else {

	}
}

func (ms *Monster) Patrol() {
	route := []Vector {
		//координаты маршрута
	}
	target := route[0]
	direction := Vector{MHx - ms.x, MHy - ms.y}
	direction.Normalize()
	ms.Position.x += direction.x + ms.Step
	ms.Position.y += direction.y + ms.Step
	if  math.Abs(ms.x - target.x) < 0.1 && math.Abs(ms.y-target.y) < 0.1 {
		route = append(route[1:], route[0])
}

func distance(x1, x2, y1, y2 int) (int) {
	dist := math.sqrt(math.pow(x2 - x1, 2) + math.pow(y2 - y1, 2))
	return dist
}

func InitMonsters(step int, tilesize int) (*Monster, error) {
	monster := &Monster {
		x: x,
		y: y,
		SeeMH: false, 
		Damage: , 
		Step: step,
		Position: Vector{x: X, y: Y},
	}
}