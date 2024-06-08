package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	//player position
	x        int
	y        int
	score    int
	leftKey  ebiten.Key
	rightKey ebiten.Key
}

var (
	player1 = &Player{
		x:        screenWidth / 2,
		y:        screenHeight / 2,
		score:    0,
		leftKey:  ebiten.KeyA,
		rightKey: ebiten.KeyD,
	}
)

func (p *Player) Update() {
	if ebiten.IsKeyPressed(p.leftKey) && p.x > 5 {
		p.x -= playerSpeed
	}
	if ebiten.IsKeyPressed(p.rightKey) && p.x < screenWidth-5 {
		p.x += playerSpeed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(player1.x)-5, float32(player1.y)-2.5, float32(10), float32(5), color.White, true)
}
