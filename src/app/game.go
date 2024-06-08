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

var (
// score              = "SCORE\n" + strconv.Itoa(player1.score)
// firaSansRegular    []byte
// firaSansFaceSource *text.GoTextFaceSource
// op                 = &text.DrawOptions{}
)

// func init() {
// 	s, err := text.NewGoTextFaceSource(bytes.NewReader(firaSansRegular))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	firaSansFaceSource = s
// }

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
	bullet1.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{0, 0, 0, 0xff})
	player1.Draw(screen)
	bullet1.Draw(screen)

	// op.GeoM.Translate(20, 20)
	// op.LineSpacing = 30
	// text.Draw(screen, score, &text.GoTextFace{
	// 	Source: firaSansFaceSource,
	// 	Size:   20,
	// }, op)
}
