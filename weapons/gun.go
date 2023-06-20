package weapons

import (
	// "fmt"

	"fmt"
	"math"
)

func gunShoot(oX, oY float64, directionX, directionY float64, spritePath string, tilesize int) (*Bullet, error) {
	var deltaX float64 = float64(oX) - float64(directionX)
	var deltaY float64 = float64(oY) - float64(directionY)
	var startY float64 = float64(oY)
	var startX float64 = float64(oX)
	var a, b float64
	var step float64 = 3
	a = deltaY / deltaX
	if a == 0 {
		if deltaX < 0 {
			bullet, err := InitNewBullet(a, 0, step, startX+8, startY, "./assets/bullet.json", 16)
			if err != nil {
				return nil, err
			}
			bullet.CalculateNextStep = func(x, y float64, a, b, step float64) (float64, float64) {
				x += step
				return x, y
			}
			fmt.Println("shoot")
			return bullet, nil
		} else {
			bullet, err := InitNewBullet(a, 0, -step, startX-8, startY, "./assets/bullet.json", 16)
			if err != nil {
				return nil, err
			}
			bullet.CalculateNextStep = func(x, y float64, a, b, step float64) (float64, float64) {
				x += step
				return x, y
			}
			return bullet, nil
		}
	} else if a == math.Inf(-1) || a == math.Inf(1) {
		var bullet *Bullet
		var err error
		if a > 0 {
			bullet, err = InitNewBullet(a, b, (-step), startX, startY-8, "./assets/bullet.json", 16)
		} else {
			bullet, err = InitNewBullet(a, b, step, startX, startY+8, "./assets/bullet.json", 16)
		}
		if err != nil {
			return nil, err
		}
		bullet.CalculateNextStep = func(x, y float64, a, b, step float64) (float64, float64) {
			y += step
			return x, y
		}
		return bullet, nil
	} else if a > 0 && a != math.Inf(1) {
		b = startY - (startX * a)
		if math.Abs(deltaX) > math.Abs(deltaY) {
			var bullet *Bullet
			var err error
			if deltaY > 0 {
				bullet, err = InitNewBullet(a, b, (-step), startX-8, startY, "./assets/bullet.json", 16)
			} else {
				bullet, err = InitNewBullet(a, b, step, startX+8, startY, "./assets/bullet.json", 16)
			}
			if err != nil {
				return nil, err
			}
			bullet.CalculateNextStep = func(x, y float64, a, b, step float64) (float64, float64) {
				x += step
				y = (a * x) + b
				return x, y
			}

			return bullet, nil
		} else {
			var bullet *Bullet
			var err error
			if deltaX > 0 {
				bullet, err = InitNewBullet(a, b, (-step), startX, startY-8, "./assets/bullet.json", 16)
			} else {
				bullet, err = InitNewBullet(a, b, step, startX, startY+8, "./assets/bullet.json", 16)
			}
			if err != nil {
				return nil, err
			}
			bullet.CalculateNextStep = func(x, y float64, a, b, step float64) (float64, float64) {
				y += step
				x = (y - b) / a
				return x, y
			}
			return bullet, nil
		}

	} else {
		b = startY - (startX * a)

		if math.Abs(deltaX) > math.Abs(deltaY) {
			var bullet *Bullet
			var err error
			if deltaY > 0 {
				bullet, err = InitNewBullet(a, b, (step), startX+8, startY, "./assets/bullet.json", 16)
			} else {
				bullet, err = InitNewBullet(a, b, (-step), startX-8, startY, "./assets/bullet.json", 16)
			}
			if err != nil {
				return nil, err
			}
			bullet.CalculateNextStep = func(x, y float64, a, b, step float64) (float64, float64) {
				x += step
				y = (a * x) + b
				return x, y
			}

			return bullet, nil
		} else {
			var bullet *Bullet
			var err error
			if deltaX > 0 {
				bullet, err = InitNewBullet(a, b, (step), startX, startY+8, "./assets/bullet.json", 16)
			} else {
				bullet, err = InitNewBullet(a, b, (-step), startX, startY-8, "./assets/bullet.json", 16)
			}
			if err != nil {
				return nil, err
			}
			bullet.CalculateNextStep = func(x, y float64, a, b, step float64) (float64, float64) {
				y += step
				x = (y - b) / a
				return x, y
			}
			return bullet, nil
		}
	}
}
