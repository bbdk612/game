package weapons

import (
	"fmt"
	"math"
)

func shotgunShoot(startX, startY float64, deltaX, deltaY float64, spritePath string, tilesize int, a, b float64) (*Bullet, error) {
	var step float64 = 2
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
