package main

import (
	"image/color"

	_ "log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Bullet struct {
	speed       int
	x           int
	y           int
	bulletState bool
}

var (
	keyState_space bool = false
	bullet1             = Bullet{
		speed:       10,
		x:           player1.x,
		y:           player1.y,
		bulletState: false,
	}
)

func (b *Bullet) Update() {
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !keyState_space {
		b.bulletState = true
		keyState_space = true
		b.y = player1.y
		b.x = player1.x
	}
	if b.bulletState {

		if b.y > 10 {
			// log.Println("b.y")
			b.y -= b.speed
		} else {
			b.bulletState = false
			keyState_space = false
		}
	}
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	if b.bulletState {
		vector.DrawFilledRect(screen, float32(bullet1.x), float32(bullet1.y)-5.0, float32(2), float32(5), color.White, true)
	}
}
