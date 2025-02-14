package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (state *gameState) moveProjectiles() {
	for idx, projectile := range *state.Projectiles {
		newX := projectile.x + int32(projectile.scaler[0]*float32(projectile.speed))
		newY := projectile.y + int32(projectile.scaler[1]*float32(projectile.speed))
		willColide, enemy := state.willCollide(newX, newY, projectile.hitBoxRadius, -1)
		if willColide && enemy != nil {
			enemy.takeDamage(projectile.damage, state)
			state.destroyProjectile(idx)
			return
		}
		if newX+int32(projectile.hitBoxRadius) > WINDOW_WIDTH || newX-int32(projectile.hitBoxRadius) < 0 {
			state.destroyProjectile(idx)
			return
		}
		if newY+int32(projectile.hitBoxRadius) > WINDOW_HEIGHT || newY-int32(projectile.hitBoxRadius) < 0 {
			state.destroyProjectile(idx)
			return
		}
		projectile.x = newX
		projectile.y = newY
		projectile.rect.X += int32(projectile.scaler[0] * float32(projectile.speed))
		projectile.rect.Y += int32(projectile.scaler[1] * float32(projectile.speed))
		(*state.Projectiles)[idx] = projectile
	}
}

func (state *gameState) moveEnemies() {
	for idx, enemy := range *state.Enemies {
		newX := enemy.x + int32(enemy.scaler[0])
		newY := enemy.y + int32(enemy.scaler[1])
		if newX+int32(enemy.hitBoxRadius) > WINDOW_WIDTH || newX-int32(enemy.hitBoxRadius) < 0 {
			enemy.scaler[0] = enemy.scaler[0] * -1
			enemy.scaler[1] = enemy.scaler[1] * -1
			(*state.Enemies)[idx] = enemy
			continue
		}
		if newY+int32(enemy.hitBoxRadius) > WINDOW_HEIGHT || newY-int32(enemy.hitBoxRadius) < 0 {
			enemy.scaler[0] = enemy.scaler[0] * -1
			enemy.scaler[1] = enemy.scaler[1] * -1
			(*state.Enemies)[idx] = enemy
			continue
		}
		if willCollide, _ := state.willCollide(newX, newY, enemy.hitBoxRadius, idx); willCollide || state.willCollideWithPlayer(enemy.hitBoxRadius, newX, newY) {
			continue
		}
		enemy.x = newX
		enemy.y = newY
		enemy.rect.X += int32(enemy.scaler[0])
		enemy.rect.Y += int32(enemy.scaler[1])
		(*state.Enemies)[idx] = enemy
	}
}

func (enemy *Enemy) takeDamage(amount uint8, state *gameState) {
	if enemy.hp <= int8(amount) {
		enemy.destroyEnemy(state)
	} else {
		enemy.hp -= int8(amount)
	}
}

func (enemy *Enemy) destroyEnemy(state *gameState) {
	for idx, obj := range *state.Enemies {
		if obj.id == enemy.id && enemy.id != 0 {
			if idx == 0 {
				if len(*state.Enemies) > 1 {
					(*state.Enemies) = (*state.Enemies)[1:]
					return
				}
				state.Enemies = &[]Enemy{}
				return
			}
			(*state.Enemies)[idx] = (*state.Enemies)[len(*state.Enemies)-1]
			(*state.Enemies) = (*state.Enemies)[:len(*state.Enemies)-1]
			return
		}
	}
}

func (state *gameState) destroyProjectile(idx int) {
	if idx == 0 {
		if len(*state.Projectiles) > 1 {
			(*state.Projectiles) = (*state.Projectiles)[1:]
			return
		}
		state.Projectiles = &[]Projectile{}
		return
	}
	(*state.Projectiles)[idx] = (*state.Projectiles)[len(*state.Projectiles)-1]
	(*state.Projectiles) = (*state.Projectiles)[:len(*state.Projectiles)-1]
}

// Possible optimisation: store a bool if it is already colliding with smth, do a check at the beggining
// if true simply skip
func (state *gameState) willCollide(x, y int32, hitBoxRadius uint8, selfIdx int) (bool, *Enemy) {
	for idx, enemy := range *state.Enemies {
		if idx == selfIdx {
			continue
		}
		distance := math.Sqrt(math.Pow(float64(enemy.x-x), 2) + math.Pow(float64(enemy.y-y), 2))
		if distance < float64(enemy.hitBoxRadius+hitBoxRadius) {
			// using idx to create the  pointer because the variable enemy that range gives
			// is not the same as state.Enemies[idx]
			return true, &((*state.Enemies)[idx])
		}
	}
	return false, &Enemy{}
}

func (state *gameState) spawnEnemies() {
	for len(*state.Enemies) < 6 {
		color := ALL_COLORS[uint8(rand.Intn(len(ALL_COLORS)))]
		scaler := [2]int16{int16(rand.Intn(3)), int16(rand.Intn(3))}
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
		willCollideWithPlayer := state.willCollideWithPlayer(radius, x, y)
		if willCollide, _ := state.willCollide(x, y, radius, -1); !willCollide && !willCollideWithPlayer {
			state.initObject(color, scaler, x, y, w, h, hp)
		}
	}
}

func (state *gameState) willCollideWithPlayer(radius uint8, x, y int32) bool {
	distance := math.Sqrt(math.Pow(float64(state.Player.x-x), 2) + math.Pow(float64(state.Player.y-y), 2))
	return distance < float64(state.Player.hitBoxRadius+radius)
}

func (state *gameState) AssignID() uint16 {
	state.nextID++
	return state.nextID - 1
}

func calcRadius(h, w int32) float64 {
	radius := math.Sqrt(math.Pow(float64(h/2), 2) + math.Pow(float64(w/2), 2))
	if radius < 1 {
		radius = 1
	}
	return radius
}

func (state *gameState) initObject(color [3]uint8, scaler [2]int16, x, y, w, h int32, hp int8) {
	rect := sdl.Rect{X: x - w/2, Y: y - h/2, W: w, H: h}
	radius := calcRadius(h, w)
	enemy := Enemy{x: x, y: y, scaler: scaler, rect: &rect, color: color, id: state.AssignID(), hitBoxRadius: uint8(radius), hp: hp}
	*state.Enemies = append(*state.Enemies, enemy)
}

func (state *gameState) initProjectile(color [3]uint8, lifeLength uint16, speed, damage uint8, x, y, w, h int32) {
	rect := sdl.Rect{X: x - w/2, Y: y - h/2, W: w, H: h}
	radius := calcRadius(h, w)
	var scaler [2]float32
	// x = sin(a)/1
	scaler[0] = float32(math.Sin(degreeToRad(state.Player.rotation)))
	scaler[1] = -float32(math.Cos(degreeToRad(state.Player.rotation)))
	projectile := Projectile{
		x: x, id: state.AssignID(), y: y, rect: &rect, color: color,
		hitBoxRadius: uint8(radius), lifeLength: lifeLength, scaler: scaler, speed: speed, damage: damage,
	}
	*state.Projectiles = append(*state.Projectiles, projectile)
}

func degreeToRad(angle int16) float64 {
	return float64(angle) * math.Pi / 180
}

// texturePath := ,,media/name.png
func (state *gameState) initPlayer(x, y int32, speed uint8, texturePath string) {
	texture, err := img.LoadTexture(state.renderer, texturePath)
	if err != nil {
		panic(err)
	}
	eventList := make([]bool, 6)

	_, _, w, h, err := texture.Query()
	if err != nil {
		panic(err)
	}

	radius := calcRadius(h, w)
	player := Player{
		x: x, y: y, texture: texture, id: state.AssignID(), speed: speed,
		eventList: eventList, hitBoxRadius: uint8(radius),
		ammo: PLAYER_MAG_SIZE, magazine_size: PLAYER_MAG_SIZE, cooldown: 0,
	}
	state.Player = &player
}
