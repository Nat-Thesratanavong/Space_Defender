package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  int = 640
	screenHeight int = 640
	playerWidth  int = 15
	playerHeight int = 10
	playerSpeed  int = 5
)

type Game struct {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

type gameObjects interface {
	Draw(screen *ebiten.Image)
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyAlt) && inpututil.IsKeyJustPressed(ebiten.KeyF4) {
		return ebiten.Termination
	}
	player1.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{0, 0, 0, 0xff})

	player1.Draw(screen)
}
