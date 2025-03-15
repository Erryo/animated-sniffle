package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func (lvl *level) moveProjectiles() {
	for idx, projectile := range *lvl.projectiles {
		newX := projectile.x + int32(projectile.vector[0]*float32(projectile.speed))
		newY := projectile.y + int32(projectile.vector[1]*float32(projectile.speed))
		willColide, enemy := lvl.willCollide(newX, newY, projectile.hitBoxRadius, -1)
		if willColide && enemy != nil {
			enemy.takeDamage(projectile.damage, lvl)
			lvl.destroyProjectile(idx)
			return
		}
		if newX+int32(projectile.hitBoxRadius) > WINDOW_WIDTH || newX-int32(projectile.hitBoxRadius) < 0 {
			lvl.destroyProjectile(idx)
			return
		}
		if newY+int32(projectile.hitBoxRadius) > WINDOW_HEIGHT || newY-int32(projectile.hitBoxRadius) < 0 {
			lvl.destroyProjectile(idx)
			return
		}
		projectile.x = newX
		projectile.y = newY
		projectile.rect.X += int32(projectile.vector[0] * float32(projectile.speed))
		projectile.rect.Y += int32(projectile.vector[1] * float32(projectile.speed))
		(*lvl.projectiles)[idx] = projectile
	}
}

func (lvl *level) moveEnemies() {
	for idx, enemy := range *lvl.enemies {
		newX := enemy.x + int32(enemy.vector[0])
		newY := enemy.y + int32(enemy.vector[1])
		if newX+int32(enemy.hitBoxRadius) > WINDOW_WIDTH || newX-int32(enemy.hitBoxRadius) < 0 {
			enemy.vector[0] = enemy.vector[0] * -1
			(*lvl.enemies)[idx] = enemy
			continue
		}
		if newY+int32(enemy.hitBoxRadius) > WINDOW_HEIGHT || newY-int32(enemy.hitBoxRadius) < 0 {
			enemy.vector[1] = enemy.vector[1] * -1
			(*lvl.enemies)[idx] = enemy
			continue
		}
		if willCollide, _ := lvl.willCollide(newX, newY, enemy.hitBoxRadius, idx); willCollide || lvl.willCollideWithPlayer(enemy.hitBoxRadius, newX, newY) {
			continue
		}
		enemy.x = newX
		enemy.y = newY
		enemy.rect.X += int32(enemy.vector[0])
		enemy.rect.Y += int32(enemy.vector[1])
		(*lvl.enemies)[idx] = enemy
	}
}

func (enemy *enemy) takeDamage(amount uint8, lvl *level) {
	if enemy.hp <= int8(amount) {
		enemy.destroyEnemy(lvl)
	} else {
		enemy.hp -= int8(amount)
	}
}

func (enemy *enemy) destroyEnemy(lvl *level) {
	for idx, obj := range *lvl.enemies {
		if obj.id == enemy.id && enemy.id != 0 {
			if idx == 0 {
				if len(*lvl.enemies) > 1 {
					(*lvl.enemies) = (*lvl.enemies)[1:]
					return
				}
				// state.enemies = &[]enemy{}
				return
			}
			(*lvl.enemies)[idx] = (*lvl.enemies)[len(*lvl.enemies)-1]
			(*lvl.enemies) = (*lvl.enemies)[:len(*lvl.enemies)-1]
			return
		}
	}
}

func (lvl *level) destroyProjectile(idx int) {
	if idx == 0 {
		if len(*lvl.projectiles) > 1 {
			(*lvl.projectiles) = (*lvl.projectiles)[1:]
			return
		}
		lvl.projectiles = &[]projectile{}
		return
	}
	(*lvl.projectiles)[idx] = (*lvl.projectiles)[len(*lvl.projectiles)-1]
	(*lvl.projectiles) = (*lvl.projectiles)[:len(*lvl.projectiles)-1]
}

// Possible optimisation: store a bool if it is already colliding with smth, do a check at the beggining
// if true simply skip
func (lvl *level) willCollide(x, y int32, hitBoxRadius uint8, selfIdx int) (bool, *enemy) {
	for idx, enemy := range *lvl.enemies {
		if idx == selfIdx {
			continue
		}
		distance := math.Sqrt(math.Pow(float64(enemy.x-x), 2) + math.Pow(float64(enemy.y-y), 2))
		if distance < float64(enemy.hitBoxRadius+hitBoxRadius) {
			// using idx to create the  pointer because the variable enemy that range gives
			// is not the same as state.enemies[idx]
			return true, &((*lvl.enemies)[idx])
		}
	}
	return false, &enemy{}
}

func (lvl *level) spawnEnemies() {
	for len(*lvl.enemies) < 6 {
		color := ALL_COLORS[uint8(rand.Intn(len(ALL_COLORS)))]
		vector := [2]int16{int16(rand.Intn(3)), int16(rand.Intn(3))}
		x, y := int32(rand.Intn(WINDOW_WIDTH)), int32(rand.Intn(WINDOW_HEIGHT))
		w, h := int32(rand.Intn(40)), int32(rand.Intn(40))
		hp := int8(math.Sqrt(float64(w)*float64(h)) / 10)
		if hp < 1 {
			continue
		}
		fmt.Println(hp)
		radius := uint8(calcRadius(h, w))

		if y+int32(radius) > WINDOW_HEIGHT || y-int32(radius) < 0 {
			continue
		}
		if x+int32(radius) > WINDOW_WIDTH || x-int32(radius) < 0 {
			continue
		}
		// Det coll with Player
		willCollideWithPlayer := lvl.willCollideWithPlayer(radius, x, y)
		if willCollide, _ := lvl.willCollide(x, y, radius, -1); !willCollide && !willCollideWithPlayer {
			lvl.initObject(color, vector, x, y, w, h, hp)
		}
	}
}

func (lvl *level) willCollideWithPlayer(radius uint8, x, y int32) bool {
	distance := math.Sqrt(math.Pow(float64(lvl.player.x-x), 2) + math.Pow(float64(lvl.player.y-y), 2))
	return distance < float64(lvl.player.hitBoxRadius+radius)
}

func (lvl *level) AssignID() uint16 {
	lvl.nextID++
	return lvl.nextID - 1
}

func calcRadius(h, w int32) float64 {
	radius := math.Sqrt(math.Pow(float64(h/2), 2) + math.Pow(float64(w/2), 2))
	if radius < 1 {
		radius = 1
	}
	return radius
}

func (lvl *level) initObject(color [3]uint8, vector [2]int16, x, y, w, h int32, hp int8) {
	rect := sdl.Rect{X: x - w/2, Y: y - h/2, W: w, H: h}
	radius := calcRadius(h, w)
	enemy := enemy{x: x, y: y, vector: vector, rect: &rect, color: color, id: lvl.AssignID(), hitBoxRadius: uint8(radius), hp: hp}
	*lvl.enemies = append(*lvl.enemies, enemy)
}

func (lvl *level) initProjectile(color [3]uint8, lifeLength uint16, speed, damage uint8, x, y, w, h int32) {
	rect := sdl.Rect{X: x - w/2, Y: y - h/2, W: w, H: h}
	radius := calcRadius(h, w)
	var vector [2]float32
	// x = sin(a)/1
	vector[0] = float32(math.Sin(degreeToRad(lvl.player.rotation)))
	vector[1] = -float32(math.Cos(degreeToRad(lvl.player.rotation)))
	projectile := projectile{
		x: x, id: lvl.AssignID(), y: y, rect: &rect, color: color,
		hitBoxRadius: uint8(radius), lifeLength: lifeLength, vector: vector, speed: speed, damage: damage,
	}
	*lvl.projectiles = append(*lvl.projectiles, projectile)
}

func degreeToRad(angle int16) float64 {
	return float64(angle) * math.Pi / 180
}

// texturePath := ,,media/name.png
func (lvl *level) initPlayer(x, y int32, speed uint8, texturePath string, renderer *sdl.Renderer) {
	var err error
	var shootEff *mix.Chunk
	texture, err := img.LoadTexture(renderer, texturePath)
	if err != nil {
		panic(err)
	}
	if shootEff, err = mix.LoadWAV("media/shoot.ogg"); err != nil {
		panic(err)
	}
	eventList := make([]bool, 8)

	_, _, w, h, err := texture.Query()
	if err != nil {
		panic(err)
	}

	radius := calcRadius(h, w)
	player := player{
		x: x, y: y, texture: texture, id: lvl.AssignID(), speed: speed,
		eventList: eventList, hitBoxRadius: uint8(radius),
		ammo: PLAYER_MAG_SIZE, magazine_size: PLAYER_MAG_SIZE, cooldown: 0,
		shootEff: shootEff, vector: [2]int8{0, 0},
	}
	lvl.player = &player
}
