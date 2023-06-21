package weapons

import (
	"encoding/json"
	"fmt"
	"image"
	"math"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Weapon struct {
	angle                float64
	SpreadAngle          float64
	oX, oY               int
	endX, endY           int
	CurrentAmmo, MaxAmmo int
	NumberOfBullets      int
	Damage               int
	BulletSprite         string
	rollbackTime         time.Time
	DurationMS           string
	WeaponType           string
	ImagePath            string
	Image                *ebiten.Image
}

func (w *Weapon) CalculateAngle(x int, y int) {
	var aY float64 = (float64(w.oY) - float64(y))
	var aX float64 = (float64(w.oX) - float64(x))
	var bY float64 = (float64(w.oY) - float64(w.endY))
	var bX float64 = (float64(w.oX) - float64(w.endX))

	w.angle = (-math.Atan2(bY, bX) + math.Atan2(aY, aX))
}

func (w *Weapon) GetAmmo() (int, int) {
	return w.CurrentAmmo, w.MaxAmmo
}

func (w *Weapon) GetAngle() float64 {
	return w.angle
}

func (w *Weapon) GetOCoordinates() (int, int) {
	return w.oX, w.oY
}

func (w *Weapon) ChangePosition(x, y int) {
	w.oX = x
	w.oY = y

	w.endX = x + 8
	w.endY = y + 8
}

func (w *Weapon) Reload() {
	w.CurrentAmmo = w.MaxAmmo
}

func (w *Weapon) Shoot(directionX, directionY int, tilesize int) ([](*Bullet), error) {
	if w.CurrentAmmo != 0 {
		rlbkDur, err := time.ParseDuration(w.DurationMS)
		if err != nil {
			return nil, err
		}
		rollbk := rlbkDur.Milliseconds()

		currTime := time.Now()
		if currTime.Sub(w.rollbackTime).Milliseconds() >= rollbk {
			w.rollbackTime = time.Now()

			switch w.WeaponType {
			case "rifle":
			case "gun":
				bullet, err := gunShoot(w.Damage, float64(w.oX), float64(w.oY), float64(directionX), float64(directionY), w.BulletSprite, 16)
				if err != nil {
					return nil, err
				}

				w.CurrentAmmo -= 1
				bullets := [](*Bullet){bullet}
				return bullets, nil

			case "shotgun":
				bullets := [](*Bullet){}
				var dAngle float64 = w.SpreadAngle / float64(w.NumberOfBullets-1)
				var angle float64 = w.SpreadAngle / 2
				for i := 0; i < w.NumberOfBullets; i++ {
					var endX float64 = (-math.Sin(angle) * float64(directionY-w.oY)) + (math.Cos(angle) * float64(directionX-w.oX)) + float64(w.oX)
					var endY float64 = (math.Cos(angle) * float64(directionY-w.oY)) + (math.Sin(angle) * float64(directionX-w.oX)) + float64(w.oY)

					bullet, err := gunShoot(w.Damage, float64(w.oX), float64(w.oY), endX, endY, w.BulletSprite, 16)

					if err != nil {
						return nil, err
					}

					bullets = append(bullets, bullet)
					angle -= dAngle

				}

				w.CurrentAmmo -= 1
				return bullets, nil

			default:
				return nil, nil
			}
		}

	}

	return nil, nil
}

func InitNewWeapon(spawnX, spawnY int, JSONPath string) (*Weapon, error) {
	w := &Weapon{
		oX:   spawnX,
		oY:   spawnY,
		endX: spawnX + 8,
		endY: spawnY + 8,
	}
	data, err := os.ReadFile(JSONPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &w)
	if err != nil {
		return nil, err
	}

	imagefile, err := os.Open(w.ImagePath)
	if err != nil {
		return nil, err
	}

	imagedecoded, _, err := image.Decode(imagefile)
	if err != nil {
		return nil, err
	}

	w.Image = ebiten.NewImageFromImage(imagedecoded)

	w.ChangePosition(spawnX, spawnY)
	fmt.Println(w)
	return w, nil
}
